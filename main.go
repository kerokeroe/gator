package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kerokeroe/gator/internal/config"
	"github.com/kerokeroe/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
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
	cmds.register("register", handlerRegister)

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	dbURL := cfg.DBURL
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	db := database.New(conn)
	s := &state{cfg: &cfg, db: db}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
