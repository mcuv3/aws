package cloudtrail

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type CloudTrailListener struct {
	consumer *eventbus.Consumer
	logger   zerolog.Logger
}

func newCloudTrailListener(conf eventbus.ConsumerConfig) *CloudTrailListener {
	ct := CloudTrailListener{
		logger: helpers.NewLogger(),
	}
	cs := eventbus.NewConsumerGroup(conf, ct.onMessage)
	ct.consumer = cs
	return &ct
}

func (l *CloudTrailListener) onMessage(msg eventbus.Message) {
	var event aws.AuditEvent
	if err := proto.Unmarshal(msg.Value, &event); err != nil {
		l.logger.Err(err).Msg(" Error unmarshal message")
		return
	}

	if event.AccountId == "" {
		l.logger.Warn().Msg("AccountId is empty")
		return
	}
	var trails []model.Trail
	if err := repo.getTrails(model.Trail{
		AccountId: event.AccountId,
	}, &trails); err != nil {
		l.logger.Err(err).Msg("Error getting the trails for user.")
		return
	}

	for _, t := range trails {
		l.logger.Info().Msgf("Trail: %s", t.Name)
	}

	fmt.Println("Cloud trail", event.String())
}

func (l *CloudTrailListener) Start() {
	l.consumer.Start()
}
