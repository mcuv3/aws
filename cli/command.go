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
	*CommonFlags
	Fs *flag.FlagSet
}

type CommonFlags struct {
	Fs         *flag.FlagSet
	PortGrpc   string
	PortWeb    string
	EnableWeb  bool
	EnableGrpc bool
	Region     string
	Origin     string
	Brokers    string
	DbUrl      string
	Secret     string
	Help       bool
}

type commonFlagsDefault struct {
	PortWeb  string
	PortGrpc string
}

func registerCommonFlags(fs *flag.FlagSet, def commonFlagsDefault) *CommonFlags {
	cmd := CommonFlags{}
	fs.StringVar(&cmd.PortGrpc, "port-grpc", def.PortGrpc, "port to listen on grpc")
	fs.StringVar(&cmd.PortWeb, "port-web", def.PortWeb, "port to listen on web browser")
	fs.BoolVar(&cmd.EnableGrpc, "enable-grpc", true, "port to listen on web browser")
	fs.BoolVar(&cmd.EnableWeb, "enable-web", true, "port to listen on web browser")
	fs.StringVar(&cmd.Region, "region", "us-east-1", "aws region")
	fs.StringVar(&cmd.Origin, "origin", "http://localhost:3000", "allowed origin.")
	fs.StringVar(&cmd.Brokers, "brokers", "broker:29092", "kafka brokers. brokers=broker1,broker2")
	fs.StringVar(&cmd.DbUrl, "db", "", "database url")
	fs.StringVar(&cmd.Secret, "secret", "", "jwt secret required")
	fs.BoolVar(&cmd.Help, "help", false, "show help of the command")
	return &cmd
}

func newSqsCmd() *SqsCmd {
	fs := flag.NewFlagSet("sqs", flag.ExitOnError)
	common := registerCommonFlags(fs, commonFlagsDefault{
		PortGrpc: "6001",
		PortWeb:  "7001",
	})
	cmd := &SqsCmd{
		Fs:          fs,
		CommonFlags: common,
	}

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
	*CommonFlags
	EventsPerInvoke  int
	ContainerRuntime string
	Workers          int
	Fs               *flag.FlagSet
}

func newLambdaCmd() *LambdaCmd {
	fs := flag.NewFlagSet("lambda", flag.ExitOnError)
	common := registerCommonFlags(fs, commonFlagsDefault{
		PortGrpc: "6002",
		PortWeb:  "7002",
	})
	cmd := &LambdaCmd{
		Fs:          fs,
		CommonFlags: common,
	}
	cmd.Fs.IntVar(&cmd.EventsPerInvoke, "events-per-invoke", 10, "number of concurrent events to process on a single lambda invocation")
	cmd.Fs.IntVar(&cmd.Workers, "workers", 3, "amount of workers to process lambdas")
	cmd.Fs.StringVar(&cmd.ContainerRuntime, "runtime", "docker", "container runtime to use [docker|containered|crio]")
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
	*CommonFlags
	Fs *flag.FlagSet
}

func newIamCmd() *IamCmd {
	fs := flag.NewFlagSet("iam", flag.ExitOnError)
	common := registerCommonFlags(fs, commonFlagsDefault{
		PortGrpc: "6000",
		PortWeb:  "7000",
	})
	cmd := &IamCmd{
		Fs:          fs,
		CommonFlags: common,
	}
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

type EventBridgeCmd struct {
	*CommonFlags
	Fs *flag.FlagSet
}

func newEventBridgeCmd() *EventBridgeCmd {
	fs := flag.NewFlagSet("eventbridge", flag.ExitOnError)
	common := registerCommonFlags(fs, commonFlagsDefault{
		PortGrpc: "6003",
		PortWeb:  "7003",
	})
	cmd := &EventBridgeCmd{
		Fs:          fs,
		CommonFlags: common,
	}

	return cmd
}

func (c *EventBridgeCmd) GetHelp() bool {
	return c.Help
}

func (c *EventBridgeCmd) Usage() {
	c.Fs.Usage()
}

func (c *EventBridgeCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *EventBridgeCmd) Name() string {
	return "eventbridge"
}

type CloudTrailCmd struct {
	*CommonFlags
	Fs *flag.FlagSet
}

func newCloudTrailCmd() *CloudTrailCmd {
	fs := flag.NewFlagSet("cloudtrail", flag.ExitOnError)
	common := registerCommonFlags(fs, commonFlagsDefault{
		PortGrpc: "6004",
		PortWeb:  "7004",
	})
	cmd := &CloudTrailCmd{
		Fs:          fs,
		CommonFlags: common,
	}

	return cmd
}

func (c *CloudTrailCmd) GetHelp() bool {
	return c.Help
}

func (c *CloudTrailCmd) Usage() {
	c.Fs.Usage()
}

func (c *CloudTrailCmd) init(args []string) error {
	return c.Fs.Parse(args)
}

func (c *CloudTrailCmd) Name() string {
	return "cloudtrail"
}
