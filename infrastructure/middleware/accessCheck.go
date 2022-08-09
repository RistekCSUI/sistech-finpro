package middleware

import (
	"context"
	"fmt"
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

type Middleware struct {
	application application.Holder
	shared      shared.Holder
}

const AccessTokenCacheName = "sistech-access-token-"

func (m *Middleware) AccessCheck(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if len(authHeader) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header not provided"})
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong authorization format"})
	}

	key := fmt.Sprintf("%s%s", AccessTokenCacheName, authToken[1])

	_, err := m.shared.Redis.Get(context.Background(), key).Result()
	if err == nil {
		c.Locals("user", authToken[1])
		m.shared.Logger.Infof("access token exist in cache for token: %s", authToken[1])
		return c.Next()
	}

	_, err = m.application.AccessService.FindByToken(authToken[1])
	if err != nil {
		m.shared.Logger.Errorf("no access key found for token: %s", authToken[1])
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go func() {
		m.shared.Logger.Infof("setting cache for user token: %s", authToken[1])
		m.shared.Redis.SetEx(context.Background(), key, authToken[1], time.Minute*3)
	}()

	c.Locals("user", authToken[1])

	return c.Next()
}

func NewMiddleware(application application.Holder, shared shared.Holder) Middleware {
	return Middleware{
		application: application,
		shared:      shared,
	}
}
