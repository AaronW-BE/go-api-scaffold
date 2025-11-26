package main

import (
	"context"
	"errors"
	"fmt"
	"go-api-scaffold/internal/app"
	"go-api-scaffold/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "go-api-scaffold/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Go Api scaffold API
// @version         1.0
// @description     chi + dig + auto-registry + swagger
// @host            :8080
// @BasePath        /
func main() {
	c := app.Bootstrap()

	_ = c.Invoke(func(r chi.Router, cfg *config.Config) {
		r.Get("/swagger/*", httpSwagger.Handler())

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
			Handler: r,
		}

		go func() {
			fmt.Println("Server running on", srv.Addr)
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("ListenAndServe(): %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, os.Kill)
		<-quit
		fmt.Println("\nShutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}

		fmt.Println("Server exited gracefully")

	})

}
