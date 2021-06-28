package iam

import (
	"context"
	"fmt"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreateUser(ctx context.Context, req *aws.CreateUserRequest) (*aws.CreateUserResponse, error) {
	// created, err := s.storage.CreateUser(req.User)

	// if err != nil {
	// 	return nil, status.Error(codes.Internal, "Internal error")
	// }

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot extract data from context")
	}

	fmt.Println("Context Data")
	fmt.Println(md)

	return &aws.CreateUserResponse{
		Created: &aws.User{},
	}, nil
}

func (s *IAMService) GetUser(ctx context.Context, req *aws.GetUserRequest) (*aws.GetUserResponse, error) {
	_, err := s.storage.FindUser("")
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetUserResponse{}, nil
}

func (s *IAMService) UpdateUser(ctx context.Context, req *aws.UpdateUserRequest) (*aws.UpdateUserResponse, error) {
	updated, err := s.storage.UpdateUser(req.Updated)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.UpdateUserResponse{
		Result: updated,
	}, nil
}

func (s *IAMService) DeleteUser(ctx context.Context, req *aws.DeleteUserRequest) (*aws.DeleteUserResponse, error) {
	err := s.storage.DeleteUser(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.DeleteUserResponse{}, nil
}
