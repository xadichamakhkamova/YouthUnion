package main 

import (
	api "api-gateway/internal/https"
	config "api-gateway/internal/config"
	_ "api-gateway/docs"

	userService "api-gateway/internal/clients/user-service"

	service "api-gateway/internal/service"
	"api-gateway/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log := logger.NewLogger()

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}
	log.Info("Configuration loaded successfully")

	conn, err := userService.DialWithUserService(*cfg)
	if err != nil {
		log.Fatal("Failed to connect to User Service:", err)
	}
	log.Info("Connected to User Service")

	clientService := service.NewServiceRepositoryClient(conn)
	log.Info("Service clients initialized")

	srv := api.NewGin(clientService, cfg.ApiGateway.Port)
	addr := fmt.Sprintf(":%d", cfg.ApiGateway.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Starting API Gateway on: ", addr)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			log.Fatal(err)
		}
	}()
	log.Info("Starting API Gateway on address:", addr)

	signalReceived := <-sigChan
	log.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Info("Graceful shutdown complete.")
}
