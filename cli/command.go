package cli

import (
	"flag"
)

type Command interface {
	init(args []string) error
	Name() string
}

type SqsCmd struct {
	Port   string
	DbUrl  string
	Secret string
	Help   bool
	Fs     *flag.FlagSet
}

func newSqsCmd() *SqsCmd {

	cmd := &SqsCmd{
		Fs: flag.NewFlagSet("sqs", flag.ExitOnError),
	}

	cmd.Fs.StringVar(&cmd.Port, "port", "8080", "port to listen on")
	cmd.Fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	cmd.Fs.StringVar(&cmd.Secret, "secret", "", "secret")
	cmd.Fs.BoolVar(&cmd.Help, "help", false, "show help of the command")

	return cmd
}

func (c *SqsCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *SqsCmd) Name() string {
	return "sqs"
}

type LambdaCmd struct {
	Port             string
	DbUrl            string
	Secret           string
	Workers          int
	ContainerRuntime string
	Help             bool
	Fs               *flag.FlagSet
}

func newLambdaCmd() *LambdaCmd {

	cmd := &LambdaCmd{
		Fs: flag.NewFlagSet("lambda", flag.ExitOnError),
	}

	cmd.Fs.StringVar(&cmd.Port, "port", "8080", "port to listen on")
	cmd.Fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	cmd.Fs.StringVar(&cmd.Secret, "secret", "", "secret")
	cmd.Fs.IntVar(&cmd.Workers, "workers", 3, "amount of workers to process lambdas")
	cmd.Fs.StringVar(&cmd.ContainerRuntime, "runtime", "docker", "container runtime to use [docker|containered|crio]")
	cmd.Fs.BoolVar(&cmd.Help, "help", false, "show help of the command")

	return cmd
}

func (c *LambdaCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *LambdaCmd) Name() string {
	return "lambda"
}

type IamCmd struct {
	Port             string
	DbUrl            string
	Secret           string
	Workers          int
	ContainerRuntime string
	Help             bool
	Fs               *flag.FlagSet
}

func newIamCmd() *IamCmd {

	cmd := &IamCmd{
		Fs: flag.NewFlagSet("iam", flag.ExitOnError),
	}

	cmd.Fs.StringVar(&cmd.Port, "port", "6000", "port to listen on")
	cmd.Fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	cmd.Fs.StringVar(&cmd.Secret, "secret", "", "secret")
	cmd.Fs.BoolVar(&cmd.Help, "help", false, "show help of the command")

	return cmd
}

func (c *IamCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *IamCmd) Name() string {
	return "iam"
}
