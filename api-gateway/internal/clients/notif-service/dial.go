package notifservice

import (
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithNotifService(cfg config.Config) (*pb.NotificationServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.NotificationService.Host, cfg.NotificationService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewNotificationServiceClient(conn)
	return &userServiceClient, nil
}