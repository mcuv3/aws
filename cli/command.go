package cli

import (
	"flag"
)

type Command interface {
	init(args []string) error
	Name() string
	GetHelp() bool
	Usage()
}

type SqsCmd struct {
	PortGrpc string
	PortWeb  string
	Region   string
	DbUrl    string
	Secret   string
	Help     bool
	Fs       *flag.FlagSet
}

func newSqsCmd() *SqsCmd {

	cmd := &SqsCmd{
		Fs: flag.NewFlagSet("sqs", flag.ExitOnError),
	}

	cmd.Fs.StringVar(&cmd.PortGrpc, "port-grpc", "6001", "port to listen on grpc")
	cmd.Fs.StringVar(&cmd.PortWeb, "port-web", "7001", "port to listen on web browser")
	cmd.Fs.StringVar(&cmd.Region, "region", "us-east-1", "aws region [defualt=us-east-1]")
	cmd.Fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	cmd.Fs.StringVar(&cmd.Secret, "secret", "", "secret")
	cmd.Fs.BoolVar(&cmd.Help, "help", false, "show help of the command")

	return cmd
}

func (c *SqsCmd) GetHelp() bool {
	return c.Help
}

func (c *SqsCmd) Usage() {
	c.Fs.Usage()
}

func (c *SqsCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *SqsCmd) Name() string {
	return "sqs"
}

type LambdaCmd struct {
	DbUrl            string
	PortGrpc         string
	PortWeb          string
	EventsPerInvoque int
	Region           string
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

	cmd.Fs.StringVar(&cmd.PortGrpc, "port-grpc", "6002", "port to listen on gRPC")
	cmd.Fs.StringVar(&cmd.PortWeb, "port-web", "7002", "port to listen on web browser")
	cmd.Fs.IntVar(&cmd.EventsPerInvoque, "events-pre-invoque", 10, "number of concurrent events to process on a single lambda invocation")
	cmd.Fs.StringVar(&cmd.Region, "region", "us-east-1", "aws region")
	cmd.Fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	cmd.Fs.StringVar(&cmd.Secret, "secret", "", "jwt secret")
	cmd.Fs.IntVar(&cmd.Workers, "workers", 3, "amount of workers to process lambdas")
	cmd.Fs.StringVar(&cmd.ContainerRuntime, "runtime", "docker", "container runtime to use [docker|containered|crio]")
	cmd.Fs.BoolVar(&cmd.Help, "help", false, "show help of the command")

	return cmd
}

func (c *LambdaCmd) GetHelp() bool {
	return c.Help
}

func (c *LambdaCmd) Usage() {
	c.Fs.Usage()
}

func (c *LambdaCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *LambdaCmd) Name() string {
	return "lambda"
}

type IamCmd struct {
	PortGrpc         string
	PortWeb          string
	DbUrl            string
	Region           string
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

	cmd.Fs.StringVar(&cmd.PortGrpc, "port-grpc", "6000", "port to listen on gRPC")
	cmd.Fs.StringVar(&cmd.PortWeb, "port-web", "7000", "port to listen on web browser")
	cmd.Fs.StringVar(&cmd.Region, "region", "us-east-1", "aws region [defualt=us-east-1]")
	cmd.Fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	cmd.Fs.StringVar(&cmd.Secret, "secret", "", "secret")
	cmd.Fs.BoolVar(&cmd.Help, "help", false, "show help of the command")

	return cmd
}

func (c *IamCmd) GetHelp() bool {
	return c.Help
}

func (c *IamCmd) Usage() {
	c.Fs.Usage()
}

func (c *IamCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *IamCmd) Name() string {
	return "iam"
}
