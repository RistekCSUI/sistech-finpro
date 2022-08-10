package middleware

import (
	"context"
	"fmt"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"time"
)

const UserTokenCacheName = "sistech-user-"

func (m *Middleware) AuthCheck(c *fiber.Ctx) error {
	userToken := c.Get("X-USER-TOKEN")
	if len(userToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "X-USER-TOKEN header not provided"})
	}

	userId := m.shared.JWT.ExtractTokenData(userToken)
	key := fmt.Sprintf("%s%s", UserTokenCacheName, userId)

	dataCache, err := m.shared.Redis.Get(context.Background(), key).Result()
	if err == nil {
		var data dto.User
		_ = json.Unmarshal([]byte(dataCache), &data)
		c.Locals("auth", data)
		m.shared.Logger.Infof("user token exist in cache for id: %s", userId)
		return c.Next()
	}

	res, err := m.application.AuthService.FindUserByID(userId)

	if err != nil {
		m.shared.Logger.Infof("no user found for given id: %s", userId)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no user found for given token"})
	}

	if res.AccessToken != c.Locals("user") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid access for other access token"})
	}

	go func() {
		m.shared.Logger.Infof("setting cache for user id: %s", userId)
		encode, _ := json.Marshal(res)
		m.shared.Redis.SetEx(context.Background(), key, string(encode), time.Minute*1)
	}()

	c.Locals("auth", res)

	return c.Next()
}
