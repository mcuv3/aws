package eventbridge

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
)

type EventBridgeDispatcherConfig eventbus.ConsumerConfig

type EventBridgeDispatcher struct {
	consumer *eventbus.Consumer
}

func newEventBridgeDispatcher(config EventBridgeDispatcherConfig) *EventBridgeDispatcher {
	dispatcher := &EventBridgeDispatcher{}
	dispatcher.consumer = eventbus.NewConsumer(eventbus.ConsumerConfig(config), dispatcher.onMessage)
	return dispatcher
}

func (d *EventBridgeDispatcher) Start() {
	d.consumer.Start()
}

func (d *EventBridgeDispatcher) onMessage(msg eventbus.Message) {
	fmt.Printf("OnMessage:  %s Offset: %d \n", msg.Value, msg.Offset)

}

func (d *EventBridgeDispatcher) Stop() {
	d.consumer.Close()
}
