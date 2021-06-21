package backend

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePolicy(ctx context.Context, req *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	created, err := s.storage.CreatePolicy(req.Policy)
	if err != nil {
		reportError("Error in CreatePolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &CreatePolicyResponse{
		Created: created,
	}, nil
}

func (s *Server) GetPolicy(ctx context.Context, req *GetPolicyRequest) (*GetPolicyResponse, error) {
	ret, err := s.storage.GetPolicy(req.Id)
	if err != nil {
		reportError("Error in GetPolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &GetPolicyResponse{
		Policy: ret,
	}, nil
}

func (s *Server) UpdatePolicy(ctx context.Context, req *UpdatePolicyRequest) (*UpdatePolicyResponse, error) {
	updated, err := s.storage.UpdatePolicy(req.Updated)
	if err != nil {
		reportError("Error in UpdatePolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &UpdatePolicyResponse{
		Result: updated,
	}, nil
}

func (s *Server) DeletePolicy(ctx context.Context, req *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	err := s.storage.DeletePolicy(req.Id)
	if err != nil {
		reportError("Error in DeletePolicy", err)
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &DeletePolicyResponse{}, nil
}
