package main

import (
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
	"gorm.io/driver/sqlite"
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

func init() {
	Secret = os.Getenv("SECRET")
	fmt.Println(Secret)
}

func startDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}

func main() {
	fastergoding.Run()
	app := fiber.New()

	connectDatabase()

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

	fmt.Println(Secret)

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
		body := RetriveMessageBody{}

		err := json.Unmarshal(c.Body(),&body)

		if err!=nil { 
			return c.Status(400).JSON(ErrorResponse{Status: 400,Message:"Invalid payload"})
		}

		msgs := make(chan []Message)	

		
		go func(){
			
			messages := make([]Message,0)	
			var n int
			
			for  {
				m := Message{}
				 db.First(&m,"")
				 messages = append(messages, m)
				time.Sleep(1)
				n++

				if body.BatchLimit == len(messages) || body.LongPulling == n { 
					break
				} 

			}
				msgs <- messages

		}()


		return c.JSON(<- msgs)


	})

	// Send a message to the queue
	app.Post("/:accountId/:queueName", func(c *fiber.Ctx) error {

		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		accountId := claims["accountId"].(string)
		qAccountId := c.Params("accountId")

		if accountId != qAccountId {
			return c.Status(401).JSON(ErrorResponse{Status: 400, Message: "Not Authorized"})
		}

		name := c.Params("queueName")

		queue := Queue{}

		if res := db.First(&queue, "account_id = ? AND name = ?", qAccountId, name); res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Cannot find the queue."})
		}



		msg := Message{
			Message: "Testing",
			QueueID:queue.ID,
		}

		res := db.Create(&msg)

		if res.Error !=nil{
			return c.Status(400).JSON(ErrorResponse{Message: "Could not save the message",Status: 400})
		}



		fmt.Println(queue)

		return c.JSON(msg)
	})

	app.Listen(":3000")

}
