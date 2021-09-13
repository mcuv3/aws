package model

import "gorm.io/gorm"

type CloudTrailEvent struct {
	gorm.Model
	TrailID      string
	Method       string
	Source       string
	ResourceName string
	Region       string
}

type Trail struct {
	gorm.Model
	Events      []CloudTrailEvent `gorm:"foreignkey:TrailID"`
	AccountId   string
	Name        string `gorm:"not null"`
	Arn         string `gorm:"not null"`
	TrailConfig `gorm:"embedded;embedded_prefix:config_"`
}

type TrailConfig struct {
	Write bool `gorm:"type:boolean;not null;default:true"`
	Read  bool `gorm:"type:boolean;not null;default:true"`
}
