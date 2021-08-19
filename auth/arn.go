package auth

import (
	"errors"
	"fmt"
	"strings"
)

type Arn struct { 
	Region string
	Service
	AccountId string
	ResourceId string
}


func (Arn) validateArn(a string) bool {

	parts := strings.Split(a, ":")

	if len(a) == 1 && a == "*" {
		return true
	}


	return len(parts) >= 6 && parts[0] == "arn" && parts[1] == "aws" && inArr(Services, parts[2])
}

func (a Arn) String() string { 
	return fmt.Sprintf("arn:aws:%s:%s:%s:%s",a.Service,a.Region,a.AccountId,a.ResourceId)
}

func NewArn(s Service,r string,accountId string,ResourceId string) (*Arn,error) {
	arn := &Arn{
		Region: r,
		Service: s,
		AccountId: accountId,
		ResourceId: ResourceId,
	}

	fmt.Println(arn.String())
	isValid := arn.validateArn(arn.String())

	if !isValid { 
		return nil,errors.New("Invalid arn")
	}
	return arn,nil

}
