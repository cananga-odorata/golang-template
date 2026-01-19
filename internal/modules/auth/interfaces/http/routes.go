package http

import (
	"github.com/cananga-odorata/golang-template/internal/modules/auth/application"
	"github.com/go-chi/chi/v5"
)

// RegisterRoutes registers all auth module routes
func RegisterRoutes(r chi.Router, service application.AuthService) {
	handler := NewAuthHandler(service)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
		r.Post("/refresh", handler.RefreshToken)
	})
}
