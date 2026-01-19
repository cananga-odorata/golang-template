package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cananga-odorata/golang-template/internal/config"
	"github.com/cananga-odorata/golang-template/internal/infra/database"
	"github.com/cananga-odorata/golang-template/internal/server"
	"github.com/jmoiron/sqlx"
)

func main() {
	// 1. Load Config
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	slog.Info("Configuration loaded",
		"environment", cfg.Environment,
		"port", cfg.Port,
	)

	// 2. Initialize Database (optional - skip if no DB configured)
	var db *sqlx.DB
	if cfg.Database != nil && cfg.Database.Host != "" {
		db, err = database.NewPostgresDB(cfg.Database)
		if err != nil {
			slog.Warn("Database connection failed, continuing without database", "error", err)
			db = nil
		} else {
			defer db.Close()
			slog.Info("Database connected",
				"host", cfg.Database.Host,
				"dbname", cfg.Database.DBName,
			)
		}
	}

	// 3. Initialize Server
	s := server.New(cfg, db)

	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      s.Router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 4. Start Server in Goroutine
	go func() {
		slog.Info("Server is starting", "addr", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Could not listen", "error", err)
			os.Exit(1)
		}
	}()

	// 5. Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server exited")
}
