package main

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func login(c *fiber.Ctx) error {

	user := User{}

	db.First(&user)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["accountId"] = user.AccountId
	claims["userId"] = user.ID

	t, err := token.SignedString([]byte(Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
