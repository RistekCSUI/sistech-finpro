package middleware

import (
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) RoleAdminCheck(c *fiber.Ctx) error {
	user := c.Locals("auth").(dto.User)
	if user.Role == dto.USER {
		m.shared.Logger.Infof("role mismatched for user id: %s", user.ID)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "role mismatched error"})
	}
	return c.Next()
}
