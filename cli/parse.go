package cli

import (
	"errors"
	"os"
)

func Parse(args []string) (*Command, error) {
	if len(args) < 1 {
		return nil, errors.New("You need to specify a service (sqs,iam,lambda)")
	}

	cmds := []Command{
		newSqsCmd(),
		newLambdaCmd(),
		newIamCmd(),
	}

	service := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == service {
			return &cmd, cmd.init(os.Args[2:])
		}
	}

	return nil, errors.New("Unknown service: " + service)

}
