package main

import (
	"github.com/urfave/cli/v2"
	"learn_cli/learn_command"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "greeter",
		Usage: "A simple CLI application",
		Commands: []*cli.Command{
			learn_command.GreetCommand(),
			learn_command.CreatePostCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
