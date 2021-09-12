package main

import (
	"fmt"
	"os"

	"github.com/MauricioAntonioMartinez/aws/cli"
	"github.com/MauricioAntonioMartinez/aws/cloudtrail"
	"github.com/MauricioAntonioMartinez/aws/eventbridge"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/iam"
	"github.com/MauricioAntonioMartinez/aws/lambda"
	"github.com/MauricioAntonioMartinez/aws/sqs"
)

func main() {
	// fastergoding.Run()
	var err error

	cmd, err := cli.Parse(os.Args)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := *cmd

	if c.GetHelp() {
		c.Usage()
		return
	}

	l := helpers.NewLogger()
	svc := c.Name()

	switch svc {
	case "sqs":
		r := c.(*cli.SqsCmd)
		if err = sqs.Run(*r, l); err != nil {
			l.Fatal().Msg("Unable to start sqs service.")
		}
		break
	case "iam":
		r := c.(*cli.IamCmd)
		if err = iam.Run(*r, l); err != nil {
			l.Fatal().Msg("Unable to start iam service.")
		}
	case "lambda":
		r := c.(*cli.LambdaCmd)
		if err = lambda.Run(*r, l); err != nil {
			l.Fatal().Msg("Unable to start lambda service.")
		}
	case "eventbridge":
		r := c.(*cli.EventBridgeCmd)
		if err = eventbridge.Run(*r, l); err != nil {
			l.Fatal().Msg("Unable to start eventbridge service.")
		}
	case "cloudtrail":
		r := c.(*cli.CloudTrailCmd)
		if err = cloudtrail.Run(*r, l); err != nil {
			l.Fatal().Msg("Unable to start cloudtrail service.")
		}
	}

	if err != nil {
		fmt.Println(err)
	}
}
