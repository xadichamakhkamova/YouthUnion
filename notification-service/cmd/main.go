package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "notification-service/internal/app"
	config "notification-service/internal/config"
	"notification-service/internal/http/handler"
	"notification-service/internal/repository"
	pq "notification-service/internal/repository/postgres"
	service "notification-service/internal/service"
	"notification-service/logger"
)

func main() {
	// Initialize logger
	log := logger.NewLogger()
	log.Info("Starting Notification Service...")

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
	queries := repository.NewNotifStore(db, log)
	repo := repository.NewNotifRepository(queries, log)
	log.Info("Notification repository initialized")

	// Initialize service
	srv := service.NewNotifService(repo)
	log.Info("Notification service initialized")

	// Initialize gRPC application
	application := app.New(*srv, cfg.Service.Port, log)
	log.Infof("Starting gRPC server on port %d", cfg.Service.Port)

	// Run server in a goroutine
	go func() {
		application.MustRun()
	}()

	log.Info("Server is running and ready to accept requests")

	//Websocket
	h := handler.Handler{
		Log: log,
	}

	http.HandleFunc("/ws", h.HandleWebSocket)
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	addr := fmt.Sprintf(":%s", cfg.WebSocket.Port)
	websocket := &http.Server{
		Addr:      addr,
		Handler:   http.DefaultServeMux,
		TLSConfig: tlsConfig,
	}
	go func() {
		log.Info("Websocket listening on ", addr)
		if err := websocket.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			log.Fatal(err)
		}
	}()

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
