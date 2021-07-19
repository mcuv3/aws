package sqs

import (
	"github.com/MauricioAntonioMartinez/aws/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)
func validateQueueOwner(c *fiber.Ctx,queue *model.Queue) error { 

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	accountId := claims["accountId"].(string)
	qAccountId := c.Params("accountId")

	if accountId != qAccountId {
		return fiber.NewError(401,"Not Authorized")
	}

	// name := c.Params("queueName")


	// if res := db.First(&queue, "account_id = ? AND name = ?", qAccountId, name); res.Error != nil {
	// 	return fiber.NewError(404,"Cannot find the queue")
	// }

	return nil

}




