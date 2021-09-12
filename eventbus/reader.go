package eventbus

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Message = kafka.Message

type Reader interface {
}

type Consumer struct {
	id        string
	reader    *kafka.Reader
	config    ConsumerConfig
	onMessage func(Message)
}

type ConsumerConfig struct {
	Identifier string
	Verbose    bool
	Topic      Topic
	Brokers    []string
}

func NewConsumerGroup(config ConsumerConfig, onMessage func(Message)) *Consumer {
	return &Consumer{
		id:        config.Identifier,
		onMessage: onMessage,
		config:    config,
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: config.Brokers,
			Topic:   string(config.Topic),
			// Partition: 0,
			MinBytes: 10e3, // 10KB
			GroupID:  config.Identifier,
			MaxBytes: 10e6, // 10MB
		}),
	}
}

func (e *Consumer) Start() {
	if e.config.Verbose {
		fmt.Println("Start listenning for events consumer: ", e.id)
	}
	for {
		m, err := e.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		if e.config.Verbose {
			fmt.Printf("Reading from  %s offset %d key: %s value:%s ", e.id, m.Offset, m.Key, string(m.Value))
		}

		e.onMessage(m)
	}

	if err := e.reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
	fmt.Println("Listener disconnected")
}

func (e *Consumer) Close() {
	e.reader.Close()
}
