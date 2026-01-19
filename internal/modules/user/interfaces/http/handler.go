package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cananga-odorata/golang-template/internal/modules/user/application"
	"github.com/cananga-odorata/golang-template/internal/modules/user/domain"
	"github.com/cananga-odorata/golang-template/internal/shared/dto"
	"github.com/go-chi/chi/v5"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	service application.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Create handles POST /users
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req application.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.Error(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request payload")
		return
	}

	// Basic validation
	if req.Email == "" || req.Password == "" {
		dto.Error(w, http.StatusBadRequest, "VALIDATION_ERROR", "Email and password are required")
		return
	}

	user, err := h.service.Create(r.Context(), req)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.Created(w, user)
}

// GetByID handles GET /users/{id}
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.OK(w, user)
}

// Update handles PUT /users/{id}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req application.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.Error(w, http.StatusBadRequest, "INVALID_JSON", "Invalid request payload")
		return
	}

	user, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.OK(w, user)
}

// Delete handles DELETE /users/{id}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.service.Delete(r.Context(), id); err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.NoContent(w)
}

// List handles GET /users
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	filter := domain.UserFilter{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
		Search: r.URL.Query().Get("search"),
	}

	// Optional role filter
	if role := r.URL.Query().Get("role"); role != "" {
		r := domain.Role(role)
		filter.Role = &r
	}

	// Optional status filter
	if status := r.URL.Query().Get("status"); status != "" {
		s := domain.Status(status)
		filter.Status = &s
	}

	users, total, err := h.service.List(r.Context(), filter)
	if err != nil {
		dto.ErrorFromAppError(w, err)
		return
	}

	dto.OK(w, map[string]interface{}{
		"data":        users,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}
