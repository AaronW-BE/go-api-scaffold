package db

import (
	"database/sql"
	"fmt"
	"go-api-scaffold/internal/config"
	"log"
)

func NewDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("error opening database connection: %v", err)
	}
	return db
}
