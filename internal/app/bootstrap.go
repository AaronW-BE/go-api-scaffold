package app

import (
	"go-api-scaffold/internal/config"
	"go-api-scaffold/internal/db"
	"go-api-scaffold/internal/handler"
	"go-api-scaffold/internal/router"

	"go.uber.org/dig"
)

func Bootstrap() *dig.Container {
	config.LoadConfig()

	c := dig.New()
	_ = c.Provide(func() *config.Config { return &config.Conf })

	err := c.Provide(db.NewDB)
	if err != nil {
		panic(err)
	}

	_ = c.Provide(handler.NewBaseHandler)
	_ = c.Provide(router.BuildRouter)

	return c
}
