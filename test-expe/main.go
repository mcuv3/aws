package main

import (
	"fmt"
	"log"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	res := aws.CreateUserRequest{
		Name:        "Test",
		Description: "New",
		Password:    "123",
		Polices:     []*aws.Policy{},
	}
	value, err := proto.Marshal(&res)
	// cns:= eventbus.NewConsumerGroup(eventbus.ConsumerConfig{
	// 	Identifier: "test",
	// 	Verbose: true,
	// 	Topic: "audit",
	// 	Brokers: []string{"localhost:9092"},
	// },onMessage)
	// go cns.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	newRes := &aws.CreateUserRequest{}
	err = proto.Unmarshal(value, newRes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newRes)

}

func writeMessage(msg []byte) {

}

func onMessage(m eventbus.Message) {
	fmt.Println(m.Value)
}
