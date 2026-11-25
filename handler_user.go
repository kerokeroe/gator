package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is required")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", username)

	return nil
}
