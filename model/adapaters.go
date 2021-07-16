package model

import (
	"fmt"

	"github.com/MauricioAntonioMartinez/aws/auth"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
)


func PolicyToModel(polices []*aws.Policy,accountId string) []Policy {
	ps := []Policy{}
	
	for _,v:= range polices { 
		arn ,_ := auth.NewArn(auth.IAM,auth.REGION_NONE,accountId,fmt.Sprintf("/policy/%s",v.Name))
		 ps = append(ps, Policy{
			 Name: v.Name,
			 Arn: arn.String(),
			 AccountId: accountId,
			 Manifest: v.Manifest,
			 Description: v.Description,
		 })
	}

	return ps
}