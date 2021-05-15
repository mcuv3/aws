package main

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Time = int

const (
	Second = iota
	Minute
	Hour
	BaseUrl = "http://localhost:3000"
)

type User struct {
	gorm.Model
	AccountId string
	Queues    []Queue
}

type Queue struct {
	gorm.Model
	Name     string
	Url      string `gorm:"unique"`
	UserID   uint
	Conf     ConfigurationQueue `gorm:"embedded"`
	messages []Message
}

type Message struct {
	gorm.Model
	message string
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
	Conf      ConfigurationQueue
	AccountId string
	Name      string
}

func (u *User) JSON() string {
	data, _ := json.Marshal(u)
	return string(data)
}
