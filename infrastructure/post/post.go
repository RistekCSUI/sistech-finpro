package post

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

func (c *Controller) PostRoutes(app *fiber.App) {
	post := app.Group("/post")

	post.Use(c.Middleware.AccessCheck)

	post.Post("/", c.Middleware.AuthCheck, c.createPost)
}

func (c *Controller) createPost(ctx *fiber.Ctx) error {
	var (
		requestBody dto.CreatePostDto
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

	res, err := c.Interfaces.PostViewService.CreatePost(dto.CreatePostRequest{
		ThreadID: requestBody.ThreadID,
		Content:  requestBody.Content,
		ReplyID:  requestBody.ReplyID,
		Token:    ctx.Locals("user").(string),
		Owner:    ctx.Locals("auth").(dto.User).ID.Hex(),
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func NewPostController(middleware middleware.Middleware, interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		middleware,
		interfaces,
		shared,
	}
}
