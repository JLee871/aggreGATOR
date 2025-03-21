package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/JLee871/aggreGATOR/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Scrapes feeds after given interval
func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time between requests>", cmd.Name)
	}

	timeBtwnReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBtwnReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch next feed")
	}

	err = s.DB.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched")
	}

	rss, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	timeFormats := []string{
		time.RFC1123Z,
		time.RFC1123,
		"Mon, 2 Jan 2006 15:04:05 MST",
	}

	var publishedTime time.Time
	for _, item := range rss.Channel.Item {
		//Parse the published time
		for _, format := range timeFormats {
			parsedTime, err := time.Parse(format, item.PubDate)
			if err == nil {
				publishedTime = parsedTime
				break
			}
		}

		//Add post to db
		_, err := s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       sql.NullString{String: item.Title, Valid: len(item.Title) != 0},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: len(item.Description) != 0},
			PublishedAt: sql.NullTime{Time: publishedTime, Valid: !publishedTime.IsZero()},
			FeedID:      nextFeed.ID,
		})

		if err != nil {
			if pgerr, ok := err.(*pq.Error); ok {
				if pgerr.Code == "23505" {
					_ = pgerr
				} else {
					log.Println("Error: failed to add post to db")
				}
			}
		} else {
			fmt.Printf("added a new post: %v\n", item.Link)
		}
	}
	fmt.Println("---------")

	return nil
}
