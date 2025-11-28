package errors

import "net/http"

type ErrorCode int

const (
	ErrSuccess ErrorCode = 0
	ErrBadRequest ErrorCode = 400
	ErrUnauthorized ErrorCode = 401
	ErrForbidden ErrorCode = 403
	ErrNotFound ErrorCode = 404
	ErrInternalServer ErrorCode = 500
	ErrDatabaseError ErrorCode = 5001
)

type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func (e *AppError) HTTPStatusCode() int {
	switch e.Code {
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrForbidden:
		return http.StatusForbidden
	case ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
