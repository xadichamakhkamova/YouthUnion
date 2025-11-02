package teamservice

import (
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithTeamService(cfg config.Config) (*pb.TeamServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.TeamService.Host, cfg.TeamService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewTeamServiceClient(conn)
	return &userServiceClient, nil
}