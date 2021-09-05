package model

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/auth"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
)

var (
	AllowedLanguages = map[string]Language{
		"golang": {Name: "golang", Runtimes: []string{"16"}, Extension: ".go", Activator: "go run"},
		"nodejs": {Name: "node", Runtimes: []string{"14"}, Extension: ".js", Activator: "node"},
		"python": {Name: "python", Runtimes: []string{"3.7"}, Extension: ".py", Activator: "python3"},
	}
)

type Language struct {
	Name      string
	Runtimes  []string
	Extension string
	Activator string
}

func PolicyToModel(polices []*aws.Policy, accountId string) []Policy {
	ps := []Policy{}

	for _, v := range polices {
		arn, _ := auth.NewArn(auth.IAM, "", accountId, fmt.Sprintf("/policy/%s", v.Name))
		ps = append(ps, Policy{
			Name:        v.Name,
			Arn:         arn.String(),
			AccountId:   accountId,
			Manifest:    v.Manifest,
			Description: v.Description,
		})
	}

	return ps
}
