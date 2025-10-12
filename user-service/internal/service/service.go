package service

import (
	"context"
	"user-service/internal/repository"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo repository.IUserRepository
}

func NewUserService (repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

//! ------------------- Authorization -------------------

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	return s.repo.CreateUser(ctx, req)
}

func (s *UserService) GetUserByIdentifier(ctx context.Context, req *pb.GetUserByIdentifierRequest) (*pb.GetUserByIdentifierResponse, error) {
	return s.repo.GetUserByIdentifier(ctx, req)
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	return s.repo.ChangePassword(ctx, req)
}

//! ------------------ User Functions -------------------

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	return s.repo.GetUserById(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	return s.repo.UpdateUser(ctx, req)
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UserList, error) {
	return s.repo.ListUsers(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.repo.DeleteUser(ctx, req)
}

//! ------------------- User Roles -------------------

func (s *UserService) AssignRoleToUser(ctx context.Context, req *pb.AssignRoleRequest) (*pb.UserRole, error) {
	return s.repo.AssignRoleToUser(ctx, req)
}

func (s *UserService) RemoveRoleFromUser(ctx context.Context, req *pb.RemoveRoleRequest) (*pb.RemoveRoleResponse, error) {
	return s.repo.RemoveRoleFromUser(ctx, req)
}

func (s *UserService) ListUserRoles(ctx context.Context, req *pb.ListUserRolesRequest) (*pb.UserRoleList, error) {
	return s.repo.ListUserRoles(ctx, req)
}


//! ------------------- Role functions -------------------

func (s *UserService) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleType, error) {
	return s.repo.CreateRole(ctx, req)
}

func (s *UserService) GetRoleById(ctx context.Context, req *pb.GetRoleByIdRequest) (*pb.RoleType, error) {
	return s.repo.GetRoleById(ctx, req)
}

func (s *UserService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.RoleType, error) {
	return s.repo.UpdateRole(ctx, req)
}

func (s *UserService) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.RoleTypeList, error) {
	return s.repo.ListRoles(ctx, req)
}

func (s *UserService) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return s.repo.DeleteRole(ctx, req)
}
