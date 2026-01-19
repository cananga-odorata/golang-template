package http

import (
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/modules/user/application"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all user module routes
func RegisterRoutes(r chi.Router, service application.UserService, authMiddleware func(http.Handler) http.Handler) {
	handler := NewUserHandler(service)

	r.Route("/users", func(r chi.Router) {
		// All user routes require authentication
		r.Use(authMiddleware)

		r.Post("/", handler.Create)
		r.Get("/", handler.List)
		r.Get("/{id}", handler.GetByID)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}
