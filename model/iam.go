package model

import (
	"time"

	"gorm.io/gorm"
)

type RootUser struct {
	AccountId string `gorm:"primaryKey"`
	Email     string
	Password  string
	Polices   []Policy  `gorm:"foreignkey:AccountId"`
	Users     []User `gorm:"foreignkey:AccountId"`
	Groups    []Group `gorm:"foreignKey:AccountId"`
}

type Policy struct {
	gorm.Model
	Name        string
	Arn         string
	Description string
	Manifest    string // the actual json policy
	AccountId   string
	Users       []*User `gorm:"many2many:policy_user"`
	Roles       []*Role    `gorm:"many2many:policy_role"`
}

type Role struct {
	gorm.Model
	Arn         string
	Name        string
	Description string
	Polices     []*Policy `gorm:"many2many:policy_role"`
	AccountId   string
}

type Group struct {
	gorm.Model
	Arn         string
	Name        string
	Description string
	Users       []User `gorm:"foreignKey:GroupID"`
	AccountId   string
}

type User struct {
	gorm.Model
	Name        string 
	Password    string
	Arn         string
	Description string
	Polices     []Policy   `gorm:"many2many:policy_user"`
	AccessKeys  []AccessKey `gorm:"foreignKey:UserID"`
	AccountId   string
	GroupID     *uint 
}

type AccessKey struct {
	AccessKeyId string `gorm:"primaryKey"`
	Arn         string
	CreatedAt   time.Time
	UserID   uint
}

