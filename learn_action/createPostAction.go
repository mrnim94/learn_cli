package learn_action

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func CreatePostAction(c *cli.Context) error {
	title := c.String("title")
	content := c.String("content")

	// Here, we just print the post details. In a real application, you might save this data to a file or database.
	fmt.Printf("Post Created:\nTitle: %s\nContent: %s\n", title, content)
	return nil
}
