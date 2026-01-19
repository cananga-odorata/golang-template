package user

import (
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/modules/user/application"
	"github.com/cananga-odorata/golang-template/internal/modules/user/infrastructure"
	httpinterface "github.com/cananga-odorata/golang-template/internal/modules/user/interfaces/http"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// Module represents the user module with all its dependencies
type Module struct {
	Service application.UserService
}

// NewModule initializes the user module with all dependencies
func NewModule(db *sqlx.DB) *Module {
	repo := infrastructure.NewPostgresUserRepository(db)
	service := application.NewUserService(repo)

	return &Module{
		Service: service,
	}
}

// RegisterRoutes registers the module's HTTP routes
func (m *Module) RegisterRoutes(r chi.Router, authMiddleware func(http.Handler) http.Handler) {
	httpinterface.RegisterRoutes(r, m.Service, authMiddleware)
}
