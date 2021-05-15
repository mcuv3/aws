package main

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["accountId"] = "23425235"
	claims["userId"] = "123123"

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
