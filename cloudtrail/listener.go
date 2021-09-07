package main

import (
	"fmt"
	"log"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/golang/protobuf/proto"
)

func main() {

	fmt.Println("Start listening CLOUDTRAIL!!!")
	cloudTrailConsumer := eventbus.NewConsumerGroup(eventbus.ConsumerConfig{
		Identifier: "cloudtrail",
		Verbose:    true,
		Topic:      "audit",
		Brokers:    []string{"localhost:9092"},
	}, onMessage)

	cloudTrailConsumer.Start()

}

func onMessage(msg eventbus.Message) {

	event := aws.AuditEvent{}
	err := proto.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Fatal("Invalid message")
	}

	fmt.Printf("FROM CLOUDTRAIL OnMessage:  %v ", event)

}
