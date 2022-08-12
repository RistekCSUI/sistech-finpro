package interfaces

import (
	"github.com/RistekCSUI/sistech-finpro/interfaces/access"
	"github.com/RistekCSUI/sistech-finpro/interfaces/authentication"
	"github.com/RistekCSUI/sistech-finpro/interfaces/category"
	"github.com/RistekCSUI/sistech-finpro/interfaces/thread"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	AccessViewService   access.ViewService
	AuthViewService     authentication.ViewService
	CategoryViewService category.ViewService
	ThreadViewService   thread.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(access.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide access service")
	}

	if err := container.Provide(authentication.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide auth service")
	}

	if err := container.Provide(category.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide category service")
	}

	if err := container.Provide(thread.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide thread service")
	}

	return nil
}
