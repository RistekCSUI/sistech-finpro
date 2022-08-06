package infrastructure

import (
	"github.com/RistekCSUI/sistech-finpro/infrastructure/access"
	"github.com/RistekCSUI/sistech-finpro/infrastructure/authentication"
	"github.com/RistekCSUI/sistech-finpro/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Access         access.Controller
		Authentication authentication.Controller
		Middleware     middleware.Middleware
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(access.NewAccessController); err != nil {
		return errors.Wrap(err, "failed to provide access controller")
	}

	if err := container.Provide(authentication.NewAuthenticationController); err != nil {
		return errors.Wrap(err, "failed to provide auth controller")
	}

	if err := container.Provide(middleware.NewMiddleware); err != nil {
		return errors.Wrap(err, "failed to provide middleware")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Access.AccessRoutes(app)
	controller.Authentication.AuthenticationRoutes(app)
}
