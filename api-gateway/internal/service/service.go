package service

import (
	"context"

	pUser "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

type ServiceRepositoryClient struct {
	userClient pUser.UserServiceClient
}

func NewServiceRepositoryClient(
	conn1 *pUser.UserServiceClient,
) *ServiceRepositoryClient {
	return &ServiceRepositoryClient{
		userClient: *conn1,
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
