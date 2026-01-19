package http

import (
	"encoding/json"
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/modules/auth/application"
	"github.com/cananga-odorata/golang-template/internal/shared/dto"
)

// AuthHandler handles HTTP requests for authentication
type AuthHandler struct {
	service application.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(service application.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register handles POST /auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req application.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.Error(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request payload")
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		dto.Error(w, http.StatusBadRequest, "VALIDATION_ERROR", "Email and password are required")
		return
	}

	res, err := h.service.Register(r.Context(), req)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.Created(w, res)
}

// Login handles POST /auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req application.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.Error(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request payload")
		return
	}

	if req.Email == "" || req.Password == "" {
		dto.Error(w, http.StatusBadRequest, "VALIDATION_ERROR", "Email and password are required")
		return
	}

	res, err := h.service.Login(r.Context(), req)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.OK(w, res)
}

// RefreshToken handles POST /auth/refresh
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req application.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.Error(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request payload")
		return
	}

	res, err := h.service.RefreshToken(r.Context(), req)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.OK(w, res)
}
