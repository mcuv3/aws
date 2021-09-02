package model

import "gorm.io/gorm"

type Rule struct {
	gorm.Model
	Name           string   `gorm:"type:varchar(100);not null" validation:"required;gte=5"`
	Arn            string   `gorm:"type:varchar(200);not null;uniqueIndex"`
	Description    string   `gorm:"type:varchar(255);not null"`
	ServiceEventID *uint    // /lambda.LambdaService/CreateLambda
	EventPattern   string   //	cron(* 30 * * * *)
	Targets        []Target `gorm:"foreignkey:RuleID"`
}

type ServiceEvent struct {
	gorm.Model
	Path  string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Rules []Rule `gorm:"foreignkey:ServiceEventID"`
}

type Target struct {
	gorm.Model
	RuleID uint   `gorm:"not null"`
	Type   string `gorm:"type:varchar(100);not null"`
	Arn    string `gorm:"type:varchar(255);not null"`
}
