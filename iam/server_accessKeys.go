package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *IAMService) CreateAccessKey(ctx context.Context, req *aws.CreateAccessKeyRequest) (*aws.CreateAccessKeysResponse, error) {
	// created, err := s.storage.CreateAccessKeys()
	// if err != nil {
	// 	// s.Error("Error in CreateAccessKeys", err)
	// 	return nil, status.Error(codes.Internal, "Internal error")
	// }
	// return &aws.CreateAccessKeysResponse{
	// 	Created: created,
	// }, nil
	return nil,nil
}

func (s *IAMService) GetAccessKeys(ctx context.Context, req *aws.GetAccessKeysRequest) (*aws.GetAccessKeysResponse, error) {
	ret, err := s.storage.GetAccessKeys(req.GetUserId())
	if err != nil {
		// s.Error("Error in GetAccessKeys", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetAccessKeysResponse{
		AccessKeys: ret,
	}, nil
}

func (s *IAMService) DeleteAccessKeys(ctx context.Context, req *aws.DeleteAccessKeysRequest)  (*emptypb.Empty,error) {
	err := s.storage.DeleteAccessKeys(req.AccessKeyId)
	if err != nil {
		return  nil,status.Error(codes.Internal, "Internal error")
	}
	return &emptypb.Empty{},nil
}
