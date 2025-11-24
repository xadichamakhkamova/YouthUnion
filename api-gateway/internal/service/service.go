package service

import (
	"context"

	pEvent "github.com/xadichamakhkamova/YouthUnionContracts/genproto/eventpb"
	pNotif "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
	pScoring "github.com/xadichamakhkamova/YouthUnionContracts/genproto/scoringpb"
	pTeam "github.com/xadichamakhkamova/YouthUnionContracts/genproto/teampb"
	pUser "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

type ServiceRepositoryClient struct {
	userClient    pUser.UserServiceClient
	eventClient   pEvent.EventServiceClient
	teamClient    pTeam.TeamServiceClient
	scoringClient pScoring.ScoringServiceClient
	notifClient   pNotif.NotificationServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pUser.UserServiceClient,
	conn2 *pEvent.EventServiceClient,
	conn3 *pTeam.TeamServiceClient,
	conn4 *pScoring.ScoringServiceClient,
	conn5 *pNotif.NotificationServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		userClient:    *conn1,
		eventClient:   *conn2,
		teamClient:    *conn3,
		scoringClient: *conn4,
		notifClient:   *conn5,
	}
}

//! ------------------- Authorization -------------------

func (s *ServiceRepositoryClient) CreateUser(ctx context.Context, req *pUser.CreateUserRequest) (*pUser.User, error) {
	return s.userClient.CreateUser(ctx, req)
}

func (s *ServiceRepositoryClient) GetUserByIdentifier(ctx context.Context, req *pUser.GetUserByIdentifierRequest) (*pUser.GetUserByIdentifierResponse, error) {
	return s.userClient.GetUserByIdentifier(ctx, req)
}

func (s *ServiceRepositoryClient) ChangePassword(ctx context.Context, req *pUser.ChangePasswordRequest) (*pUser.ChangePasswordResponse, error) {
	return s.userClient.ChangePassword(ctx, req)
}

//! ------------------ User Functions -------------------

