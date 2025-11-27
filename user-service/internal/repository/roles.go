package repository

import (
	"context"
	"database/sql"
	"user-service/internal/storage"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

//! ------------------- User Roles -------------------

func (r *UserREPO) AssignRoleToUser(ctx context.Context, req *pb.AssignRoleRequest) (*pb.UserRole, error) {
	r.log.Info("Assigning role to user started")
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		r.log.WithError(err).Error("Invalid user UUID format")
		return nil, err
	}
	roleID, err := uuid.Parse(req.RoleId)
	if err != nil {
		r.log.WithError(err).Error("Invalid role UUID format")
		return nil, err
	}

	resp, err := r.queries.AssignRoleToUser(ctx, storage.AssignRoleToUserParams{
		UserID: userID,
		RoleID: roleID,
	})
	if err != nil {
		r.log.WithError(err).Error("Failed to assign role in database")
		return nil, err
	}

	r.log.WithField("assigned_role_id", resp.ID.String()).Info("Role assigned successfully")

	return &pb.UserRole{
		Id:         resp.ID.String(),
		UserId:     resp.UserID.String(),
		RoleId:     resp.RoleID.String(),
		AssignedAt: resp.AssignedAt.Time.String(),
	}, nil
}

func (r *UserREPO) RemoveRoleFromUser(ctx context.Context, req *pb.RemoveRoleRequest) (*pb.StatusResponse, error) {
	r.log.Info("Removing role from user started")

	id, err := uuid.Parse(req.UserRoleId)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	message, err := r.queries.RemoveRoleFromUser(ctx, id)
	if err != nil {
		r.log.WithError(err).Error("Failed to remove role from user in database")
		return nil, err
	}

	is_success := false
	if message == "removed" {
		is_success = true
		r.log.Info("Role removed successfully")
	} else {
		r.log.Warn("Role removal unsuccessful")
	}

	return &pb.StatusResponse{
		Success: is_success,
		Message: message,
	}, nil
}

func (r *UserREPO) ListUserRoles(ctx context.Context, req *pb.ListUserRolesRequest) (*pb.UserRoleList, error) {
	r.log.Info("Listing user roles started")
	offset := (req.Page - 1) * req.Limit

	id, err := uuid.Parse(req.UserId)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	params := storage.ListUserRolesParams{
		UserID: id,
		Limit:  req.Limit,
		Offset: offset,
	}

	resp, err := r.queries.ListUserRoles(ctx, params)
	if err != nil {
		r.log.WithError(err).Error("Failed to fetch user roles from database")
		return nil, err
	}

	var userRoles []*pb.UserRole
	var totalCount int64
	for _, r := range resp {
		userRoles = append(userRoles, &pb.UserRole{
			Id:         r.ID.String(),
			UserId:     r.UserID.String(),
			RoleId:     r.RoleName,
			AssignedAt: r.AssignedAt.Time.String(),
		})
		totalCount = r.TotalCount
	}

	r.log.WithFields(logrus.Fields{
		"count": totalCount,
	}).Info("User roles listed successfully")

	return &pb.UserRoleList{
		UserRoles:  userRoles,
		TotalCount: int32(totalCount),
	}, nil
}

//! ------------------- Role CRUD -------------------

func (r *UserREPO) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.RoleType, error) {
	r.log.Info("Creating role started")

	resp, err := r.queries.CreateRole(ctx, storage.CreateRoleParams{
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
	})
	if err != nil {
		r.log.WithError(err).Error("Failed to create role in database")
		return nil, err
	}

	r.log.WithField("role_id", resp.ID.String()).Info("Role created successfully")

	return &pb.RoleType{
		Id:          resp.ID.String(),
		Name:        resp.Name,
		Description: resp.Description.String,
		CreatedAt:   resp.CreatedAt.Time.String(),
		UpdatedAt:   resp.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) GetRoleById(ctx context.Context, req *pb.GetRoleByIdRequest) (*pb.RoleType, error) {
	r.log.Info("Fetching role by ID started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	resp, err := r.queries.GetRoleByID(ctx, id)
	if err != nil {
		r.log.WithError(err).Error("Failed to fetch role from database")
		return nil, err
	}

	r.log.WithField("role_name", resp.Name).Info("Role fetched successfully")

	return &pb.RoleType{
		Id:          resp.ID.String(),
		Name:        resp.Name,
		Description: resp.Description.String,
		CreatedAt:   resp.CreatedAt.Time.String(),
		UpdatedAt:   resp.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.RoleType, error) {
	r.log.Info("Updating role started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	resp, err := r.queries.UpdateRole(ctx, storage.UpdateRoleParams{
		ID:          id,
		Name:        req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
	})
	if err != nil {
		r.log.WithError(err).Error("Failed to update role in database")
		return nil, err
	}

	r.log.WithField("role_id", resp.ID.String()).Info("Role updated successfully")

	return &pb.RoleType{
		Id:          resp.ID.String(),
		Name:        resp.Name,
		Description: resp.Description.String,
		CreatedAt:   resp.CreatedAt.Time.String(),
		UpdatedAt:   resp.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) ListRoles(ctx context.Context, req *pb.ListRolesRequest) (*pb.RoleTypeList, error) {
	r.log.Info("Listing roles started")
	offset := (req.Page - 1) * req.Limit

	params := storage.ListRolesParams{
		Limit:  req.Limit,
		Offset: offset,
	}

	resp, err := r.queries.ListRoles(ctx, params)
	if err != nil {
		r.log.WithError(err).Error("Failed to fetch role list from database")
		return nil, err
	}

	var roles []*pb.RoleType
	var totalCount int64
	for _, r := range resp {
		roles = append(roles, &pb.RoleType{
			Id:          r.ID.String(),
			Name:        r.Name,
			Description: r.Description.String,
			CreatedAt:   r.CreatedAt.Time.String(),
			UpdatedAt:   r.UpdatedAt.Time.String(),
		})
		totalCount = r.TotalCount
	}

	r.log.WithField("total", totalCount).Info("Roles listed successfully")

	return &pb.RoleTypeList{
		Roles:      roles,
		TotalCount: int32(totalCount),
	}, nil
}

func (r *UserREPO) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.StatusResponse, error) {
	r.log.Info("Deleting role started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	message, err := r.queries.DeleteRole(ctx, id)
	if err != nil {
		r.log.WithError(err).Error("Failed to delete role from database")
		return nil, err
	}

	is_success := false
	if message == "deleted" {
		is_success = true
		r.log.Info("Role deleted successfully")
	} else {
		r.log.Warn("Role deletion unsuccessful")
	}

	return &pb.StatusResponse{
		Success: is_success,
		Message: message,
	}, nil
}
