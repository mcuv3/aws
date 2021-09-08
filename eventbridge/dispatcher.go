package eventbridge

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

type EventBridgeDispatcherConfig eventbus.ConsumerConfig

type EventBridgeDispatcher struct {
	consumer *eventbus.Consumer
	logger   zerolog.Logger
}

func newEventBridgeDispatcher(config EventBridgeDispatcherConfig) *EventBridgeDispatcher {
	dispatcher := &EventBridgeDispatcher{
		logger: helpers.NewLogger(),
	}
	dispatcher.consumer = eventbus.NewConsumerGroup(eventbus.ConsumerConfig(config), dispatcher.onMessage)
	return dispatcher
}

func (d *EventBridgeDispatcher) Start() {
	d.logger.Info().Msg("Starting EventBridgeDispatcher ....")
	d.consumer.Start()
}

func (d *EventBridgeDispatcher) onMessage(msg eventbus.Message) {
	payload := aws.AuditEvent{}

	if err := proto.Unmarshal(msg.Value, &payload); err != nil {
		d.logger.Err(err).Msgf("Error unmarshal message")
		return
	}
	d.logger.Debug().Msgf("EventBridge: processing method -> %s", payload.Method)

	if payload.AccountId == "" {
		d.logger.Warn().Msg("AccountId is empty")
		return
	}

	se := model.ServiceEvent{
		Path: payload.Method,
	}
	if err := repo.findServiceEvent(&se); err != nil {
		d.logger.Warn().Msg("RPC method not supported")
		return
	}
	// FIXME: create a rule again and test with the old method
	rules := []model.Rule{}
	fmt.Println(se.ID)
	if err := repo.findRulesQuery("account_id = ? AND service_event_id = ?",
		&rules, payload.AccountId, se.ID); err != nil {
		d.logger.Err(err).Msgf("Error retreiving rules")
	}

	fmt.Println(rules)

	for _, rule := range rules {
		fmt.Println(rule)
	}

	// get the triggers for each rule and execute the integration code
	// ex. push a message to a topic corresponding to the action like lambda-invoke, sqs-send-msg

}

func (d *EventBridgeDispatcher) Stop() {
	d.consumer.Close()
}
