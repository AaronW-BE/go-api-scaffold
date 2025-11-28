package handler

import (
	"go-api-scaffold/internal/logger"
	"go-api-scaffold/internal/service"
	"go-api-scaffold/internal/util"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Base        *BaseHandler
	UserService service.UserService
}

func NewUserHandler(base *BaseHandler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Base:        base,
		UserService: userService,
	}
}

// GetUser @Summary Get user info
// @Param id path int true "User ID"
// @Success 200 {string} string "ok"
// @Router /user/{id} [get]
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	name := r.URL.Query().Get("name")

	user, err := h.UserService.GetUser(r.Context(), id, name)
	if err != nil {
		logger.Log.Error("failed to get user", "error", err, "id", id)
		util.Json(w, nil, 500, "internal server error")
		return
	}

	util.Json(w, user, 0, "ok")
}

func (h *UserHandler) RegisterRoutes(r chi.Router) {
	r.Get("/user/{id}", h.GetUser)
}
