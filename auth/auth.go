package auth

import (
	"context"
	"errors"

	"github.com/form3tech-oss/jwt-go"
	"google.golang.org/grpc/metadata"
)

type UserClaims struct {
	jwt.StandardClaims
	Username   string
	IsRootUser bool
	AccountId  string
}

type RootClaims struct {
	jwt.StandardClaims
	Email     string `json:"email"`
	AccountId string
}

type UserMetadata struct {
	Email     string
	AccountId string
}

func GetUserMetadata(ctx context.Context) (*UserMetadata, error) {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errors.New("unable to get the context")
	}

	accountId, okAccount := md["accountid"]
	email, okEmail := md["email"]

	if !okEmail || !okAccount {
		return nil, errors.New("incomplete metadata")
	}

	if len(accountId) == 0 || len(email) == 0 {
		return nil, errors.New("incomplete user metadata")
	}
	return &UserMetadata{Email: email[0], AccountId: accountId[0]}, nil

}
