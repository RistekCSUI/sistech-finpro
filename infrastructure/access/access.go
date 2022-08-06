package access

import (
	"github.com/RistekCSUI/sistech-finpro/interfaces"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) AccessRoutes(app *fiber.App) {
	acces := app.Group("/access")
	acces.Post("/register", c.register)
}

func (c *Controller) register(ctx *fiber.Ctx) error {
	var (
		requestBody dto.CreateAccessDto
	)

	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	err = validate.Struct(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	c.Shared.Logger.Infof("registering new access key for user: %s", requestBody.Name)

	res, err := c.Interfaces.AccessViewService.Register(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func NewAccessController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}
