package auth

import (
	"github.com/cananga-odorata/golang-template/internal/modules/auth/application"
	"github.com/cananga-odorata/golang-template/internal/modules/auth/infrastructure"
	httpinterface "github.com/cananga-odorata/golang-template/internal/modules/auth/interfaces/http"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// Module represents the auth module with all its dependencies
type Module struct {
	Service application.AuthService
}

// NewModule initializes the auth module with all dependencies
func NewModule(db *sqlx.DB, jwtSecret string) *Module {
	repo := infrastructure.NewPostgresAuthRepository(db)
	service := application.NewAuthService(repo, jwtSecret)

	return &Module{
		Service: service,
	}
}

// RegisterRoutes registers the module's HTTP routes
func (m *Module) RegisterRoutes(r chi.Router) {
	httpinterface.RegisterRoutes(r, m.Service)
}
