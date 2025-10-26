package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
)

func NewTeamRepository(repo *TeamRepo, log *logrus.Logger) ITeamRepository {
	return &TeamRepo{
		queries: repo.queries,
		log: log,
	}
}

type ITeamRepository interface {
	//Team
	CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.Team, error)
	UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.Team, error)
	GetTeamsByEvent(ctx context.Context, req *pb.GetTeamsByEventRequest) (*pb.TeamList, error) 
	GetTeamMembers(ctx context.Context, req *pb.GetTeamRequest) (*pb.MemberList, error)
	RemoveTeamMember(ctx context.Context, req *pb.RemoveTeamMemberRequest) (*pb.StatusResponse, error)
}