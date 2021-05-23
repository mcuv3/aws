package main

import (
	"encoding/json"

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

type AddDestinationBody struct {
	Url         string
	LongPolling int
}

func (u *User) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}


type RetriveMessageBody struct { 
	LongPulling int
	BatchLimit int
}