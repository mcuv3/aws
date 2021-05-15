package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	db := connectDatabase()

	db.AutoMigrate(&Queue{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Message{})

	app.Use(recover.New())

	app.Post("/login", login)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

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
		SigningKey: []byte("secret"),
	}))

	app.Post("/queue/:accountId/:queueName", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		// accountId := claims["accountId"].(string)
		userId := claims["userId"].(string)

		queue := Queue{}

		if res := db.First(&queue, "user_id = ?", userId); res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Cannot find the queue."})
		}

		message := Message{}

		if err := json.Unmarshal(c.Body(), &message); err != nil {
			return c.JSON(ErrorResponse{Status: 400, Message: "Invalid JSON format."})
		}

		db.Create(&message)

		return c.JSON(message)
	})

	app.Post("/queue", func(c *fiber.Ctx) error {

		queue := Queue{}

		err := json.Unmarshal(c.Body(), &queue)

		if err != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Invalid payload."})
		}

		user := User{}

		res := db.First(&user, queue.UserID)

		if res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Unable to find a user"})
		}

		queue.Url = BaseUrl + "/" + user.AccountId + "/" + queue.Name

		fmt.Println(user.JSON())

		res = db.Create(&queue)

		if res.Error != nil {
			return c.Status(400).JSON(ErrorResponse{Status: 400, Message: "Couldn't create record."})
		}

		return c.SendString(string(c.Body()))
	})

	app.Listen(":3000")

}
