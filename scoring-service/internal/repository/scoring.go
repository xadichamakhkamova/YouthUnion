package repository

import (
	"context"
	"database/sql"
	"scoring-service/internal/storage"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/scoringpb"
)

type ScoringRepo struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewScoringStore(db *sql.DB, log *logrus.Logger) *ScoringRepo {
	return &ScoringRepo{
		queries: storage.New(db),
		log:     log,
	}
}

func (q *ScoringRepo) GiveScore(ctx context.Context, req *pb.GiveScoreRequest) (*pb.Score, error) {

	event_id, err := uuid.Parse(req.EventId)
	if err != nil {
		return nil, err
	}
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	scored_by_id, err := uuid.Parse(req.ScoredById)
	if err != nil {
		return nil, err
	}
	resp, err := q.queries.GiveScore(ctx, storage.GiveScoreParams{
		EventID:      event_id,
		UserID:       user_id,
		Points:       req.Points,
		Comment:      sql.NullString{String: req.Comment},
		ScoredByID:   scored_by_id,
		ScoredByType: storage.ScoredByType(req.ScoredByType.String()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.Score{
		Id:           resp.ID.String(),
		EventId:      resp.EventID.String(),
		TeamId:       resp.TeamID.UUID.String(),
		UserId:       resp.UserID.String(),
		Points:       resp.Points,
		Comment:      resp.Comment.String,
		ScoredById:   resp.ScoredByID.String(),
		ScoredByType: pb.ScoredByType(pb.ScoredByType_value[string(resp.ScoredByType)]),
		CreatedAt:    resp.CreatedAt.Time.String(),
	}, nil
}

func (q *ScoringRepo) GetScoresByEvent(ctx context.Context, req *pb.GetScoresByEventRequest) (*pb.ScoreList, error) {

	event_id, err := uuid.Parse(req.EventId)
	if err != nil {
		return nil, err
	}
	params := storage.GetScoreByEventParams{
		EventID: event_id,
	}
	resp, err := q.queries.GetScoreByEvent(ctx, params)
	if err != nil {
		return nil, err
	}

	var scores []*pb.Score
	var total_count int64
	for _, r := range resp {
		scores = append(scores, &pb.Score{
			Id:           r.ID.String(),
			EventId:      r.EventID.String(),
			TeamId:       r.TeamID.UUID.String(),
			UserId:       r.UserID.String(),
			Points:       r.Points,
			Comment:      r.Comment.String,
			ScoredById:   r.ScoredByID.String(),
			ScoredByType: pb.ScoredByType(pb.ScoredByType_value[string(r.ScoredByType)]),
			CreatedAt:    r.CreatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.ScoreList{
		Scores:     scores,
		TotalCount: int32(total_count),
	}, nil
}

func (q *ScoringRepo) GetScoresByUser(ctx context.Context, req *pb.GetScoresByUserRequest) (*pb.ScoreList, error) {

	event_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	params := storage.GetScoreByEventParams{
		EventID: event_id,
	}
	resp, err := q.queries.GetScoreByEvent(ctx, params)
	if err != nil {
		return nil, err
	}

	var scores []*pb.Score
	var total_count int64
	for _, r := range resp {
		scores = append(scores, &pb.Score{
			Id:           r.ID.String(),
			EventId:      r.EventID.String(),
			TeamId:       r.TeamID.UUID.String(),
			UserId:       r.UserID.String(),
			Points:       r.Points,
			Comment:      r.Comment.String,
			ScoredById:   r.ScoredByID.String(),
			ScoredByType: pb.ScoredByType(pb.ScoredByType_value[string(r.ScoredByType)]),
			CreatedAt:    r.CreatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.ScoreList{
		Scores:     scores,
		TotalCount: int32(total_count),
	}, nil
}

func (q *ScoringRepo) GetScoresByTeam(ctx context.Context, req *pb.GetScoresByTeamRequest) (*pb.ScoreList, error) {

	event_id, err := uuid.Parse(req.TeamId)
	if err != nil {
		return nil, err
	}
	params := storage.GetScoreByEventParams{
		EventID: event_id,
	}
	resp, err := q.queries.GetScoreByEvent(ctx, params)
	if err != nil {
		return nil, err
	}

	var scores []*pb.Score
	var total_count int64
	for _, r := range resp {
		scores = append(scores, &pb.Score{
			Id:           r.ID.String(),
			EventId:      r.EventID.String(),
			TeamId:       r.TeamID.UUID.String(),
			UserId:       r.UserID.String(),
			Points:       r.Points,
			Comment:      r.Comment.String,
			ScoredById:   r.ScoredByID.String(),
			ScoredByType: pb.ScoredByType(pb.ScoredByType_value[string(r.ScoredByType)]),
			CreatedAt:    r.CreatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.ScoreList{
		Scores:     scores,
		TotalCount: int32(total_count),
	}, nil
}


func (q *ScoringRepo) GetGlobalRanking(ctx context.Context, req *pb.GetGlobalRankingRequest) (*pb.RankingList, error) {
	offset := (req.Page - 1) * req.Limit

	params := storage.GetGlobalRankingParams{
		Limit:  req.Limit,
		Offset: offset,
	}
	resp, err := q.queries.GetGlobalRanking(ctx, params)
	if err != nil {
		return nil, err
	}

	var rankings []*pb.Ranking
	for _, r := range resp {
		rankings = append(rankings, &pb.Ranking{
			UserId:      r.UserID.String(),
			TotalPoints: int32(r.TotalPoints),
		})
	}
	return &pb.RankingList{
		Rankings: rankings,
	}, nil
}
