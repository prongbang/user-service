package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prongbang/uam-service/internal/pkg/token"
)

func Payload(c *fiber.Ctx) token.Claims {
	a := new(token.AccessToken)
	err := c.BodyParser(a)
	if err != nil {
		return token.Claims{}
	}
	payload, err := token.Payload(a.Token)
	if err != nil {
		return token.Claims{}
	}
	return *payload
}

func PayloadByToken(accessToken string) token.Claims {
	payload, err := token.Payload(accessToken)
	if err != nil {
		return token.Claims{}
	}
	return *payload
}
