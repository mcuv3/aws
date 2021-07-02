package iam

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/MauricioAntonioMartinez/aws/model"
	aws "github.com/MauricioAntonioMartinez/aws/proto"
	"golang.org/x/crypto/bcrypt"
)

func (storage *IamRepository) CreateUser(us *aws.User,accountId string) (*model.UserIam, error) {


	encrypted, err := bcrypt.GenerateFromPassword([]byte(us.Password), 10)

	if err != nil {
		return nil, err
	}


	user := model.UserIam{Password: string(encrypted),Arn: "",Description: us.Description,Polices: []*model.Policy{{Name: "FullAccess",AccountId: accountId,Arn: "", Description: "Test"}}, Name: us.Name, AccountId: accountId}

	res := storage.db.Create(&us)

	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

func (storage *IamRepository) CreateRootUser(email string, password string, confPass string) (*model.AwsUser, error) {

	if password != confPass {
		return nil, errors.New("Passwords does not match")
	}
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())

	accountId := rand.Intn(9999999999-1000000000) + 1000000000

	us := model.AwsUser{Password: string(encrypted), Email: email, AccountId: strconv.Itoa(accountId)}

	res := storage.db.Create(&us)

	if res.Error != nil {
		return nil, res.Error
	}

	return &us, nil

}

func (storage *IamRepository) DeleteUser(id uint32) error {

	return nil
}

func (storage *IamRepository) FindUser(query string, args ...interface{}) (*model.UserIam, error) {
	us := &model.UserIam{}

	tx := storage.db.Where(query, args).First(us)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return us, nil
}

func (storage *IamRepository) FindAwsUser(query, args string) (*model.AwsUser, error) {
	us := &model.AwsUser{}
	tx := storage.db.Where(query, args).First(us)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return us, nil
}

func (storage *IamRepository) UpdateUser(updated *aws.User) (*aws.User, error) {

	return nil, nil
}
