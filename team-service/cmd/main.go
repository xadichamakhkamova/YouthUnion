package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "team-service/internal/app"
	config "team-service/internal/config"
	"team-service/internal/repository"
	pq "team-service/internal/repository/postgres"
	service "team-service/internal/service"
	"team-service/logger"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()
	log.Info("Starting Team Service...")

	// Load configuration
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Info("Configuration loaded successfully")

	// Connect to Postgres
	db, err := pq.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres database: %v", err)
	}
	log.Info("Connected to Postgres database")

	// Initialize repositories
	queries := repository.NewTeamStore(db, log)
	repo := repository.NewTeamRepository(queries, log)
	log.Info("Team repository initialized")

	// Initialize service
	srv := service.NewTeamService(repo)
	log.Info("Team service initialized")

	// Initialize gRPC application
	application := app.New(*srv, cfg.Service.Port, log)
	log.Infof("Starting gRPC server on port %d", cfg.Service.Port)

	// Run server in a goroutine
	go func() {
		application.MustRun()
	}()

	log.Info("Server is running and ready to accept requests")

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop
	log.Warnf("Received termination signal: %v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	application.Stop()
	log.Info("Server stopped gracefully")

	<-ctx.Done()
	log.Info("Shutdown complete")
}
