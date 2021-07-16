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




func (storage *IamRepository) CreateUser(us aws.User,password,accountId string) (*model.User, error) {

	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	arn,_ := auth.NewArn(auth.IAM,auth.REGION_NONE,accountId,fmt.Sprintf("/user/%s",us.Name))

	user := model.User{
	Password: string(encrypted),
	Arn: arn.String(),	
	Polices: model.PolicyToModel(us.Policies,accountId),
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

func (storage *IamRepository) DeleteUser(id uint32,accountId string) error {

	tx  := storage.db.Where("account_id = ? AND id = ?",accountId,id)

	if tx.Error != nil { 
		return tx.Error
	}
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
	us := model.User{}

	tx := storage.db.Where(query, args...).First(&us)

	
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &us, nil
} 

func (storage *IamRepository) UpdateUser(accountId string , updated *aws.User) (*model.User, error) {

	us,err  := storage.FindUser("account_id = ? AND id = ?",accountId,updated.Id)

	if err !=nil { 
		return nil,err
	}

	arn := us.Arn

	if us.Name != updated.Name { 
		newArn,err := auth.NewArn(auth.IAM,auth.REGION_NONE,accountId,fmt.Sprintf("/user/%s",us.Name))
		if err !=nil { return nil, err }
		arn = newArn.String()
	}


	user := model.User{
		Arn: arn,	
		Description: us.Description,
		Name: us.Name,
		AccountId: accountId}

	res := storage.db.Omit("AccessKeys").Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil

}
