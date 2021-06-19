package main

import (
	"context"
	"fmt"
	"os"

	database "github.com/MauricioAntonioMartinez/aws/db"
	"github.com/MauricioAntonioMartinez/aws/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var (
	Secret string
	db *gorm.DB
)

var ctx = context.Background()

func init() {
	Secret = os.Getenv("SECRET")
	fmt.Println(Secret)
}


func logger() zerolog.Logger { 
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	w := zerolog.ConsoleWriter{Out: os.Stderr}
	l := zerolog.New(w).With().Timestamp().Caller().Logger()
	return l 
}

func main() {
	fastergoding.Run()
	l := logger()	
	app := fiber.New()
	db,err :=database.New("MyDb")

	if err !=nil {
		err = fmt.Errorf("Failed to connect database: %w",err)
		l.Fatal().Err(err).Msg("Filed to connect the database")
	}

	l.Info().Str("name",db.Dialector.Name()).Str("database",db.Debug().Name()).Msg("Succeeded to connect to the database")


	redis := database.NewRedis() 

	// db.AutoMigrate(&model.Queue{},
	// 	&model.User{},&model.Message{})
	db.AutoMigrate(db)

	handler(app, redis)
	app.Listen(":3000")

}





func validateQueueOwner(c *fiber.Ctx,queue *model.Queue) error { 

	    user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		accountId := claims["accountId"].(string)
		qAccountId := c.Params("accountId")

		if accountId != qAccountId {
			return fiber.NewError(401,"Not Authorized")
		}

		name := c.Params("queueName")


		if res := db.First(&queue, "account_id = ? AND name = ?", qAccountId, name); res.Error != nil {
			return fiber.NewError(404,"Cannot find the queue")
		}

		return nil

}


