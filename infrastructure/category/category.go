package category

import (
	"github.com/RistekCSUI/sistech-finpro/infrastructure/middleware"
	"github.com/RistekCSUI/sistech-finpro/interfaces"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Middleware middleware.Middleware
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) CategoryRoutes(app *fiber.App) {
	category := app.Group("/category")

	category.Use(c.Middleware.AccessCheck)
	category.Use(c.Middleware.AuthCheck)

	category.Post("/", c.Middleware.RoleAdminCheck, c.createCategory)
}

func (c *Controller) createCategory(ctx *fiber.Ctx) error {
	var (
		requestBody dto.CreateCategoryDto
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

	res, err := c.Interfaces.CategoryViewService.CreateCategory(dto.CreateCategoryRequest{
		CreateCategoryDto: requestBody,
		Token:             ctx.Locals("user").(string),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func NewCategoryController(middleware middleware.Middleware, interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		middleware,
		interfaces,
		shared,
	}
}
