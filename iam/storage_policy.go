package iam

import (
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"gorm.io/gorm"
)

type IamRepository struct {
	db  *gorm.DB
}




func (storage *IamRepository) CreatePolicy(in *aws.Policy) (* model.Policy, error) {
	
	p := model.Policy{Name: in.Name,Description: in.GetDescription(),
		Manifest: in.GetManifest(),Arn: "aws",AccountId: 231}

	tx := storage.db.Create(&p)
	return &p,tx.Error
	
}


func (storage *IamRepository) DeletePolicy(id uint32) error {

	return nil
}


func (storage *IamRepository) GetPolicy(id uint32) (*aws.Policy, error) {

	return nil, nil
}


func (storage *IamRepository) UpdatePolicy(updated *aws.Policy) (*aws.Policy, error) {

	return nil,nil
}
