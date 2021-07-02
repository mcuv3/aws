package iam

import (
	"context"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) CreateUser(ctx context.Context, req *aws.CreateUserRequest) (*aws.CreateUserResponse, error) {
	// created, err := s.storage.CreateUser(req.User)

	// if err != nil {
	// 	return nil, status.Error(codes.Internal, "Internal error")
	// }

	user ,err := s.auth.GetUserMetadata(ctx)

	if err != nil { 
		s.logger.Err(err).Msg(err.Error())
		return nil,s.Error(err,codes.Unauthenticated,"Unauthorized.")
	}
	
 	created ,err := s.storage.CreateUser(req.User,user.AccountId)

	 if err !=nil  { 
		 return nil,s.Error(err,codes.Internal,"Unable to create a iam user.")
	 }

	return &aws.CreateUserResponse{
		Created: &aws.User{Id: uint32(created.ID),Name: created.Name,Description: created.Description,Arn: created.Arn},
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
