package main

import (
	"fmt"
	"os"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.registeredCommands[cmd.Name]
	if !ok {
		os.Exit(1)
		return fmt.Errorf("invalid Command")
	}
	return cmdFunc(s, cmd)
}
