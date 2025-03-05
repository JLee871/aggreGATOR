package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/JLee871/aggreGATOR/internal/config"
	"github.com/JLee871/aggreGATOR/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	DB     *database.Queries
	Config *config.Config
}

func main() {
	dbURL := "postgres://postgres:postgres@localhost:5432/gator"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(db)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	newState := state{Config: &cfg, DB: dbQueries}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	//Reset Handler
	cmds.register("reset", handlerReset)

	//Handlers related to users found in users.go
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerUsers)

	//Aggregation Handler
	cmds.register("agg", handlerAgg)

	//Handlers found in feeds.go
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", handlerFollowing)

	osArgs := os.Args
	if len(osArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmd := command{
		Name: osArgs[1],
		Args: osArgs[2:],
	}

	err = cmds.run(&newState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