func (s *ServiceRepositoryClient) GetUserById(ctx context.Context, req *pUser.GetUserByIdRequest) (*pUser.User, error) {
	return s.userClient.GetUserById(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateUser(ctx context.Context, req *pUser.UpdateUserRequest) (*pUser.User, error) {
	return s.userClient.UpdateUser(ctx, req)
}

func (s *ServiceRepositoryClient) ListUsers(ctx context.Context, req *pUser.ListUsersRequest) (*pUser.UserList, error) {
	return s.userClient.ListUsers(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteUser(ctx context.Context, req *pUser.DeleteUserRequest) (*pUser.DeleteUserResponse, error) {
	return s.userClient.DeleteUser(ctx, req)
}

//! ------------------- User Roles -------------------

func (s *ServiceRepositoryClient) AssignRoleToUser(ctx context.Context, req *pUser.AssignRoleRequest) (*pUser.UserRole, error) {
	return s.userClient.AssignRoleToUser(ctx, req)
}

func (s *ServiceRepositoryClient) RemoveRoleFromUser(ctx context.Context, req *pUser.RemoveRoleRequest) (*pUser.RemoveRoleResponse, error) {
	return s.userClient.RemoveRoleFromUser(ctx, req)
}

func (s *ServiceRepositoryClient) ListUserRoles(ctx context.Context, req *pUser.ListUserRolesRequest) (*pUser.UserRoleList, error) {
	return s.userClient.ListUserRoles(ctx, req)
}

//! ------------------- Role functions -------------------

func (s *ServiceRepositoryClient) CreateRole(ctx context.Context, req *pUser.CreateRoleRequest) (*pUser.RoleType, error) {
	return s.userClient.CreateRole(ctx, req)
}

func (s *ServiceRepositoryClient) GetRoleById(ctx context.Context, req *pUser.GetRoleByIdRequest) (*pUser.RoleType, error) {
	return s.userClient.GetRoleById(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateRole(ctx context.Context, req *pUser.UpdateRoleRequest) (*pUser.RoleType, error) {
	return s.userClient.UpdateRole(ctx, req)
}

func (s *ServiceRepositoryClient) ListRoles(ctx context.Context, req *pUser.ListRolesRequest) (*pUser.RoleTypeList, error) {
	return s.userClient.ListRoles(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteRole(ctx context.Context, req *pUser.DeleteRoleRequest) (*pUser.DeleteRoleResponse, error) {
	return s.userClient.DeleteRole(ctx, req)
}

//! ------------------- Event -------------------

func (s *ServiceRepositoryClient) CreateEvent(ctx context.Context, req *pEvent.CreateEventRequest) (*pEvent.Event, error) {
	return s.eventClient.CreateEvent(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateEvent(ctx context.Context, req *pEvent.UpdateEventRequest) (*pEvent.Event, error) {
	return s.eventClient.UpdateEvent(ctx, req)
}

func (s *ServiceRepositoryClient) GetEvent(ctx context.Context, req *pEvent.GetEventRequest) (*pEvent.Event, error) {
	return s.eventClient.GetEvent(ctx, req)
}

func (s *ServiceRepositoryClient) ListEvents(ctx context.Context, req *pEvent.ListEventsRequest) (*pEvent.ListEventsResponse, error) {
	return s.eventClient.ListEvents(ctx, req)
}

func (s *ServiceRepositoryClient) DeleteEvent(ctx context.Context, req *pEvent.DeleteEventRequest) (*pEvent.DeleteEventResponse, error) {
	return s.eventClient.DeleteEvent(ctx, req)
}

func (s *ServiceRepositoryClient) RegisterEvent(ctx context.Context, req *pEvent.RegisterEventRequest) (*pEvent.EventParticipant, error) {
	return s.eventClient.RegisterEvent(ctx, req)
}

func (s *ServiceRepositoryClient) ListParticipants(ctx context.Context, req *pEvent.EventParticipantRequest) (*pEvent.EventParticipantResponse, error) {
	return s.eventClient.ListParticipants(ctx, req)
}

//! ------------------- Team -------------------

func (s *ServiceRepositoryClient) CreateTeam(ctx context.Context, req *pTeam.CreateTeamRequest) (*pTeam.Team, error) {
	return s.teamClient.CreateTeam(ctx, req)
}

func (s *ServiceRepositoryClient) UpdateTeam(ctx context.Context, req *pTeam.UpdateTeamRequest) (*pTeam.Team, error) {
	return s.teamClient.UpdateTeam(ctx, req)
}

func (s *ServiceRepositoryClient) GetTeamsByEvent(ctx context.Context, req *pTeam.GetTeamsByEventRequest) (*pTeam.TeamList, error) {
	return s.teamClient.GetTeamsByEvent(ctx, req)
}

func (s *ServiceRepositoryClient) GetTeamMembers(ctx context.Context, req *pTeam.GetTeamRequest) (*pTeam.MemberList, error) {
	return s.teamClient.GetTeamMembers(ctx, req)
}

func (s *ServiceRepositoryClient) RemoveTeamMember(ctx context.Context, req *pTeam.RemoveTeamMemberRequest) (*pTeam.StatusResponse, error) {
	return s.teamClient.RemoveTeamMember(ctx, req)
}

func (s *ServiceRepositoryClient) InviteMember(ctx context.Context, req *pTeam.InviteMemberRequest) (*pTeam.InvitationsResponse, error) {
	return s.teamClient.InviteMember(ctx, req)
}

func (s *ServiceRepositoryClient) RespondInvite(ctx context.Context, req *pTeam.RespondInviteRequest) (*pTeam.InvitationsResponse, error) {
	return s.teamClient.RespondInvite(ctx, req)
}

//! ------------------- Scoring -------------------

func (s *ServiceRepositoryClient) GiveScore(ctx context.Context, req *pScoring.GiveScoreRequest) (*pScoring.Score, error) {
	return s.scoringClient.GiveScore(ctx, req)
}

func (s *ServiceRepositoryClient) GetScoresByEvent(ctx context.Context, req *pScoring.GetScoresByEventRequest) (*pScoring.ScoreList, error) {
	return s.scoringClient.GetScoresByEvent(ctx, req)
}

func (s *ServiceRepositoryClient) GetScoresByUser(ctx context.Context, req *pScoring.GetScoresByUserRequest) (*pScoring.ScoreList, error) {
	return s.scoringClient.GetScoresByUser(ctx, req)
}

func (s *ServiceRepositoryClient) GetScoresByTeam(ctx context.Context, req *pScoring.GetScoresByTeamRequest) (*pScoring.ScoreList, error) {
	return s.scoringClient.GetScoresByTeam(ctx, req)
}

func (s *ServiceRepositoryClient) GetGlobalRanking(ctx context.Context, req *pScoring.GetGlobalRankingRequest) (*pScoring.RankingList, error) {
	return s.scoringClient.GetGlobalRanking(ctx, req)
}

//! ------------------- Notification -------------------

func (s *ServiceRepositoryClient) SendNotification(ctx context.Context, req *pNotif.SendNotificationRequest) (*pNotif.SendNotificationResponse, error) {
	return s.notifClient.SendNotification(ctx, req)
}