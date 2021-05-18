package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func connectDatabase() {
	dsn := "host=db user=mcuve password=maotrix1 dbname=sqs port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = database

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected with the database")

}
