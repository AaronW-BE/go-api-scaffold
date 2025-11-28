package db

import (
	"database/sql"
	"fmt"
	"go-api-scaffold/internal/config"
	"go-api-scaffold/internal/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Log.Info("Database connection established successfully",
		"host", cfg.DB.Host,
		"port", cfg.DB.Port,
		"database", cfg.DB.Database,
	)
	return db, nil
}
