package eventbus

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

type Writer struct {
	kafka *kafka.Writer
}

type WriteConfig struct {
	Brokers []string
	Topic   string
}

func NewWriter(config WriteConfig) *Writer {
	kafka := kafka.NewWriter(kafka.WriterConfig{
		Brokers: config.Brokers,
		Topic:   config.Topic,
	})
	return &Writer{
		kafka: kafka,
	}
}

func (w *Writer) WriteMessage(key, value []byte, headers ...protocol.Header) error {
	return w.kafka.WriteMessages(context.Background(), kafka.Message{
		Key:     key,
		Value:   value,
		Headers: headers,
	})
}

func (w *Writer) Close() {
	w.kafka.Close()
}
