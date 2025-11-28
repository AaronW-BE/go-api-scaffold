package router

import (
	"go-api-scaffold/internal/handler"
	"go-api-scaffold/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func BuildRouter(userHandler *handler.UserHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.ErrorHandler)

	userHandler.RegisterRoutes(r)

	return r
}
