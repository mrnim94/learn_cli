package learn_command

import (
	"github.com/urfave/cli/v2"
	"learn_cli/learn_action"
)

func CreatePostCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create a resource",
		Subcommands: []*cli.Command{
			{
				Name:   "post",
				Usage:  "create a post with title and content",
				Action: learn_action.CreatePostAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "title",
						Usage: "title of the post",
					},
					&cli.StringFlag{
						Name:  "content",
						Usage: "content of the post",
					},
				},
			},
		},
	}
}
