package service

import (
	"context"
	"database/sql"
	"fmt"
)

type UserService interface {
	GetUser(ctx context.Context, id string, name string) (map[string]string, error)
}

type userServiceImpl struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return &userServiceImpl{
		db: db,
	}
}

func (s *userServiceImpl) GetUser(ctx context.Context, id string, name string) (map[string]string, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is not available")
	}

	return map[string]string{
		"id":   id,
		"name": name,
		"desc": "hello " + name,
	}, nil
}
