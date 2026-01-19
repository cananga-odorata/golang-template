package domain

import "net/http"

// ErrorCode represents application error codes
type ErrorCode string

const (
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrCodeConflict     ErrorCode = "CONFLICT"
	ErrCodeInternal     ErrorCode = "INTERNAL_SERVER_ERROR"
)

// AppError represents a structured application error
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// HTTPStatus returns the corresponding HTTP status code
func (e *AppError) HTTPStatus() int {
	return ErrorCodeToHTTPStatus(e.Code)
}

// Error constructors
func NewNotFoundError(message string) *AppError {
	return &AppError{Code: ErrCodeNotFound, Message: message}
}

func NewValidationError(message string) *AppError {
	return &AppError{Code: ErrCodeValidation, Message: message}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{Code: ErrCodeUnauthorized, Message: message}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{Code: ErrCodeForbidden, Message: message}
}

func NewConflictError(message string) *AppError {
	return &AppError{Code: ErrCodeConflict, Message: message}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{Code: ErrCodeInternal, Message: message, Err: err}
}

// ErrorCodeToHTTPStatus maps error codes to HTTP status codes
func ErrorCodeToHTTPStatus(code ErrorCode) int {
	switch code {
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeValidation:
		return http.StatusBadRequest
	case ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case ErrCodeForbidden:
		return http.StatusForbidden
	case ErrCodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
