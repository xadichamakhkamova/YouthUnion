package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/scoringpb"
)

func NewScoringRepository(repo *ScoringRepo, log *logrus.Logger) *ScoringRepo {
	return &ScoringRepo{
		queries: repo.queries,
		log:     log,
	}
}

type IScoringRepository interface {
	GiveScore(ctx context.Context, req *pb.GiveScoreRequest) (*pb.Score, error)
	GetScoresByEvent(ctx context.Context, req *pb.GetScoresByEventRequest) (*pb.ScoreList, error)
	GetScoresByUser(ctx context.Context, req *pb.GetScoresByUserRequest) (*pb.ScoreList, error)
	GetScoresByTeam(ctx context.Context, req *pb.GetScoresByTeamRequest) (*pb.ScoreList, error)
	GetGlobalRanking(ctx context.Context, req *pb.GetGlobalRankingRequest) (*pb.RankingList, error)
}
