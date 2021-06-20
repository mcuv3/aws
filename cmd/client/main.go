package main

import (
	"context"
	"fmt"
	"log"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc"
)

func main(){

	itc := NewAuthInterceptor("---")
	
	conn, err  := grpc.Dial("localhost:6002",itc.Unary())

	if err !=nil { 
		log.Fatal("Unable to connect with the server.")
	}
	defer conn.Close()

	client := aws.NewSQSServiceClient(conn)


	res,err := client.CreateQueue(context.Background(),&aws.CreateQueueRequest{Queue: &aws.Queue{Name: "MyQueue",AccountId: "123123"}})

	if err !=nil { 
		log.Fatalf("Unable to create queue due to %v.",err)
	}

	fmt.Println(res.Queue)

}