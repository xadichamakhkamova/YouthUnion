package app

import (
	s "team-service/internal/service"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
	log        *logrus.Logger
}

func New(srv s.TeamService, port int, log *logrus.Logger) *App {
	grpcServer := grpc.NewServer()
	pb.RegisterTeamServiceServer(grpcServer, &srv)

	return &App{
		gRPCServer: grpcServer,
		port:       port,
		log:        log,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		a.log.Fatalf("gRPC server failed to start: %v", err)
	}
}

func (a *App) Run() error {
	addr := fmt.Sprintf(":%d", a.port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		a.log.Errorf("Failed to start TCP listener: %v", err)
		return err
	}

	a.log.Infof("gRPC server is starting on port %d", a.port)

	// Run server
	go func() {
		if err := a.gRPCServer.Serve(listener); err != nil {
			a.log.Errorf("gRPC server stopped unexpectedly: %v", err)
		}
	}()

	// Wait for OS signals (Ctrl+C, kill)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit

	a.log.Warnf("Received shutdown signal: %v", sig)
	a.Stop()
	return nil
}

func (a *App) Stop() {
	a.log.Infof("Stopping gRPC server on port %d...", a.port)
	a.gRPCServer.GracefulStop()
	a.log.Info("gRPC server stopped cleanly")
}
