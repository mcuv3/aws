package model

import (
	"time"

	"gorm.io/gorm"
)

type AwsUser struct { 
	gorm.Model 
	AccountId string `gorm:"foreignkey;unique"`
	Email string
	Password string
	Polices []Policy `gorm:"foreignkey:AccountId"`
	Users []UserIam `gorm:"foreignkey:AccountId"`
	// Groups []Group `gorm:"foreignKey:AccountId"`
}

type Policy struct {
	gorm.Model
	Name             string
	Arn				 string
	Description 	 string
	Manifest  		 string // the actual json policy
	AccountId        string 
	Users []*UserIam   `gorm:"many2many:policy_user"` 
	Roles []*Role   `gorm:"many2many:policy_role"`
}


type Role struct { 
	gorm.Model
	Arn				string
	Name 			string
	Description 	string
	Polices []*Policy `gorm:"many2many:policy_role"`
	AccountId 		string
}

type Group struct { 
	gorm.Model
	Arn				 string
	Name string
	Description string
	Users []UserIam `gorm:"foreignKey:GroupID"`
	AccountId string 
}

type UserIam struct { 
	gorm.Model
	Name string
	Arn				 string
	Description string
	Polices []*Policy `gorm:"many2many:policy_user"`
	AccessKeys []AccessKey`gorm:"foreignKey:UserIamID"`
	AccountId string
	GroupID uint
}


type AccessKey struct { 
	gorm.Model
	Arn			string
	CreatedAt time.Time 
	AccessKeyId string
	UserIamID uint
}


