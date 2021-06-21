package main

import (
	"log"

	database "github.com/MauricioAntonioMartinez/aws/db"
)


func main(){
	db,err := database.New()
	_ = db

	if err!=nil{
		log.Fatal("Unable to connect to database Lambda.")
	}
}