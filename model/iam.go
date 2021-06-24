package model

import (
	"time"

	"gorm.io/gorm"
)

type Policy struct {
	gorm.Model
	Name             string
	Arn				 string
	Description 	 string
	Manifest  		 string // the actual json policy
	AccountId        uint
}


type Role struct { 
	gorm.Model
	Arn				 string
	Name string
	Description string
	Polices []Policy `gorm:"foreignKey:PolicyRef"`
}

type UserIam struct { 
	gorm.Model
	Name string
	Arn				 string
	Description string
	Polices []Policy `gorm:"foreignKey:PolicyRef"`
	AccessKeys []AccessKey`gorm:"foreignKey:AccessKey"`
}


type AccessKey struct { 
	gorm.Model
	Arn			string
	CreatedAt time.Time 
	AccessKeyId string
}

type Group struct { 
	gorm.Model
	Arn				 string
	Name string
	Description string
	Users []User`gorm:"foreignKey:UserRef"`
}

