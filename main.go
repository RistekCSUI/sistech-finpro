package main

import (
	"github.com/RistekCSUI/sistech-finpro/di"
	"github.com/RistekCSUI/sistech-finpro/infrastructure"
	"github.com/RistekCSUI/sistech-finpro/shared/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, env *config.EnvConfig, infra infrastructure.Holder) error {
		infrastructure.Routes(http, infra)
		err := http.Listen(":" + env.ServerPort)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
