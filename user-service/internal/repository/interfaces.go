package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

func NewUserRepository(repo *UserREPO, log *logrus.Logger) IUserRepository{
	return &UserREPO{
		queries: repo.queries,
		log: log,
	}
}

type IUserRepository interface {
	//User
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error)
	GetUserByIdentifier(ctx context.Context, req *pb.GetUserByIdentifierRequest) (*pb.LoginResponse, error)
	ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error)
	GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error)
	ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UserList, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)

	//User Roles
	AssignRoleToUser(ctx context.Context, req *pb.AssignRoleRequest) (*pb.UserRole, error)
	RemoveRoleFromUser(ctx context.Context, req *pb.RemoveRoleRequest) (*pb.StatusResponse, error)
	ListUserRoles(ctx context.Context, req *pb.ListUserRolesRequest) (*pb.UserRoleList, error)

	//Roles
	CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleType, error)
	GetRoleById(ctx context.Context, req *pb.GetRoleByIdRequest) (*pb.RoleType, error)
	UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.RoleType, error)
	ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.RoleTypeList, error)
	DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.StatusResponse, error)
}
