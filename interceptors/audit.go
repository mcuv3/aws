package interceptors

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
	_ "github.com/segmentio/kafka-go/snappy"
	"google.golang.org/grpc"
)

type AuditInterceptor struct {
	Verbose bool
	writer  *kafka.Writer
}

func NewAuditInterceptor(KafkaBroker, topic string, partition int) *AuditInterceptor {
	// _, err := kafka.DialLeader(context.Background(), "tcp", KafkaBroker, topic, 0)
	// if err != nil {
	// 	panic(err.Error())
	// }

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{KafkaBroker},
		Topic:   topic,
	})
	conn, _ := kafka.Dial("tcp", KafkaBroker)
	b, _ := conn.Brokers()

	fmt.Println(b)

	return &AuditInterceptor{
		writer:  w,
		Verbose: true,
	}
}

func (a *AuditInterceptor) Stop() {
	a.writer.Close()
}

func (a *AuditInterceptor) PublishEvent(key, value string, headers *[]protocol.Header) error {
	err := a.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		return err
	}
	fmt.Println("CREATED!!!")
	return nil
}

func (a *AuditInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if a.Verbose {
			log.Printf("[AUDIT] %s %s", info.FullMethod, req)
			err := a.PublishEvent("unarycall", info.FullMethod, nil)
			if err != nil {
				log.Printf("[AUDIT] %s", err)
			}
		}
		return handler(ctx, req)
	}
}

func (a *AuditInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if a.Verbose {
			log.Printf("[AUDIT] %s ", info.FullMethod)
			err := a.PublishEvent("stream-rpc-call", info.FullMethod, nil)
			if err != nil {
				log.Printf("[AUDIT] %s", err)
			}
		}
		return handler(srv, stream)
	}
}
