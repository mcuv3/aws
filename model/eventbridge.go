package model

import "gorm.io/gorm"

type Rule struct {
	gorm.Model
	Name          string   `gorm:"type:varchar(100);not null"`
	Arn           string   `gorm:"type:varchar(200);not null"`
	Description   string   `gorm:"type:varchar(255);not null"`
	ScheduleEvent string   // /lambda.LambdaService/CreateLambda
	EventPattern  string   // cron(* 30 * * * *)
	Targets       []Target `gorm:"foreignkey:RuleID"`
}

type Target struct {
	gorm.Model
	RuleID       uint   `gorm:"not null"`
	TargetTypeID uint   `gorm:"not null"`
	Arn          string `gorm:"type:varchar(255);not null"`
}

type TargetType struct {
	gorm.Model
	Name    string
	Targets []Target `gorm:"foreignkey:TargetTypeID"`
}
