package scoringservice

import (
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/scoringpb"
	config "api-gateway/internal/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithScoringService(cfg config.Config) (*pb.ScoringServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.ScoringService.Host, cfg.ScoringService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userServiceClient := pb.NewScoringServiceClient(conn)
	return &userServiceClient, nil
}