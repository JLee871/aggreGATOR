package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/JLee871/aggreGATOR/internal/database"
	"github.com/google/uuid"
)

// Sets current user
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	_, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		os.Exit(1)
		return err
	}

	err = s.Config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been set to %s.\n", name)
	return nil
}

// Adds a new user to the db
func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	_, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name})
	if err != nil {
		return err
	}
	fmt.Printf("New user created: %s\n", name)
	//fmt.Println(user)

	err = s.Config.SetUser(name)
	if err != nil {
		return err
	}

	return nil
}

// Returns all users in db
func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	users, err := s.DB.GetAllUsers(context.Background())
	if err != nil {
		os.Exit(1)
		return err
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}

	return nil
}
