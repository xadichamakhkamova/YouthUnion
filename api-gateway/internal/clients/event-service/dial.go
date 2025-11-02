package eventservice

import (
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/eventpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithEventService(cfg config.Config) (*pb.EventServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.EventService.Host, cfg.EventService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewEventServiceClient(conn)
	return &userServiceClient, nil
}