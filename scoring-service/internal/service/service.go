package service

import (
	"context"
	"scoring-service/internal/repository"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/scoringpb"
)

type ScoringService struct {
	pb.UnimplementedScoringServiceServer
	repo repository.IScoringRepository
}

func NewScoringService(repo repository.IScoringRepository) *ScoringService {
	return &ScoringService{
		repo: repo,
	}
}

func (s *ScoringService) GiveScore(ctx context.Context, req *pb.GiveScoreRequest) (*pb.Score, error) {
	return s.repo.GiveScore(ctx, req)
}

func (s *ScoringService) GetScoresByEvent(ctx context.Context, req *pb.GetScoresByEventRequest) (*pb.ScoreList, error) {
	return s.repo.GetScoresByEvent(ctx, req)
}

func (s *ScoringService) GetScoresByUser(ctx context.Context, req *pb.GetScoresByUserRequest) (*pb.ScoreList, error) {
	return s.repo.GetScoresByUser(ctx, req)
}

func (s *ScoringService) GetScoresByTeam(ctx context.Context, req *pb.GetScoresByTeamRequest) (*pb.ScoreList, error) {
	return s.repo.GetScoresByTeam(ctx, req)
}

func (s *ScoringService) GetGlobalRanking(ctx context.Context, req *pb.GetGlobalRankingRequest) (*pb.RankingList, error) {
	return s.repo.GetGlobalRanking(ctx, req)
}
