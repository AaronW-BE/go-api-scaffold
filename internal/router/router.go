package router

import (
	"go-api-scaffold/internal/handler"
	"go-api-scaffold/internal/registry"

	"github.com/go-chi/chi/v5"
)

func BuildRouter(base *handler.BaseHandler) chi.Router {
	r := chi.NewRouter()

	for _, ctor := range registry.List() {
		h := invokeHandler(ctor, base)
		h.RegisterRoutes(r)
	}

	return r
}

func invokeHandler(ctor interface{}, base *handler.BaseHandler) handler.Router {
	return ctor.(func(*handler.BaseHandler) handler.Router)(base)
}
