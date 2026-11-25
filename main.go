package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kerokeroe/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	cfg, err := config.Read()
	s := &state{cfg: &cfg}
	if err != nil {
		log.Fatal(err)
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
