package interceptors

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/MauricioAntonioMartinez/aws/eventbus"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go/protocol"
	_ "github.com/segmentio/kafka-go/snappy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuditInterceptor struct {
	verbose bool
	writer  *eventbus.Writer
}

type AuditInterceptorConfig struct {
	Brokers []string
	Topic   eventbus.Topic
	Verbose bool
}

func NewAuditInterceptor(config AuditInterceptorConfig) *AuditInterceptor {
	w := eventbus.NewWriter(eventbus.WriteConfig{
		Brokers: config.Brokers,
		Topic:   config.Topic,
	})

	return &AuditInterceptor{
		writer:  w,
		verbose: config.Verbose,
	}
}

func (a *AuditInterceptor) Stop() {
	a.writer.Close()
}

func (a *AuditInterceptor) publishEvent(key, value []byte, headers *[]protocol.Header) error {

	err := a.writer.WriteMessage(eventbus.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AuditInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		a.audit("unarycall", ctx, info.FullMethod, req)
		data, _ := json.Marshal(req)
		fmt.Println("json ->", string(data))
		return handler(ctx, req)
	}
}

func (a *AuditInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		a.audit("streamcall", stream.Context(), info.FullMethod, nil)
		return handler(srv, stream)
	}
}

func (a *AuditInterceptor) audit(key string, ctx context.Context, method string, req interface{}) {
	if a.verbose {
		log.Printf("[AUDIT] %s ", method)
	}
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	accountId := ""
	ac := md["accountid"]
	fmt.Println(md)
	if len(ac) > 0 {
		accountId = ac[0]
	}
	payload, _ := json.Marshal(req)
	event := aws.AuditEvent{
		AccountId: accountId,
		Method:    method,
		Payload:   string(payload),
	}
	if value, err := proto.Marshal(&event); err != nil {
		fmt.Printf("Error %v", err)
		return
	} else {
		err := a.publishEvent([]byte(key), value, nil)
		if err != nil && a.verbose {
			log.Printf("[AUDIT] %s", err)
		}
	}

}
