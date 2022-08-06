package di

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/infrastructure"
	"github.com/RistekCSUI/sistech-finpro/interfaces"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"go.uber.org/dig"
	"log"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := application.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := interfaces.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := infrastructure.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}
