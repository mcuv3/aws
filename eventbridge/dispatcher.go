package eventbridge

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
	"github.com/MauricioAntonioMartinez/aws/helpers"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/proto"
)

var (
	typeToTopic = map[string]eventbus.Topic{
		"LambdaFunction": eventbus.InvokeFunction,
	}
)

type TriggerEvent struct {
	RuleID    uint      `json:"rule_id"`
	Type      string    `json:"type"`
	TargetArn *string   `json:"arn,omitempty"`
	AccountId string    `json:"account_id"`
	Cron      *string   `json:"cron,omitempty"`
	RuleArn   string    `json:"rule_arn"`
	CreatedAt time.Time `json:"created_at"`
}

type EventBridgeDispatcherConfig eventbus.ConsumerConfig

type EventBridgeDispatcher struct {
	consumer *eventbus.Consumer
	logger   zerolog.Logger
	writer   eventbus.Writer
}

func newEventBridgeDispatcher(config EventBridgeDispatcherConfig) *EventBridgeDispatcher {
	dispatcher := &EventBridgeDispatcher{
		logger: helpers.NewLogger(),
		writer: *eventbus.NewWriter(eventbus.WriteConfig{
			Brokers: config.Brokers,
		}),
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

	triggers, err := d.getTriggers(payload.AccountId, &se.ID)
	if err != nil {
		d.logger.Err(err).Msgf("Error getting triggers")
		return
	}

	for _, tg := range triggers {
		topic, err := d.getTopic(tg.Type)
		if err != nil {
			d.logger.Err(err).Msgf("Error getting topic")
			continue
		}

		payload, err := d.getPayload(tg)
		if err != nil {
			d.logger.Err(err).Msg("Error getting the payload")
			continue
		}

		go d.writer.WriteMessage(eventbus.Message{
			Topic: string(topic),
			Key:   []byte("event-bridge-actor"),
			Value: payload,
		})

	}
}

func (d *EventBridgeDispatcher) getTriggers(accountId string, serviceId *uint) ([]TriggerEvent, error) {
	rows, err := repo.findRulesWithTargets(accountId, serviceId)
	if err != nil {
		return nil, err
	}

	targets := []TriggerEvent{}
	for rows.Next() {
		var target TriggerEvent
		if err := rows.Scan(&target.RuleID, &target.Type,
			&target.TargetArn, &target.RuleArn, &target.Cron); err != nil {
			d.logger.Err(err).Msgf("Error scanning rule")
			return nil, fmt.Errorf("Error scaning rules and triggers")
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (d *EventBridgeDispatcher) getTopic(typeEvent string) (eventbus.Topic, error) {
	topic, ok := typeToTopic[typeEvent]
	if !ok {
		return eventbus.Topic(""), fmt.Errorf("Topic not found")
	}
	return topic, nil
}

func (d *EventBridgeDispatcher) getPayload(event TriggerEvent) ([]byte, error) {
	payloadBinary := []byte{}
	payload, err := json.MarshalIndent(event, "", "\t")
	if err != nil {
		return payloadBinary, err
	}

	switch event.Type {
	case "LambdaFunction":
		event := aws.InvokeLambdaEvent{
			Arn:       *event.TargetArn,
			AccountId: event.AccountId,
			Payload:   string(payload),
		}
		bin, err := proto.Marshal(&event)
		if err != nil {
			return payloadBinary, err
		}
		payloadBinary = bin
	case "SqsLambda":
		event := aws.SendMessageToQueueEvent{
			Arn:       *event.TargetArn,
			AccountId: event.AccountId,
			Payload:   string(payload),
		}
		bin, err := proto.Marshal(&event)
		if err != nil {
			return payloadBinary, err
		}
		payloadBinary = bin

	}

	return payloadBinary, nil
}

func (d *EventBridgeDispatcher) Stop() {
	d.consumer.Close()
}
