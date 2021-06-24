package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreatePolicy(ctx context.Context, req *aws.CreatePolicyRequest) (*aws.CreatePolicyResponse, error) {
	_, err := s.storage.CreatePolicy(req.Policy)
	if err != nil {
		reportError("Error in CreatePolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.CreatePolicyResponse{
		// Created: created,
	}, nil
}

func (s *IAMService) GetPolicy(ctx context.Context, req *aws.GetPolicyRequest) (*aws.GetPolicyResponse, error) {
	ret, err := s.storage.GetPolicy(req.Id)
	if err != nil {
		reportError("Error in GetPolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.GetPolicyResponse{
		Policy: ret,
	}, nil
}

func (s *IAMService) UpdatePolicy(ctx context.Context, req *aws.UpdatePolicyRequest) (*aws.UpdatePolicyResponse, error) {
	updated, err := s.storage.UpdatePolicy(req.Updated)
	if err != nil {
		reportError("Error in UpdatePolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &aws.UpdatePolicyResponse{
		Result: updated,
	}, nil
}

func (s *IAMService) DeletePolicy(ctx context.Context, req *aws.DeletePolicyRequest) (*aws.DeletePolicyResponse, error) {
	err := s.storage.DeletePolicy(req.Id)
	if err != nil {
		reportError("Error in DeletePolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &aws.DeletePolicyResponse{}, nil
}
