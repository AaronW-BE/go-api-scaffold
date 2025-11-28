package app

import (
	"go-api-scaffold/internal/config"
	"go-api-scaffold/internal/db"
	"go-api-scaffold/internal/handler"
	"go-api-scaffold/internal/router"
	"go-api-scaffold/internal/service"

	"go.uber.org/dig"
)

func Bootstrap() *dig.Container {
	config.LoadConfig()

	c := dig.New()
	_ = c.Provide(func() *config.Config { return &config.Conf })

	if err := c.Provide(db.NewDB); err != nil {
		panic(err)
	}

	_ = c.Provide(service.NewUserService)
	_ = c.Provide(handler.NewBaseHandler)
	_ = c.Provide(handler.NewUserHandler)
	_ = c.Provide(router.BuildRouter)

	return c
}
