package iam

import (
	"context"
	"errors"
	"fmt"
	"time"

	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)


type UserClaims struct { 
	jwt.StandardClaims
	Username string `json:"username"`
	Policy []*aws.Policy `json:"policy"`
}

type AuthInterceptor struct {
	Mannager *JWTMannger
	Issuer string
	l zerolog.Logger
}

type JWTMannger struct { 
	SecretKey string
	Duration time.Duration
}

func NewJWTMannager(SecretKey string ,Duration time.Duration) *JWTMannger{
	return &JWTMannger{SecretKey: SecretKey,Duration: Duration}
}

func NewAuthInterceptor(mannager *JWTMannger,issuer string) *AuthInterceptor{
	return &AuthInterceptor{Mannager: mannager,Issuer:issuer }
}



func (a *AuthInterceptor) Validate(token string) (*UserClaims,error) {

	tk,err := jwt.ParseWithClaims(token,UserClaims{},func (token *jwt.Token)(interface{},error){
		 if ok ,_ := token.Method.(*jwt.SigningMethodRSA) ; ok == nil {
			 a.l.Err(errors.New("Invalid token")).Msg("Invalid token was provided")
			 return nil,fmt.Errorf("Invalid token")
		 }
		return a.Mannager.SecretKey,nil
	})

	if err !=nil { 
		a.l.Err(err).Msg("Token with an invalid signature.")
		return nil,err
	}

	 claims , ok := tk.Claims.(*UserClaims)
	 if !ok { 
		a.l.Info().Str("token",tk.Raw).Msg("The token does not have a valid structure")
		return nil,errors.New("Not a valid structure")
	 }
	 return claims , nil
}


func (a *AuthInterceptor) GetToken(user aws.User) (string, error) { 
	claims := UserClaims{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(a.Mannager.Duration).Unix(),
		IssuedAt: time.Now().Unix(),
		Issuer: a.Issuer,
	},
	Username:  user.Name,
	Policy:  user.Policies,
	}

	token :=  jwt.NewWithClaims(jwt.SigningMethodRS256,&claims) // creates the token from the claims using the provided method

	return token.SignedString([]byte(a.Mannager.SecretKey)) // this signs the token
}


func (a *AuthInterceptor) Authorize(ctx context.Context,method string) error { 
    md,ok:=	metadata.FromIncomingContext(ctx)

	if !ok { 
		return status.Errorf(codes.InvalidArgument,"Cannot extract data from context")
	}
	fmt.Println("Context Data")
	fmt.Println(md["name"])
	return nil
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func (ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(interface{},error){

		a.l.Info().Bool("validating",true).Msg("Validating incoming request")

		if err := a.Authorize(ctx,info.FullMethod); err !=nil { 
			return nil,err
		}

		return handler(ctx,req)
	}
}