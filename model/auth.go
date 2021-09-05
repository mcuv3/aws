package model

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/form3tech-oss/jwt-go"
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

func (u *UserMetadata) String() string {
	var (
		v = &url.Values{}
		s = reflect.ValueOf(*u)
		t = reflect.TypeOf(*u)
	)

	for i := 0; i < s.NumField(); i++ {
		v.Add(t.Field(i).Name, fmt.Sprintf("%v", s.Field(i).Interface()))
	}
	return v.Encode()
}
