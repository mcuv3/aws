package sqs

import (
	"context"
	"fmt"
	"log"

	"github.com/MauricioAntonioMartinez/aws/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func login(c *fiber.Ctx) error {

	user := model.User{}

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


func checkCredentials(token *jwt.Token) string {

	claims := token.Claims.(jwt.MapClaims)

	accountId := claims["accountId"].(string)
	
	return accountId
}


func withJwt(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Println(ctx)
    log.Println("--> unary interceptor: ", info.FullMethod)
    return handler(ctx, req)
}
