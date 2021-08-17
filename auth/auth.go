package auth

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	Email string
	AccountId string
}



type AuthInterceptor struct {
	Mannager      *JWTMannger
	Issuer        string
	Logger        zerolog.Logger
	PublicMethods []string
	ServerPrefix string
}

type JWTMannger struct {
	SecretKey string
	Duration  time.Duration
}

func NewJWTMannager(SecretKey string, Duration time.Duration) *JWTMannger {
	return &JWTMannger{SecretKey: SecretKey, Duration: Duration}
}

func NewAuthInterceptor(mannager *JWTMannger, issuer string) *AuthInterceptor {
	return &AuthInterceptor{Mannager: mannager, Issuer: issuer}
}




func (a *AuthInterceptor) GetUserMetadata(ctx context.Context) (*UserMetadata, error) {

	md,ok:= metadata.FromIncomingContext(ctx)

	if !ok { 
		return nil,errors.New("Unable to get the context")
	}

	accountId,okAccount :=  md["AccountId"]
	email,okEmail := md["Email"] 

	if !okEmail || !okAccount {
		return nil,errors.New("Incomplete metadata")
	}


	if len(accountId) == 0 || len(email) == 0 { 
		return nil,errors.New("Incomplete user metadata.")
	}
	return &UserMetadata{Email: email[0],AccountId: accountId[0]},nil


}

func (a *AuthInterceptor) Validate(token string) (*UserClaims, error) {

	tk, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if ok, _ := token.Method.(*jwt.SigningMethodHMAC); ok == nil {
			a.Logger.Err(errors.New("Invalid token")).Msg("Invalid token was provided")
			return nil, fmt.Errorf("Invalid token")
		}
		return []byte(a.Mannager.SecretKey), nil
	})

	if err != nil {
		a.Logger.Err(err).Msg("Token with an invalid signature.")
		return nil, err
	}

	claims, ok := tk.Claims.(*UserClaims)

	if !ok {
		a.Logger.Info().Str("token", tk.Raw).Msg("The token does not have a valid structure")
		return nil, errors.New("Not a valid structure")
	}

	return claims, nil
}

func (a *AuthInterceptor) GetToken(clm UserClaims) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(a.Mannager.Duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Username:   clm.Username,
		IsRootUser: clm.IsRootUser,
		AccountId:  clm.AccountId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // creates the token from the claims using the provided method
	return token.SignedString([]byte(a.Mannager.SecretKey))    // this signs the token
}

func (a *AuthInterceptor) Authorize(ctx context.Context, method string) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot extract data from context")
	}

	tk := md.Get("authorization")

	if len(tk) == 0 {
		return nil, errors.New("Authorization token required.")
	}

	token := tk[0]

	claims, err := a.Validate(token)

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		a.Logger.Info().Bool("validating", true).Msg("Validating incoming request")

		var claims *UserClaims
		var err error
		skipAuth := false


		for _,method := range a.PublicMethods { 
			if fmt.Sprintf("%s%s",a.ServerPrefix,method) == info.FullMethod {
				skipAuth = true
			}
		}

		if !skipAuth {
			if claims, err = a.Authorize(ctx, info.FullMethod); err != nil {
				return nil, err
			}
			ctxWithAuth := metadata.NewIncomingContext(ctx,metadata.MD{
				"AccountId":[]string{ claims.AccountId},
				"Email":[]string{ claims.Username},
			}) 
			return handler(ctxWithAuth, req)
		} 

		return handler(ctx,req)


	}
}

func (a *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
    return func(
        srv interface{},
        stream grpc.ServerStream,
        info *grpc.StreamServerInfo,
        handler grpc.StreamHandler,
    ) error {

		if info.IsServerStream {
			return handler(srv, stream)
		}


		a.Logger.Info().Bool("validating", true).Msg("Validating incoming request")

		var claims *UserClaims
		var err error
		skipAuth := false


		for _,method := range a.PublicMethods { 
			if fmt.Sprintf("%s%s",a.ServerPrefix,method) == info.FullMethod {
				skipAuth = true
			}
		}
 
		if !skipAuth {
			if claims, err = a.Authorize(stream.Context(), info.FullMethod); err != nil {
				return  err
			}

			
			stream.SetTrailer(metadata.MD{
				"AccountId":[]string{claims.AccountId},
				// "Email":[]string{claims.Username},
			}) 
			stream.SetHeader(metadata.MD{
				"AccountId":[]string{claims.AccountId},
				// "Email":[]string{claims.Username},
			})
		
			metadata.AppendToOutgoingContext(stream.Context(),"TEST","TEST")
	stream.Context()
			// fmt.Println(claims)
		} 

		stream.SetHeader(metadata.MD{
			"Something":[]string{"testin"},
		})


		return handler(srv,stream)
    }
}


func (u *UserMetadata) String() string { 
	var(
		v = &url.Values{
			
		}
		s = reflect.ValueOf(*u)
		t = reflect.TypeOf(*u)
	)

	for i := 0 ; i < s.NumField(); i++ { 
		v.Add(t.Field(i).Name,fmt.Sprintf("%v",s.Field(i).Interface()))
	}
	return v.Encode()
}