package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/qinains/fastergoding"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var (
	Secret string
)

var ctx = context.Background()

func init() {
	Secret = os.Getenv("SECRET")
	fmt.Println(Secret)
}


func main() {
	fastergoding.Run()
	app := fiber.New()

	connectDatabase()
	redis := connectCache()

	db.AutoMigrate(&Queue{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})

	app.Use(recover.New())

	app.Post("/login", login) // just for testing porpuses

	app.Post("/user", func(c *fiber.Ctx) error {
		AccountId := c.Query("accountId")

		user := User{
			AccountId: AccountId,
		}

		res := db.Create(&user)

		if res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Cannot create user"})
		}

		return c.JSON(user)

	})

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(Secret),
	}))

	app.Patch("/purge/:id", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		accountId := claims["accountId"]

		fmt.Println(accountId)
		return c.JSON(claims)
	})



	app.Post("/create", func(c *fiber.Ctx) error {

		body := CreateQueueBody{}

		err := json.Unmarshal(c.Body(), &body)

		if err != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Invalid payload."})
		}

		us := c.Locals("user").(*jwt.Token)
		claims := us.Claims.(jwt.MapClaims)

		accountId := claims["accountId"]

		user := User{}

		if res := db.First(&user, "account_id = ?", accountId); res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Cannot the user."})
		}

		fmt.Println(body.Configuration)

		if body.Configuration.DeliveryDelayTime == "" || body.Configuration.MessageRetentionTime == "" || body.Configuration.VisibilityTime == "" {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Invalid configuration please check."})
		}

		name := strings.ReplaceAll(body.Name, " ", "")

		queue := Queue{
			Url:           BaseUrl + "/" + user.AccountId + "/" + name,
			Configuration: body.Configuration,
			AccountId:     user.AccountId,
			Name:          name,
		}

		res := db.Create(&queue)

		if res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Couldn't create queue."})
		}

		return c.JSON(queue)
	})

	// Retreive manually a message from the queue
	app.Get("/:accountId/:queueName", func(c *fiber.Ctx) error {
		queue := Queue{}
		err := validateQueueOwner(c,&queue)
		
		if err !=nil { 
			return c.JSON(err)
		}
		
		body := RetriveMessageBody{}
		err = json.Unmarshal(c.Body(),&body)
		
 
		if err!=nil { 
			return c.Status(400).JSON(ErrorResponse{Status: 400,Message:"Invalid payload"})
		}


		msgs := make(chan []Message)	
		
		fmt.Println(body)
		go func(){
			
			messages := make([]Message,0)	 
			var n int
			
			for  {
				m := Message{}

				keys,err:=  redis.Keys(ctx,queue.Pattern()).Result()

				if err == nil {
					for _,k:= range keys {
						val,_ :=  redis.Get(ctx,k).Result()
						redis.Del(ctx,k)
						json.Unmarshal([]byte(val),&m)

						messages = append(messages, m)
					}
				}

				
				
				time.Sleep(time.Second)
				n++
					
				if body.BatchLimit == len(messages) || body.LongPooling <= n { 
					break
				}
				
			}

				msgs <- messages

		}()


		return c.JSON(<- msgs)


	})

	// Send a message to the queue
	app.Post("/:accountId/:queueName", func(c *fiber.Ctx) error {
		queue := Queue{}
		err := validateQueueOwner(c,&queue)

		if err !=nil { 
			return c.JSON(err)
		}

		message := SendMessageBody{}

		err = json.Unmarshal(c.Body(),&message)

		if err!=nil {
			return c.Status(400).JSON(ErrorResponse{Message: "Invalid payload."})
		}

		msg := Message{
			Message: message.Message,
			QueueID:queue.ID,
		}

		// redis.Publish(ctx, "__keyspace@0__ DEL", "hello").Err()

		
		res := db.Create(&msg)
		m,_:= json.Marshal( msg)

		err = redis.Set(ctx,msg.Key(),m,0).Err()

		if res.Error !=nil || err !=nil {
			return c.Status(400).JSON(ErrorResponse{Message: "Could not save the message",Status: 400})
		}


		return c.JSON(msg)
	})

	app.Delete("/:accountId/:queueName",func(c *fiber.Ctx)error{
		
		queue:= Queue{}

		err:= validateQueueOwner(c,&queue)

		if err!=nil { 
			return c.Status(401).JSON(ErrorResponse{Message: "Not Authorized"})
		}

		body := DeleteMessageBody{}

		err = json.Unmarshal(c.Body(),&body)

		if err !=nil { 
			return c.Status(400).JSON(ErrorResponse{Message:"Invalid payload sorry 😥"})
		}
		var messages []Message

		
		var toDelete []int
		
		for _,m := range messages { 
			v,_ :=redis.Get(ctx,fmt.Sprintf("%d.%d",queue.ID,m.ID)).Result()
			if v == "" {
				toDelete = append(toDelete, int(m.ID))
			}
		}

		db.Delete(&messages,toDelete)
		
		return c.JSON(SuccessResponse{Message: "Successfully deleted"})
	})

	pubsub := redis.Subscribe(ctx, "__keyspace@0__:*")

// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)
	if err != nil {
		panic(err)
	}
	redis.Set(ctx,"TEST","!",0).Err()

	// Go channel which receives messages.
	ch := pubsub.Channel()


	// time.AfterFunc(time.Second, func() {
	// fmt.Println("CHANNEL CLOSED")
	// // _ = pubsub.Close()
	// })


 
	go func(){
		for msg := range ch {
			fmt.Println("NEW MESSAGE")
			fmt.Println(msg.Channel, msg.Payload)
		}
	}()


	app.Listen(":3000")

}



func validateQueueOwner(c *fiber.Ctx,queue *Queue) error { 

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


