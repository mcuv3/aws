package auth

import (
	"encoding/json"
	"strings"
)

// type Arn  = string this is an alias
type Arn string // this is a new type

type Effect string

type Action string

const (
	Allow = "Allow"
	Deny  = "Deny"
)

var (
	Services = []string{"iam", "sqs", "lambda", "ses"}
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

func (a Arn) GetService() []string {
	parts := strings.Split(string(a), ":")

	return parts
}

func in(arr []string, value string) bool {
	isIn := false
	for _, v := range arr {
		isIn = v == value
	}
	return isIn
}

func (a Arn) Validate() bool {

	parts := strings.Split(string(a), ":")

	if len(a) == 1 && a == "*" {
		return true
	}

	return len(parts) >= 6 && parts[0] == "arn" && parts[1] == "aws" && in(Services, parts[2])
}

func (a Action) Validate(service string, avaliableMethods []string) bool {
	parts := strings.Split(string(a), ":")

	if len(parts) != 2 && parts[0] != service {
		return false
	}

	return in(avaliableMethods, parts[1])

}
