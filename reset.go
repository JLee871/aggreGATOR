package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		os.Exit(1)
		return err
	}

	fmt.Println("user database reset")

	return nil
}
