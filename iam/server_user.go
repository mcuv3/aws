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
		Policies: req.GetPolices(),
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

	us ,err := s.auth.GetUserMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	} 
 
	id := req.GetId()
	user , err := s.storage.FindUser("id = ? AND account_id = ?",id,us.AccountId)
	
	groupId := user.GroupID

	fetchUser :=  &aws.User{
		Id: uint32(user.ID),
		Name: user.Name,
		Description: user.Description,
		Arn: user.Arn,
	}

	if groupId !=nil {
		fetchUser.GroupId = uint32(*groupId)
	}

	return &aws.GetUserResponse{
		User:fetchUser,
	}, nil
}


func (s *IAMService) UpdateUser(ctx context.Context, req *aws.UpdateUserRequest) (*aws.UpdateUserResponse, error) {
	us ,err := s.auth.GetUserMetadata(ctx)

	if err !=nil { 
		return nil, s.Error(err,codes.Unauthenticated, "Unauthenticated!")
	}

	updated, err := s.storage.UpdateUser(us.AccountId,req.Updated)

	if err != nil {
		return nil, status.Error(codes.Internal, "Couldn't update user.")
	}

	return &aws.UpdateUserResponse{
		Result: &aws.User{
			Id: uint32(updated.ID),
			Name: updated.Name,
			Description: updated.Description,
			Arn: updated.Arn,
			GroupId: uint32(*updated.GroupID),
		},
	}, nil
}

func (s *IAMService) DeleteUser(ctx context.Context, req *aws.DeleteUserRequest) (*emptypb.Empty, error) {

	us ,err := s.auth.GetUserMetadata(ctx)

	if err !=nil { 
		return nil, s.Error(err,codes.Unauthenticated, "Unauthenticated!")
	}


	err = s.storage.DeleteUser(req.Id,us.AccountId)
	if err != nil {
		return nil, s.Error(err,codes.FailedPrecondition, "Unable to delete user.")
	}

	return &emptypb.Empty{}, nil
}
