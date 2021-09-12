package cli

import (
	"errors"
	"os"
)

func Parse(args []string) (*Command, error) {
	if len(args) < 2 {
		return nil, errors.New("you need to specify a service (sqs,iam,lambda)")
	}

	commands := []Command{
		newSqsCmd(),
		newLambdaCmd(),
		newIamCmd(),
		newEventBridgeCmd(),
		newCloudTrailCmd(),
	}

	service := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name() == service {
			return &cmd, cmd.init(os.Args[2:])
		}
	}

	return nil, errors.New("Unknown service: " + service)

}
