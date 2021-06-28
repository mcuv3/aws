package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreateRole(ctx context.Context, req *aws.CreateRoleRequest) (*aws.CreateRoleResponse, error) {
	created, err := s.storage.CreateRole(req.Role)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.CreateRoleResponse{
		Created: created,
	}, nil
}

func (s *IAMService) GetRole(ctx context.Context, req *aws.GetRoleRequest) (*aws.GetRoleResponse, error) {
	ret, err := s.storage.GetRole(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetRoleResponse{
		Role: ret,
	}, nil
}

func (s *IAMService) UpdateRole(ctx context.Context, req *aws.UpdateRoleRequest) (*aws.UpdateRoleResponse, error) {
	updated, err := s.storage.UpdateRole(req.Updated)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.UpdateRoleResponse{
		Result: updated,
	}, nil
}

func (s *IAMService) DeleteRole(ctx context.Context, req *aws.DeleteRoleRequest) (*aws.DeleteRoleResponse, error) {
	err := s.storage.DeleteRole(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.DeleteRoleResponse{}, nil
}
