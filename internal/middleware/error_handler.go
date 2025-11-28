package middleware

import (
	"go-api-scaffold/internal/errors"
	"go-api-scaffold/internal/logger"
	"go-api-scaffold/internal/util"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error("panic recovered",
					"error", err,
					"path", r.URL.Path,
					"method", r.Method,
				)
				util.Json(w, nil, int(errors.ErrInternalServer), "internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}
