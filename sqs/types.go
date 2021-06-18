package main

import (
	"encoding/json"
	"fmt"

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

type UserMessageId string;

type User struct {
	gorm.Model
	AccountId string `gorm:"primaryKey"`
	Queues    []Queue
}

type Queue struct {
	gorm.Model
	Name             string
	Url              string `gorm:"unique"`
	UserID           uint
	AccountId        string             `gorm:"primaryKey"`
	Configuration    ConfigurationQueue `gorm:"embedded"`
	messages         []Message
}


type Message struct {
	gorm.Model
	Message string
	UserMessageId UserMessageId
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

type CreateQueueBody struct {
	Configuration ConfigurationQueue
	Name          string
}


type SendMessageBody struct { 
	Id UserMessageId
	Message string
}

type DeleteMessageBody struct { 
	Ids []UserMessageId
}

func (u *User) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}


type RetriveMessageBody struct { 
	LongPooling int
	BatchLimit int
}



func (m *Message) Key() string { 
	return fmt.Sprintf("%d.%d",m.QueueID,m.ID) 
}


func (m *Queue) Pattern()string{ 
	return fmt.Sprintf("%d.*",m.ID)
}