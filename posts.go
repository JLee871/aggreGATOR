package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JLee871/aggreGATOR/internal/database"
)

const defaultLimit = 2

// Prints feeds to cli for user up to the limit
func handlerBrowse(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <number>", cmd.Name)
	}

	num := defaultLimit
	if len(cmd.Args) == 1 {
		var err error
		num, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("number argument was invalid")
		}
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{Name: s.Config.CurrentUserName, Limit: int32(num)})
	if err != nil {
		return fmt.Errorf("could not get posts from db")
	}

	for _, post := range posts {
		fmt.Println(post.Title.String)
		fmt.Println(post.Url)
		fmt.Println(post.PublishedAt.Time)
		fmt.Println("=================================")
	}

	return nil
}
