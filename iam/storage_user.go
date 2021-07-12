package iam

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/MauricioAntonioMartinez/aws/auth"
	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"golang.org/x/crypto/bcrypt"
)


func policyToGorm(polices []*aws.Policy,accountId string) []model.Policy {
	ps := []model.Policy{}

	
	for _,v:= range polices { 
		arn ,_ := auth.NewArn(auth.IAM,auth.REGION_NONE,accountId,fmt.Sprintf("/policy/%s",v.Name))
		 ps = append(ps, model.Policy{
			 Name: v.Name,
			 Arn: arn.String(),
		 })
	}

	return ps
}

func (storage *IamRepository) CreateUser(us aws.User,password,accountId string) (*model.User, error) {


	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}
	arn,_ := auth.NewArn(auth.IAM,auth.REGION_NONE,accountId,fmt.Sprintf("/user/%s",us.Name))


	user := model.User{
	Password: string(encrypted),
	Arn: arn.String(),	
	Polices: policyToGorm(us.Policies,accountId),
	Description: us.Description, Name: us.Name,
	 AccountId: accountId}

	res := storage.db.Omit("AccessKeys").Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (storage *IamRepository) CreateRootUser(email string, password string, confPass string) (*model.RootUser, error) {

	if password != confPass {
		return nil, errors.New("Passwords does not match")
	}
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())

	accountId := rand.Intn(9999999999-1000000000) + 1000000000

	us := model.RootUser{Password: string(encrypted), Email: email, AccountId: strconv.Itoa(accountId)}

	res := storage.db.Create(&us)

	if res.Error != nil {
		return nil, res.Error
	}

	return &us, nil

}

func (storage *IamRepository) DeleteUser(id uint32) error {

	return nil
}

func (storage *IamRepository) FindRootUser(query string, args ...interface{}) (*model.RootUser, error) {
	us := &model.RootUser{}

	tx := storage.db.Where(query, args).First(us)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return us, nil
}

func (storage *IamRepository) FindUser(query string, args ...interface{}) (*model.User, error) {
	us := &model.User{}
	tx := storage.db.Where(query, args).First(us)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return us, nil
}

func (storage *IamRepository) UpdateUser(updated *aws.User) (*aws.User, error) {

	return nil, nil
}
