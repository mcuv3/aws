package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() (db *gorm.DB) {
	dsn := "host=db user=mcuve password=maotrix1 dbname=sqs port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected with the database")

	return db

}
