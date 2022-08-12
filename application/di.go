package application

import (
	"github.com/RistekCSUI/sistech-finpro/application/access"
	"github.com/RistekCSUI/sistech-finpro/application/authentication"
	"github.com/RistekCSUI/sistech-finpro/application/category"
	"github.com/RistekCSUI/sistech-finpro/application/thread"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	AccessService   access.Service
	AuthService     authentication.Service
	CategoryService category.Service
	ThreadService   thread.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(access.NewService); err != nil {
		return errors.Wrap(err, "failed to provide access app")
	}

	if err := container.Provide(authentication.NewService); err != nil {
		return errors.Wrap(err, "failed to provide auth app")
	}

	if err := container.Provide(category.NewService); err != nil {
		return errors.Wrap(err, "failed to provide category app")
	}

	if err := container.Provide(thread.NewService); err != nil {
		return errors.Wrap(err, "failed to provide thread app")
	}

	return nil
}
