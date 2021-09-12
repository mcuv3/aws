package eventbus

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Writer struct {
	kafka *kafka.Writer
}

type WriteConfig struct {
	Brokers []string
	Topic   Topic
}

func NewWriter(config WriteConfig) *Writer {
	kafka := kafka.NewWriter(kafka.WriterConfig{
		Brokers: config.Brokers,
		Topic:   string(config.Topic),
	})
	return &Writer{
		kafka: kafka,
	}
}

func (w *Writer) WriteMessage(msg Message) error {
	return w.kafka.WriteMessages(context.Background(), kafka.Message(msg))
}

func (w *Writer) Close() {
	w.kafka.Close()
}
