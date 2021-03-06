package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Time = string

const (
	Second = iota
	Minute
	Hour
	BaseUrl = "http://localhost:3000"
)

var times = [3]string{"second", "minute", "hour"}

type MessageStatus int 

const (
	Available MessageStatus = iota 
	Processing
	Consumed
)


type UserMessageId string;

type Queue struct {
	gorm.Model
	Name             string
	Arn 			 string   		    `gorm:"not null"`
	AccountId        string             `gorm:"primaryKey"`
	Configuration    ConfigurationQueue `gorm:"embedded"`
	messages         []Message
}


type Message struct {
	gorm.Model
	Message string
	UserMessageId UserMessageId
	QueueName string
	Status MessageStatus `gorm:"default:0"`
	DeleteOn time.Time
	QueueID uint
}

type ConfigurationQueue struct { 
	VisibilityTimeout    int
	VisibilityTime       Time
	MessageRetention     int
	MessageRetentionTime Time
	DeliveryDelay        int
	DeliveryDelayTime    Time
}


func (m *Message) Key() string { 
	return fmt.Sprintf("%d.%d",m.QueueID,m.ID) 
}


func (m *Queue) Pattern()string{ 
	return fmt.Sprintf("%d.*",m.ID)
}