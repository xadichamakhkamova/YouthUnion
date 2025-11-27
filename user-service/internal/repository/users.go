package repository

import (
	"context"
	"database/sql"
	"time"
	"user-service/internal/storage"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
)

type UserREPO struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewUserStore(db *sql.DB, log *logrus.Logger) *UserREPO {
	return &UserREPO{
		queries: storage.New(db),
		log:     log,
	}
}

//! ------------------- User Functions -------------------

func (r *UserREPO) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	r.log.Info("Fetching user by ID started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	user, err := r.queries.GetUserById(ctx, id)
	if err != nil {
		r.log.WithError(err).Error("Database query failed")
		return nil, err
	}

	r.log.WithFields(logrus.Fields{
		"identifier": user.Identifier,
	}).Info("User fetched successfully")

	return &pb.User{
		Id:          req.Id,
		Identifier:  user.Identifier,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Faculty:     user.Faculty.String,
		Course:      user.Course.Int32,
		BirthDate:   user.BirthDate,
		Gender:      string(user.Gender),
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	r.log.Info("Updating user started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	user, err := r.queries.UpdateUser(ctx, storage.UpdateUserParams{
		ID:          id,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Faculty:     sql.NullString{String: req.Faculty, Valid: req.Faculty != ""},
		Course:      sql.NullInt32{Int32: int32(req.Course), Valid: req.Course != 0},
		BirthDate:   req.BirthDate,
		Gender:      storage.Gender(req.Gender),
	})
	if err != nil {
		r.log.WithError(err).Error("Failed to update user in database")
		return nil, err
	}

	r.log.WithFields(logrus.Fields{
		"identifier": user.Identifier,
	}).Info("User updated successfully")

	return &pb.User{
		Id:          user.ID.String(),
		Identifier:  user.Identifier,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Faculty:     user.Faculty.String,
		Course:      user.Course.Int32,
		BirthDate:   user.BirthDate,
		Gender:      string(user.Gender),
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UserList, error) {
	r.log.Info("Listing users started")
	offset := (req.Page - 1) * req.Limit

	params := storage.ListUsersParams{
		Limit:  req.Limit,
		Offset: offset,
	}

	resp, err := r.queries.ListUsers(ctx, params)
	if err != nil {
		r.log.WithError(err).Error("Failed to fetch user list")
		return nil, err
	}

	var users []*pb.User
	var total_count int64
	for _, r := range resp {
		users = append(users, &pb.User{
			Id:          r.ID.String(),
			Identifier:  r.Identifier,
			FirstName:   r.FirstName,
			LastName:    r.LastName,
			PhoneNumber: r.PhoneNumber,
			Faculty:     r.Faculty.String,
			Course:      r.Course.Int32,
			BirthDate:   r.BirthDate,
			Gender:      string(r.Gender),
			CreatedAt:   r.CreatedAt.Time.String(),
			UpdatedAt:   r.UpdatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}

	r.log.WithFields(logrus.Fields{
		"user_count": len(users),
		"total":      total_count,
	}).Info("User list retrieved successfully")

	return &pb.UserList{
		Users:      users,
		TotalCount: int32(total_count),
	}, nil
}

func (r *UserREPO) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	r.log.Info("Deleting user started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		r.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	resp, err := r.queries.DeleteUser(ctx, storage.DeleteUserParams{
		ID:        id,
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	})
	if err != nil {
		r.log.WithError(err).Error("Failed to delete user from database")
		return nil, err
	}

	is_success := false
	if resp.Message == "deleted" {
		is_success = true
		r.log.Info("User deleted successfully")
	} else {
		r.log.Warn("User deletion unsuccessful")
	}

	return &pb.DeleteUserResponse{
		Success:       is_success,
		Message:       resp.Message,
		DeletedUserId: resp.ID.String(),
	}, nil
}
