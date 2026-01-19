package server

import (
	"log/slog"
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/config"
	"github.com/cananga-odorata/golang-template/internal/modules/affiliate"
	"github.com/cananga-odorata/golang-template/internal/modules/auth"
	"github.com/cananga-odorata/golang-template/internal/modules/product"
	"github.com/cananga-odorata/golang-template/internal/modules/user"
	"github.com/cananga-odorata/golang-template/internal/shared/middleware"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
)

// Server holds the HTTP server dependencies
type Server struct {
	Router *chi.Mux
	Config *config.Config
	DB     *sqlx.DB
}

// New creates a new server with all modules wired
func New(cfg *config.Config, db *sqlx.DB) *Server {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.RequestID)

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Request-ID"},
		ExposedHeaders:   []string{"Link", "X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check
	r.Get("/health", healthHandler)

	// Initialize modules
	authModule := auth.NewModule(db, cfg.JWTSecret)
	userModule := user.NewModule(db)
	productModule := product.NewModule(db)
	affiliateModule := affiliate.NewModule(db)

	// Auth middleware for protected routes
	authMiddleware := middleware.JWTAuth(cfg.JWTSecret)

	// API v1 routes
	r.Route("/api/v1", func(api chi.Router) {
		// Auth routes (public)
		authModule.RegisterRoutes(api)

		// Protected routes
		userModule.RegisterRoutes(api, authMiddleware)
		productModule.RegisterRoutes(api, authMiddleware)
		affiliateModule.RegisterRoutes(api, authMiddleware)
	})

	slog.Info("Server initialized",
		"modules", []string{"auth", "user", "product", "affiliate"},
		"environment", cfg.Environment,
	)

	return &Server{
		Router: r,
		Config: cfg,
		DB:     db,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
