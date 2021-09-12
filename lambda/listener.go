package lambda

import (
	"github.com/MauricioAntonioMartinez/aws/eventbus"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type lambdaListener struct {
	consumer *eventbus.Consumer
	logger   zerolog.Logger
}

func newLambdaListener() *lambdaListener {
	lis := &lambdaListener{
		logger: helpers.NewLogger(),
	}
	lis.consumer = eventbus.NewConsumerGroup(eventbus.ConsumerConfig{
		Identifier: "lambda",
		Topic:      eventbus.InvokeFunction,
		Verbose:    true,
		Brokers:    []string{"broker:29092"},
	}, lis.onMessage)
	return lis
}

func (lis *lambdaListener) onMessage(msg eventbus.Message) {
	msgEvent := aws.InvokeLambdaEvent{}
	if err := proto.Unmarshal(msg.Value, &msgEvent); err != nil {
		lis.logger.Err(err).Msg("Error unmarshal message")
		return
	}
	fn := model.Function{
		Arn: msgEvent.Arn,
	}

	lis.logger.Info().Msgf("Invoking lambda -> %s", fn.Arn)
	if err := repo.findFunction(&fn); err != nil {
		lis.logger.Err(err).Msg("Function not found.")
		return
	}

	if err := execution.Run(LambdaExecutionConfig{
		ID:      fn.Arn,
		image:   fn.Image,
		ram:     int64(fn.Memory),
		data:    msgEvent.Payload,
		handler: fn.Handler,
	}); err != nil {
		lis.logger.Err(err).Msg("Error running function")

	}
}

func (lis *lambdaListener) Start() {
	lis.consumer.Start()
}
