package application

import (
	"github.com/RistekCSUI/sistech-finpro/application/access"
	"github.com/RistekCSUI/sistech-finpro/application/authentication"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	AccessService access.Service
	AuthService   authentication.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(access.NewService); err != nil {
		return errors.Wrap(err, "failed to provide access app")
	}

	if err := container.Provide(authentication.NewService); err != nil {
		return errors.Wrap(err, "failed to provide auth app")
	}

	return nil
}
