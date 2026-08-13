package errors

import (
	"fmt"
	"net/http"
)

type Code string

const (
	CodeValidation   Code = "VALIDATION_ERROR"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeForbidden    Code = "FORBIDDEN"
	CodeNotFound     Code = "NOT_FOUND"
	CodeConflict     Code = "CONFLICT"
	CodeRateLimited  Code = "RATE_LIMITED"
	CodeBadRequest   Code = "BAD_REQUEST"
	CodeInternal     Code = "INTERNAL_ERROR"
	CodeDatabase     Code = "DATABASE_ERROR"
	CodeBlockchain   Code = "BLOCKCHAIN_ERROR"
	CodeStorage      Code = "STORAGE_ERROR"
	CodeUnavailable  Code = "SERVICE_UNAVAILABLE"
)

type AppError struct {
	Code       Code
	Message    string
	HTTPStatus int
	Cause      error
}

func (e *AppError) Error() string {
	if e.Cause == nil {
		return string(e.Code) + ": " + e.Message
	}

	return fmt.Sprintf(
		"%s: %s: %v",
		e.Code,
		e.Message,
		e.Cause,
	)
}

func (e *AppError) Unwrap() error {
	return e.Cause
}

func (e *AppError) SafeMessage() string {
	return e.Message
}

func New(
	code Code,
	message string,
	status int,
) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
	}
}

func Wrap(
	code Code,
	message string,
	status int,
	cause error,
) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
		Cause:      cause,
	}
}

func Validation(message string) *AppError {
	return New(
		CodeValidation,
		message,
		http.StatusBadRequest,
	)
}

func Unauthorized() *AppError {
	return New(
		CodeUnauthorized,
		"Authentication is required.",
		http.StatusUnauthorized,
	)
}

func Forbidden() *AppError {
	return New(
		CodeForbidden,
		"You are not authorized to perform this operation.",
		http.StatusForbidden,
	)
}

func NotFound(resource string) *AppError {
	return New(
		CodeNotFound,
		resource+" was not found.",
		http.StatusNotFound,
	)
}

func Conflict(message string) *AppError {
	return New(
		CodeConflict,
		message,
		http.StatusConflict,
	)
}

func Internal(cause error) *AppError {
	return Wrap(
		CodeInternal,
		"An internal error occurred.",
		http.StatusInternalServerError,
		cause,
	)
}

func Database(cause error) *AppError {
	return Wrap(
		CodeDatabase,
		"An internal database error occurred.",
		http.StatusInternalServerError,
		cause,
	)
}

func Blockchain(cause error) *AppError {
	return Wrap(
		CodeBlockchain,
		"An internal blockchain error occurred.",
		http.StatusInternalServerError,
		cause,
	)
}

func Storage(cause error) *AppError {
	return Wrap(
		CodeStorage,
		"An internal storage error occurred.",
		http.StatusInternalServerError,
		cause,
	)
}

func Unavailable() *AppError {
	return New(
		CodeUnavailable,
		"The service is temporarily unavailable.",
		http.StatusServiceUnavailable,
	)
}
