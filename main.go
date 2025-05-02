package main

import (
	"errors"
	"fmt"
	"gator/internal/config"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	functions map[string]func(*state, command) error
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := &commands{
		functions: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("expected a command")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New(
			"the login handler expects a single argument, the username",
		)
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("The user %s has been set.\n", cmd.args[0])
	return nil

}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.functions[cmd.name]
	if !exists {
		return fmt.Errorf("command %s not found", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.functions[name] = f
}
