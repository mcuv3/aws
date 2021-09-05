package iam

import (
	"context"
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *IAMService) RootUserLogin(ctx context.Context, req *aws.RootUserLoginRequest) (*aws.LoginResponse, error) {

	email := req.GetEmail()
	password := req.GetPassword()

	us, err := s.storage.FindRootUser("email = ?", email)

	if err != nil {
		return nil, s.Error(err, codes.Internal, "Unable to find user")
	}

	isValid := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(password))

	if isValid != nil {
		return nil, s.Error(isValid, codes.PermissionDenied, "Incorrect Password")
	}

	token, err := s.GetTokenForUser(model.UserClaims{Username: us.Email, IsRootUser: true, AccountId: us.AccountId})

	if err != nil {
		return nil, s.Error(isValid, codes.Internal, "Unable to generate token")
	}

	return &aws.LoginResponse{Token: token}, nil
}

func (s *IAMService) UserLogin(ctx context.Context, req *aws.UserLoginRequest) (*aws.LoginResponse, error) {

	name := req.GetName()
	accountId := req.GetAccountId()
	password := req.GetPassword()

	us, err := s.storage.FindUser("name = ? AND account_id = ?", name, accountId)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Unable to find a user")
	}

	isValid := bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(password))

	if isValid != nil {
		return nil, status.Errorf(codes.PermissionDenied, fmt.Sprint("Incorrect password."))
	}

	token, err := s.GetTokenForUser(model.UserClaims{Username: us.Name, IsRootUser: false, AccountId: us.AccountId})

	if err != nil {
		s.logger.Err(err).Msg("Unable to generate token")
		return nil, status.Errorf(codes.Internal, "Unable to generate a token.")
	}

	return &aws.LoginResponse{Token: token}, nil
}

func (s *IAMService) SignUp(ctx context.Context, req *aws.SignUpRequest) (*aws.SignUpResponse, error) {

	email := req.GetEmail()
	pass := req.GetPassword()
	passConfirm := req.GetConfirmPassword()

	user, err := s.storage.CreateRootUser(email, pass, passConfirm)

	if err != nil {
		s.logger.Err(err).AnErr("Aws-User", err).Msg("Error createing an aws root user.")
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Unable to create a user reason: %v", err))
	}

	return &aws.SignUpResponse{Succeed: true, AccountId: user.AccountId}, nil
}
