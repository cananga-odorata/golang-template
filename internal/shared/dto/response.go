package dto

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/shared/domain"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorBody  `json:"error,omitempty"`
}

// ErrorBody represents the error details in a response
type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JSON sends a successful JSON response
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    data,
	})
}

// Error sends an error JSON response
func Error(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Success: false,
		Error: &ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}

// ErrorFromAppError sends an error response from AppError
func ErrorFromAppError(w http.ResponseWriter, err error) {
	var appErr *domain.AppError
	if errors.As(err, &appErr) {
		Error(w, appErr.HTTPStatus(), string(appErr.Code), appErr.Message)
		return
	}
	Error(w, http.StatusInternalServerError, string(domain.ErrCodeInternal), "Internal server error")
}

// NoContent sends a 204 No Content response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Created sends a 201 Created response
func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}

// OK sends a 200 OK response
func OK(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, data)
}
