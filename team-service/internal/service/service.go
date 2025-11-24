package service

import (
	"context"
	"team-service/internal/repository"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
)

type TeamService struct {
	pb.UnimplementedTeamServiceServer
	repo repository.ITeamRepository
}

func NewTeamService(repo repository.ITeamRepository) *TeamService {
	return &TeamService{
		repo: repo,
	}
}

func (s *TeamService) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.Team, error) {
	return s.repo.CreateTeam(ctx, req)
}

func (s *TeamService) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.Team, error) {
	return s.repo.UpdateTeam(ctx, req)
}

func (s *TeamService) GetTeamsByEvent(ctx context.Context, req *pb.GetTeamsByEventRequest) (*pb.TeamList, error) {
	return s.repo.GetTeamsByEvent(ctx, req)
}

func (s *TeamService) GetTeamMembers(ctx context.Context, req *pb.GetTeamRequest) (*pb.MemberList, error) {
	return s.repo.GetTeamMembers(ctx, req)
}

func (s *TeamService) RemoveTeamMember(ctx context.Context, req *pb.RemoveTeamMemberRequest) (*pb.StatusResponse, error) {
	return s.repo.RemoveTeamMember(ctx, req)
}

func (s *TeamService) InviteMember(ctx context.Context, req *pb.InviteMemberRequest) (*pb.InvitationsResponse, error) {
	return s.repo.InviteMember(ctx, req)
}

func (s *TeamService) RespondInvite(ctx context.Context, req *pb.RespondInviteRequest) (*pb.InvitationsResponse, error) {
	return s.repo.RespondInvite(ctx, req)
}