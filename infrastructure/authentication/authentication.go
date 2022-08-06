package authentication

import (
	"github.com/RistekCSUI/sistech-finpro/infrastructure/middleware"
	"github.com/RistekCSUI/sistech-finpro/interfaces"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Middleware middleware.Middleware
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) AuthenticationRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Use(c.Middleware.AccessCheck)
	auth.Post("/register", c.register)
	auth.Post("/login", c.login)
}

func (c *Controller) register(ctx *fiber.Ctx) error {
	var (
		requestBody dto.RegisterDto
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

	token := ctx.Locals("user").(string)
	req, err := c.Interfaces.AuthViewService.BuildRegisterRequest(&requestBody, token)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := c.Interfaces.AuthViewService.RegisterUser(*req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	c.Shared.Logger.WithFields(logrus.Fields{
		"access": token,
	}).Infof("registering new user: %s", requestBody.Username)

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *Controller) login(ctx *fiber.Ctx) error {
	var (
		requestBody dto.LoginDto
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

	token := ctx.Locals("user").(string)

	res, err := c.Interfaces.AuthViewService.Login(dto.LoginRequest{
		AccessToken: token,
		LoginDto:    requestBody,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	c.Shared.Logger.WithFields(logrus.Fields{
		"access": token,
	}).Infof("login user: %s", requestBody.Username)

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func NewAuthenticationController(middleware middleware.Middleware, shared shared.Holder, interfaces interfaces.Holder) Controller {
	return Controller{
		middleware,
		interfaces,
		shared,
	}
}
