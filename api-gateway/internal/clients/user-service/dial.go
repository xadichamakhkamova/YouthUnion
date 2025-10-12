package userservice

import (
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithUserService(cfg config.Config) (*pb.UserServiceClient, error) {

	target := fmt.Sprintf("%s:%s", cfg.UserService.Host, cfg.UserService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewUserServiceClient(conn)
	return &userServiceClient, nil
}