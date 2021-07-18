package auth

import (
	"encoding/json"
	"strings"
)

// type Arn  = string this is an alias

type Effect string

type Action string

type Service string 

type Region string

const (
	Allow = "Allow"
	Deny  = "Deny"
)

var (
	Services = []string{"iam", "sqs", "lambda", "ses"}
	GlobalServices = []string{"s3","cloudfront","organization"}
)

var (
	Svs = map[Service]string{
		IAM:"iam",
	}
)

const (
	IAM Service = "iam"
	SQS 		= "sqs"
	Lambda		= "lambda"
	SES			= "ses"
	CloudFront  = "cloudfront"
)

const (
	REGION_NONE Region = ""
	US_EAST_1  = "us-east-1"
)

type Statement struct {
	Sid       string   `json:"sid"`
	Principal []Arn    `json:"principal"`
	Effect    Effect   `json:"effect"`
	Actions   []Action `json:"action"`
	Resource  []Arn    `json:"resource"`
}

type ResourcePolicy struct {
	Version    string      `json:"version"`
	Statements []Statement `json:"statement"`
}


// TODO: logic to validate and create policies easier.

func ParsePolicy(p string) (*ResourcePolicy, error) {
	policy := ResourcePolicy{}
	err := json.Unmarshal([]byte(p), &policy)
	if err != nil {
		return nil, err
	}
	return &policy, nil
}

func (r *ResourcePolicy) CheckPermission(action Action, resources ...Arn) bool {

	statements := r.Statements

	isAllowed := false

	for _, s := range statements {
		for _, p := range s.Actions {
			if p == action {
				isAllowed = s.Effect == Allow
			}
		}
	}

	return isAllowed
}



func (a Action) Validate(service string, avaliableMethods []string) bool {
	parts := strings.Split(string(a), ":")

	if len(parts) != 2 && parts[0] != service {
		return false
	}

	return inArr(avaliableMethods, parts[1])

}
