package sqs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/MauricioAntonioMartinez/aws/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
)

func handler(app *fiber.App, redis *redis.Client) {

	app.Use(recover.New())

	app.Post("/login", login) // just for testing porpuses

	app.Post("/user", func(c *fiber.Ctx) error {
		AccountId := c.Query("accountId")

		user := model.User{
			AccountId: AccountId,
		}

		res := db.Create(&user)

		if res.Error != nil {
			return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Cannot create user"})
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

		body := model.CreateQueueBody{}

		err := json.Unmarshal(c.Body(), &body)

		if err != nil {
			return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Invalid payload."})
		}

		us := c.Locals("user").(*jwt.Token)
		claims := us.Claims.(jwt.MapClaims)

		accountId := claims["accountId"]

		user := model.User{}

		if res := db.First(&user, "account_id = ?", accountId); res.Error != nil {
			return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Cannot the user."})
		}

		fmt.Println(body.Configuration)

		if body.Configuration.DeliveryDelayTime == "" || body.Configuration.MessageRetentionTime == "" || body.Configuration.VisibilityTime == "" {
			return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Invalid configuration please check."})
		}

		name := strings.ReplaceAll(body.Name, " ", "")

		queue := model.Queue{
			Url:           model.BaseUrl + "/" + user.AccountId + "/" + name,
			Configuration: body.Configuration,
			AccountId:     user.AccountId,
			Name:          name,
		}

		res := db.Create(&queue)

		if res.Error != nil {
			return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Couldn't create queue."})
		}

		return c.JSON(queue)
	})

	// Retreive manually a message from the queue
	app.Get("/:accountId/:queueName", func(c *fiber.Ctx) error {
		queue := model.Queue{}
		err := validateQueueOwner(c, &queue)

		if err != nil {
			return c.JSON(err)
		}

		body := model.RetriveMessageBody{}
		err = json.Unmarshal(c.Body(), &body)

		if err != nil {
			return c.Status(400).JSON(model.ErrorResponse{Status: 400, Message: "Invalid payload"})
		}

		msgs := make(chan []model.Message)

		fmt.Println(body)
		go func() {

			messages := make([]model.Message, 0)
			var n int

			for {
				m := model.Message{}

				keys, err := redis.Keys(ctx, queue.Pattern()).Result()

				if err == nil {
					for _, k := range keys {
						val, _ := redis.Get(ctx, k).Result()
						redis.Del(ctx, k)
						json.Unmarshal([]byte(val), &m)

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

		return c.JSON(<-msgs)

	})

	// Send a message to the queue
	app.Post("/:accountId/:queueName", func(c *fiber.Ctx) error {
		queue := model.Queue{}
		err := validateQueueOwner(c, &queue)

		if err != nil {
			return c.JSON(err)
		}

		message := model.SendMessageBody{}

		err = json.Unmarshal(c.Body(), &message)

		if err != nil {
			return c.Status(400).JSON(model.ErrorResponse{Message: "Invalid payload."})
		}

		msg := model.Message{
			Message: message.Message,
			QueueID: queue.ID,
		}

		// redis.Publish(ctx, "__keyspace@0__ DEL", "hello").Err()

		res := db.Create(&msg)

		m, _ := json.Marshal(msg)

		err = redis.Set(ctx, msg.Key(), m, 0).Err()

		if res.Error != nil || err != nil {
			return c.Status(400).JSON(model.ErrorResponse{Message: "Could not save the message", Status: 400})
		}

		return c.JSON(msg)
	})

	app.Delete("/:accountId/:queueName", func(c *fiber.Ctx) error {

		queue := model.Queue{}

		err := validateQueueOwner(c, &queue)

		if err != nil {
			return c.Status(401).JSON(model.ErrorResponse{Message: "Not Authorized"})
		}

		body := model.DeleteMessageBody{}

		err = json.Unmarshal(c.Body(), &body)

		if err != nil {
			return c.Status(400).JSON(model.ErrorResponse{Message: "Invalid payload sorry 😥"})
		}
		var messages []model.Message

		var toDelete []int

		for _, m := range messages {
			v, _ := redis.Get(ctx, fmt.Sprintf("%d.%d", queue.ID, m.ID)).Result()
			if v == "" {
				toDelete = append(toDelete, int(m.ID))
			}
		}

		db.Delete(&messages, toDelete)

		return c.JSON(model.SuccessResponse{Message: "Successfully deleted"})
	})

	pubsub := redis.Subscribe(ctx, "__keyspace@0__:*")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive(ctx)

	if err != nil {
		panic(err)
	}
	redis.Set(ctx, "TEST", "!", 0).Err()

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// time.AfterFunc(time.Second, func() {
	// fmt.Println("CHANNEL CLOSED")
	// // _ = pubsub.Close()
	// })

	go func() {
		for msg := range ch {
			fmt.Println("NEW MESSAGE")
			fmt.Println(msg.Channel, msg.Payload)
		}
	}()

}
