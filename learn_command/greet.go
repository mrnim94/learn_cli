package learn_command

import (
	"github.com/urfave/cli/v2"
	"learn_cli/learn_action"
)

func GreetCommand() *cli.Command {
	return &cli.Command{
		Name:    "greet",
		Aliases: []string{"g"},
		Usage:   "greets a user",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "the name of the person to greet",
			},
			&cli.IntFlag{
				Name:  "age",
				Value: 0,
				Usage: "the age of the person",
			},
		},
		Action: learn_action.GreetAction,
	}
}
