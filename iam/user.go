package iam

import (
	"context"
	"fmt"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)


func(*IAMService) Login(ctx context.Context,req *aws.LoginRequest)(*aws.LoginResponse ,error){

	return nil,nil
}

func(s *IAMService) SignUp(ctx context.Context,req *aws.SignUpRequest)(*aws.SignUpResponse,error){

	email := req.GetEmail()
	pass := req.GetPassword()
	passConfirm := req.GetConfirmPassword()

	user ,err :=  s.storage.CreateAwsUser(email,pass,passConfirm)

	if err !=nil { 	
		s.l.Err(err).AnErr("Aws-User",err).Msg("Error createing an aws root user.")
		return nil,grpc.Errorf(codes.Internal,fmt.Sprintf("Unable to create a user reason: %v",err))
	}


	return &aws.SignUpResponse{Succeed: true,AccountId: user.AccountId },nil
}