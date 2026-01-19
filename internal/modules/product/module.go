package product

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// Module represents the product module
type Module struct {
	db *sqlx.DB
}

// NewModule initializes the product module
func NewModule(db *sqlx.DB) *Module {
	return &Module{db: db}
}

// RegisterRoutes registers the module's HTTP routes
func (m *Module) RegisterRoutes(r chi.Router, authMiddleware func(http.Handler) http.Handler) {
	r.Route("/products", func(r chi.Router) {
		r.Use(authMiddleware)
		// TODO: Implement product handlers
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message":"product list - not implemented"}`))
		})
	})
}
