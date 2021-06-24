package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreateUser(ctx context.Context, req *aws.CreateUserRequest) (*aws.CreateUserResponse, error) {
	created, err := s.storage.CreateUser(req.User)
	if err != nil {
		reportError("Error in CreateUser", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.CreateUserResponse{
		Created: created,
	}, nil
}

func (s *IAMService) GetUser(ctx context.Context, req *aws.GetUserRequest) (*aws.GetUserResponse, error) {
	ret, err := s.storage.GetUser(req.Id)
	if err != nil {
		reportError("Error in GetUser", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetUserResponse{
		User: ret,
	}, nil
}

func (s *IAMService) UpdateUser(ctx context.Context, req *aws.UpdateUserRequest) (*aws.UpdateUserResponse, error) {
	updated, err := s.storage.UpdateUser(req.Updated)
	if err != nil {
		reportError("Error in UpdateUser", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.UpdateUserResponse{
		Result: updated,
	}, nil
}

func (s *IAMService) DeleteUser(ctx context.Context, req *aws.DeleteUserRequest) (*aws.DeleteUserResponse, error) {
	err := s.storage.DeleteUser(req.Id)
	if err != nil {
		reportError("Error in DeleteUser", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.DeleteUserResponse{}, nil
}
