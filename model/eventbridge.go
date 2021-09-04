package model

import "gorm.io/gorm"

type Rule struct {
	gorm.Model
	Name           string   `gorm:"type:varchar(100);not null;index:rule_idx,unique"`
	AccountId      string   `gorm:"index:rule_idx,unique"`
	Arn            string   `gorm:"type:varchar(200);not null"`
	Description    string   `gorm:"type:varchar(255);not null"`
	ServiceEventID *uint    // /lambda.LambdaService/CreateLambda
	EventPattern   *string  //	cron(* 30 * * * *)
	Targets        []Target `gorm:"foreignkey:RuleID"`
	Active         bool     `gorm:"type:boolean;not null;default:true"`
}

type ServiceEvent struct {
	gorm.Model
	Method string
	Name   string
	Path   string `gorm:"type:varchar(100);not null;index:svc_idx,unique"`
	Rules  []Rule `gorm:"foreignkey:ServiceEventID"`
}

type Target struct {
	gorm.Model
	RuleID uint   `gorm:"not null"`
	Type   string `gorm:"type:varchar(100);not null"`
	Arn    string `gorm:"type:varchar(255);not null"`
}
