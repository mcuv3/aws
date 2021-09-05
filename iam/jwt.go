package iam

import (
	"time"

	"github.com/MauricioAntonioMartinez/aws/model"
	"github.com/form3tech-oss/jwt-go"
)

type JWTMannger struct {
	SecretKey string
	Duration  time.Duration
}

func (s *IAMService) GetTokenForUser(clm model.UserClaims) (string, error) {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.jwt.Duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Username:   clm.Username,
		IsRootUser: clm.IsRootUser,
		AccountId:  clm.AccountId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // creates the token from the claims using the provided method
	return token.SignedString([]byte(s.jwt.SecretKey))         // this signs the token
}
