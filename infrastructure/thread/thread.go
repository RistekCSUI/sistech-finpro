package thread

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

func (c *Controller) ThreadRoutes(app *fiber.App) {
	thread := app.Group("/thread")

	thread.Use(c.Middleware.AccessCheck)

	thread.Post("/", c.Middleware.AuthCheck, c.createThread)

	thread.Put("/:id", c.Middleware.AuthCheck, c.Middleware.RoleAdminCheck, c.editThread)

	thread.Delete("/:id", c.Middleware.AuthCheck, c.Middleware.RoleAdminCheck, c.deleteThread)

	thread.Get("/:id", c.getAllThreadPost)
}

func (c *Controller) createThread(ctx *fiber.Ctx) error {
	var (
		requestBody dto.CreateThreadDto
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

	res, err := c.Interfaces.ThreadViewService.CreateThread(dto.CreateThreadRequest{
		CategoryID: requestBody.CategoryID,
		Name:       requestBody.Name,
		FirstPost:  requestBody.FirstPost,
		Token:      ctx.Locals("user").(string),
		Owner:      ctx.Locals("auth").(dto.User).ID.Hex(),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *Controller) editThread(ctx *fiber.Ctx) error {
	var (
		requestBody dto.EditThreadDto
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

	res, err := c.Interfaces.ThreadViewService.EditThread(dto.EditThreadRequest{
		EditThreadDto: requestBody,
		Token:         ctx.Locals("user").(string),
		ThreadID:      ctx.Params("id"),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *Controller) deleteThread(ctx *fiber.Ctx) error {
	res, err := c.Interfaces.ThreadViewService.DeleteThread(dto.DeleteThreadRequest{
		ID:    ctx.Params("id"),
		Token: ctx.Locals("user").(string),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *Controller) getAllThreadPost(ctx *fiber.Ctx) error {
	res, err := c.Interfaces.PostViewService.GetAllPostByThread(dto.GetAllPostRequest{
		Token:    ctx.Locals("user").(string),
		ThreadID: ctx.Params("id"),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func NewThreadController(middleware middleware.Middleware, interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		middleware,
		interfaces,
		shared,
	}
}
