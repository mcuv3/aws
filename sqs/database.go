package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
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


func connectCache()  *redis.Client{ 
	rdb := redis.NewClient(&redis.Options{
        Addr:     "cache:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

	return rdb
}