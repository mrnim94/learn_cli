package learn_action

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func GreetAction(c *cli.Context) error {
	name := c.String("name")
	if name == "" {
		name = "World" // Default value if no name is provided
	}
	age := c.Int("age")

	greeting := fmt.Sprintf("Hello, %s!", name)
	if age > 0 {
		greeting += fmt.Sprintf(" You are %d years old.", age)
	}

	fmt.Println(greeting)
	return nil
}
