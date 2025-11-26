package handler

import (
	"go-api-scaffold/internal/util"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Base *BaseHandler
}

func NewUserHandler(base *BaseHandler) Router {
	return &UserHandler{Base: base}
}

// GetUser @Summary Get user info
// @Param id path int true "用户ID"
// @Success 200 {string} string "ok"
// @Router /user/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if h.Base.GetDB() != nil {
		log.Println("database is ok")
	}

	util.Json(w, map[string]string{
		"id":   chi.URLParam(r, "id"),
		"name": r.URL.Query().Get("name"),
		"desc": "hello" + r.URL.Query().Get("name"),
	}, 0, "ok")
}

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Get("/user/{id}", h.GetUser)
}
