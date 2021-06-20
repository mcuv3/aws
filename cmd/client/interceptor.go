package main

import (
	"context"
	"log"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)



type AuthInterceptor  struct { 
	accessToken string
	authMethods map[string]string
	client *aws.IAMServiceClient
}

func NewAuthInterceptor(accessToken string,authMethods map[string]string)*AuthInterceptor { 
	return &AuthInterceptor{accessToken: accessToken,authMethods: authMethods}
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}


func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor CLIENT: %s", method)

		if interceptor.authMethods[method] != "" {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}