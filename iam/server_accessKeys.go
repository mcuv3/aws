package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreateAccessKeys(ctx context.Context, req *aws.CreateAccessKeysRequest) (*aws.CreateAccessKeysResponse, error) {
	created, err := s.storage.CreateAccessKeys(req.AccessKeys)
	if err != nil {
		// s.Error("Error in CreateAccessKeys", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.CreateAccessKeysResponse{
		Created: created,
	}, nil
}

func (s *IAMService) GetAccessKeys(ctx context.Context, req *aws.GetAccessKeysRequest) (*aws.GetAccessKeysResponse, error) {
	ret, err := s.storage.GetAccessKeys(req.Id)
	if err != nil {
		// s.Error("Error in GetAccessKeys", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetAccessKeysResponse{
		AccessKeys: ret,
	}, nil
}

func (s *IAMService) DeleteAccessKeys(ctx context.Context, req *aws.DeleteAccessKeysRequest) (*aws.DeleteAccessKeysResponse, error) {
	err := s.storage.DeleteAccessKeys(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.DeleteAccessKeysResponse{}, nil
}
