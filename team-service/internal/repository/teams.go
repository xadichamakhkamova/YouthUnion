package repository

import (
	"context"
	"database/sql"
	"team-service/internal/storage"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
)

type TeamRepo struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewTeamStore(db *sql.DB, log *logrus.Logger) *TeamRepo {
	return &TeamRepo{
		queries: storage.New(db),
		log:     log,
	}
}

func (q *TeamRepo) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.Team, error) {

	event_id, err := uuid.Parse(req.EventId)
	if err != nil {
		return nil, err
	}
	leader_id, err := uuid.Parse(req.LeaderId)
	if err != nil {
		return nil, err
	}
	resp, err := q.queries.CreateTeam(ctx, storage.CreateTeamParams{
		Name:     req.Name,
		LeaderID: leader_id,
		EventID:  event_id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Team{
		Id:        resp.ID.String(),
		Name:      resp.Name,
		LeaderId:  resp.LeaderID.String(),
		EventId:   resp.EventID.String(),
		IsReady:   resp.IsReady.Bool,
		CreatedAt: resp.CreatedAt.Time.String(),
		UpdatedAt: resp.UpdatedAt.Time.String(),
	}, nil
}

func (q *TeamRepo) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.Team, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	resp, err := q.queries.UpdateTeam(ctx, storage.UpdateTeamParams{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Team{
		Id:        resp.ID.String(),
		Name:      resp.Name,
		LeaderId:  resp.LeaderID.String(),
		EventId:   resp.EventID.String(),
		IsReady:   resp.IsReady.Bool,
		CreatedAt: resp.CreatedAt.Time.String(),
		UpdatedAt: resp.UpdatedAt.Time.String(),
	}, nil
}

func (q *TeamRepo) GetTeamsByEvent(ctx context.Context, req *pb.GetTeamsByEventRequest) (*pb.TeamList, error) {
	offset := (req.Page - 1) * req.Limit

	event_id, err := uuid.Parse(req.EventId)
	if err != nil {
		return nil, err
	}
	params := storage.GetTeamsByEventParams{
		EventID: event_id,
		Limit:   req.Limit,
		Offset:  offset,
	}
	resp, err := q.queries.GetTeamsByEvent(ctx, params)
	if err != nil {
		return nil, err
	}

	var teams []*pb.Team
	var total_count int64
	for _, r := range resp {
		teams = append(teams, &pb.Team{
			Id:        r.ID.String(),
			Name:      r.Name,
			LeaderId:  r.LeaderID.String(),
			EventId:   r.EventID.String(),
			IsReady:   r.IsReady.Bool,
			CreatedAt: r.CreatedAt.Time.String(),
			UpdatedAt: r.UpdatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.TeamList{
		Teams:      teams,
		TotalCount: int32(total_count),
	}, nil
}

func (q *TeamRepo) GetTeamMembers(ctx context.Context, req *pb.GetTeamRequest) (*pb.MemberList, error) {

	team_id, err := uuid.Parse(req.TeamId)
	if err != nil {
		return nil, err
	}
	resp, err := q.queries.GetTeamMembers(ctx, team_id)
	if err != nil {
		return nil, err
	}

	var members []*pb.TeamMember
	var total_count int64
	for _, r := range resp {
		members = append(members, &pb.TeamMember{
			Id:       r.ID.String(),
			TeamId:   r.TeamID.String(),
			UserId:   r.UserID.String(),
			Role:     pb.MemberRole(pb.MemberRole_value[r.Role]),
			JoinedAt: r.JoinedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.MemberList{
		Members:    members,
		TotalCount: int32(total_count),
	}, nil
}

func (q *TeamRepo) RemoveTeamMember(ctx context.Context, req *pb.RemoveTeamMemberRequest) (*pb.StatusResponse, error) {

	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	team_id, err := uuid.Parse(req.TeamId)
	if err != nil {
		return nil, err
	}
	err = q.queries.RemoveTeamMember(ctx, storage.RemoveTeamMemberParams{
		TeamID: team_id,
		UserID: user_id,
	})
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{
		Status: 204,
	}, nil
}
