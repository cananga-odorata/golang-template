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
	"github.com/cananga-odorata/golang-template/internal/server"
)

func main() {
	// 1. Load Config
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// 2. Init Server
	s := server.New(cfg)

	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: s.Router,
	}

	// 3. Start Server in Goroutine
	// รันแยก thread ไป เพื่อไม่ให้บล็อกการรอสัญญาณปิด
	go func() {
		slog.Info("Server is starting", "addr", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Could not listen", "error", err)
			os.Exit(1)
		}
	}()

	// 4. Graceful Shutdown
	// สร้างช่องทางดักฟังสัญญาณจาก OS (เช่น Ctrl+C หรือ Docker Stop)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// รอจนกว่าจะมีสัญญาณส่งมา (โค้ดจะค้างบรรทัดนี้จนกว่าจะกดปิด)
	<-stop

	slog.Info("Shutting down server...")

	// ให้เวลาเคลียร์งานที่ค้างอยู่ 10 วินาที
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Server exited")
}
