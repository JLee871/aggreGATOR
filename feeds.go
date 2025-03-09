package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/JLee871/aggreGATOR/internal/database"
	"github.com/google/uuid"
)

// Add a feed to database
func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	v := reflect.ValueOf(feed)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Println(v.Type().Field(i).Name, "-", field.Interface())
	}

	return nil
}

// Return all feeds
func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.DB.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		username, err := s.DB.GetUserNameFromID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println("Name:", feed.Name, ", URL:", feed.Url, ", User:", username)
	}

	return nil
}

// Follow a feed
func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	fmt.Println(feedFollow.FeedName)
	fmt.Println(feedFollow.UserName)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("failed to unfollow")
	}

	fmt.Println("Unfollowed", feed.Name)

	return nil
}

// Return all followed feeds by user
func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}
	feedsFollow, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	for _, feed := range feedsFollow {
		fmt.Println(feed.FeedName)
	}

	return nil
}
