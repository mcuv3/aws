package iam

import (
	"context"
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/auth"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *IAMService) CreateUser(ctx context.Context, req *aws.CreateUserRequest) (*aws.CreateUserResponse, error) {

	user ,err := s.auth.GetUserMetadata(ctx)

	if err != nil { 
		s.logger.Err(err).Msg(err.Error())
		return nil,s.Error(err,codes.Unauthenticated,"Unauthorized.")
	}

	arn ,err := auth.NewArn(auth.IAM,"",user.AccountId,fmt.Sprintf("/user/%s",req.Name))

	if err !=nil { 
		return nil, s.Error(err,codes.Internal,"Unable to generate arn.")
	}

 	created ,err := s.storage.CreateUser(aws.User{
		Name: req.GetName(),
		Description: req.GetDescription(),
		Arn: arn.String(),
	 },req.GetPassword(),user.AccountId)

	 if err !=nil  { 
		 return nil,s.Error(err,codes.Internal,"Unable to create a iam user.")
	 }


	return &aws.CreateUserResponse{
		Created: &aws.User{Id: uint32(created.ID),
		Name: created.Name,
		Description: created.Description,
		Arn: arn.String(),
	},
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

func (s *IAMService) DeleteUser(ctx context.Context, req *aws.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteUser(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &emptypb.Empty{}, nil
}
