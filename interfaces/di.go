package interfaces

import (
	"github.com/RistekCSUI/sistech-finpro/interfaces/access"
	"github.com/RistekCSUI/sistech-finpro/interfaces/authentication"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	AccessViewService access.ViewService
	AuthViewService   authentication.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(access.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide access service")
	}

	if err := container.Provide(authentication.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide auth service")
	}

	return nil
}
