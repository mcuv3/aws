package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreateGroup(ctx context.Context, req *aws.CreateGroupRequest) (*aws.CreateGroupResponse, error) {
	created, err := s.storage.CreateGroup(req.Group)
	if err != nil {
		reportError("Error in CreateGroup", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.CreateGroupResponse{
		Created: created,
	}, nil
}

func (s *IAMService) GetGroup(ctx context.Context, req *aws.GetGroupRequest) (*aws.GetGroupResponse, error) {
	ret, err := s.storage.GetGroup(req.Id)
	if err != nil {
		reportError("Error in GetGroup", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetGroupResponse{
		Group: ret,
	}, nil
}

func (s *IAMService) UpdateGroup(ctx context.Context, req *aws.UpdateGroupRequest) (*aws.UpdateGroupResponse, error) {
	updated, err := s.storage.UpdateGroup(req.Updated)
	if err != nil {
		reportError("Error in UpdateGroup", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.UpdateGroupResponse{
		Result: updated,
	}, nil
}

func (s *IAMService) DeleteGroup(ctx context.Context, req *aws.DeleteGroupRequest) (*aws.DeleteGroupResponse, error) {
	err := s.storage.DeleteGroup(req.Id)
	if err != nil {
		reportError("Error in DeleteGroup", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.DeleteGroupResponse{}, nil
}
