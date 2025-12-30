package server

import (
	"net/http"

	"github.com/cananga-odorata/golang-template/internal/auth"
	"github.com/cananga-odorata/golang-template/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router *chi.Mux
	Config *config.Config
}

func New(cfg *config.Config) *Server {
	r := chi.NewRouter()

	//Global Middleware
	setupMiddleware(r)

	//Handler Check Route
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"`))
	})

	authService := auth.NewService()
	authHandler := auth.NewHandler(authService)

	// API Routes เตรียมสำหรับ Versioning
	r.Route("/api/v1", func(api chi.Router) {
		// api.Post("/auth/register", authHandler.Register)
		// Create Group Auth and Register Route
		api.Route("/auth", func(authRouter chi.Router) {
			authRouter.Post("/register", authHandler.Register)
		})

	})

	return &Server{
		Router: r,
		Config: cfg,
	}
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)    // Logging
	r.Use(middleware.Recoverer) // Panic Recovery
	r.Use(middleware.RealIP)    // ดึง IP ที่แท้จริงของลูกค้า

	// CORS Setup
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}
