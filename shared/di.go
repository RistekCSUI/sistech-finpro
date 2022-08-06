package shared

import (
	"github.com/RistekCSUI/sistech-finpro/shared/config"
	"github.com/RistekCSUI/sistech-finpro/shared/depedencies"
	"github.com/RistekCSUI/sistech-finpro/shared/utils"
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Env    *config.EnvConfig
	App    *fiber.App
	Mongo  *mongo.Database
	Logger *logrus.Logger
	Redis  *redis.Client
	JWT    utils.JWT
}

func Register(container *dig.Container) error {
	if err := container.Provide(config.NewEnvConfig); err != nil {
		return errors.Wrap(err, "failed to provide env")
	}

	if err := container.Provide(depedencies.NewLogger); err != nil {
		return errors.Wrap(err, "failed to provide logger")
	}

	if err := container.Provide(depedencies.NewHttp); err != nil {
		return errors.Wrap(err, "failed to provide fiber")
	}

	if err := container.Provide(depedencies.NewMongo); err != nil {
		return errors.Wrap(err, "failed to provide mongo")
	}

	if err := container.Provide(utils.NewJWT); err != nil {
		return errors.Wrap(err, "failed to provide jwt")
	}

	if err := container.Provide(depedencies.NewRedis); err != nil {
		return errors.Wrap(err, "failed to provide redis")
	}

	return nil
}
