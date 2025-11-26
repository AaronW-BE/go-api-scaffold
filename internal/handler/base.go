package handler

import (
	"database/sql"
	"go-api-scaffold/internal/config"

	"github.com/go-chi/chi/v5"
)

type BaseHandler struct {
	Config *config.Config
	DB     *sql.DB
}

func (h *BaseHandler) GetDB() *sql.DB {
	return h.DB
}

func NewBaseHandler(cfg *config.Config, db *sql.DB) *BaseHandler {
	return &BaseHandler{
		Config: cfg,
		DB:     db,
	}
}

type Router interface {
	RegisterRoutes(r chi.Router)
}
