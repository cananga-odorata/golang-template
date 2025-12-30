package auth

import (
	"encoding/json"
	"net/http"

	"github.com/cananga-odorata/golang-template/pkg/response"
)

type Service interface {
	Register(req RegisterRequest) (*AuthResponse, error)
}

type Handler struct {
	service Service //Inject Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	// JSON from body (Structural Validation)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Basic Validation
	if req.Email == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "Email and password are required")
		return
	}

	res, err := h.service.Register(req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	// response.JSON(w, http.StatusCreated, data)
	response.JSON(w, http.StatusCreated, res)
}
