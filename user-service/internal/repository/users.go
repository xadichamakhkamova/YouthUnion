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
	
	logger := r.log.WithFields(logrus.Fields{
		"method": "GetUserById",
		"user_id": req.Id,
	})

	logger.Info("Fetching user by ID started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	user, err := r.queries.GetUserById(ctx, id)
	if err != nil {
		logger.WithError(err).Error("Database query failed")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"identifier": user.Identifier,
	}).Info("User fetched successfully")

	return &pb.User{
		Id:           req.Id,
		Identifier:   user.Identifier,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhoneNumber:  user.PhoneNumber,
		Faculty:      user.Faculty.String,
		Course:       int32(user.Course.Int16),
		BirthDate:    user.BirthDate,
		Gender:       pb.Gender(pb.Gender_value[string(user.Gender)]),
		CreatedAt:    user.CreatedAt.Time.String(),
		UpdatedAt:    user.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {

	logger := r.log.WithFields(logrus.Fields{
		"method":  "UpdateUser",
		"user_id": req.Id,
	})

	logger.Info("Updating user started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	user, err := r.queries.UpdateUser(ctx, storage.UpdateUserParams{
		ID:          id,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Faculty:     sql.NullString{String: req.Faculty, Valid: req.Faculty != ""},
		Course:      sql.NullInt16{Int16: int16(req.Course), Valid: req.Course != 0},
		BirthDate:   req.BirthDate,
		Gender:      storage.GenderEnum(req.Gender.String()),
	})
	if err != nil {
		logger.WithError(err).Error("Failed to update user in database")
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"identifier": user.Identifier,
	}).Info("User updated successfully")

	return &pb.User{
		Id:           user.ID.String(),
		Identifier:   user.Identifier,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PhoneNumber:  user.PhoneNumber,
		Faculty:      user.Faculty.String,
		Course:       int32(user.Course.Int16),
		BirthDate:    user.BirthDate,
		Gender:       pb.Gender(pb.Gender_value[string(user.Gender)]),
		CreatedAt:    user.CreatedAt.Time.String(),
		UpdatedAt:    user.UpdatedAt.Time.String(),
	}, nil
}

func (r *UserREPO) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.UserList, error) {

	logger := r.log.WithFields(logrus.Fields{
		"method": "ListUsers",
		"page":   req.Page,
		"limit":  req.Limit,
	})

	logger.Info("Listing users started")

	params := storage.ListUsersParams{
		Column1: req.Page,
		Limit:   req.Limit,
	}

	resp, err := r.queries.ListUsers(ctx, params)
	if err != nil {
		logger.WithError(err).Error("Failed to fetch user list")
		return nil, err
	}

	var users []*pb.User
	var total_count int64
	for _, r := range resp {
		users = append(users, &pb.User{
			Id:           r.ID.String(),
			Identifier:   r.Identifier,
			FirstName:    r.FirstName,
			LastName:     r.LastName,
			PhoneNumber:  r.PhoneNumber,
			Faculty:      r.Faculty.String,
			Course:       int32(r.Course.Int16),
			BirthDate:    r.BirthDate,
			Gender:       pb.Gender(pb.Gender_value[string(r.Gender)]),
			CreatedAt:    r.CreatedAt.Time.String(),
			UpdatedAt:    r.UpdatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}

	logger.WithFields(logrus.Fields{
		"user_count": len(users),
		"total":      total_count,
	}).Info("User list retrieved successfully")

	return &pb.UserList{
		Users:      users,
		TotalCount: int32(total_count),
	}, nil
}

func (r *UserREPO) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	logger := r.log.WithFields(logrus.Fields{
		"method":  "DeleteUser",
		"user_id": req.Id,
	})

	logger.Info("Deleting user started")

	id, err := uuid.Parse(req.Id)
	if err != nil {
		logger.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	message, err := r.queries.DeleteUser(ctx, storage.DeleteUserParams{
		ID:        id,
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	})
	if err != nil {
		logger.WithError(err).Error("Failed to delete user from database")
		return nil, err
	}

	status := 400
	if message == "deleted" {
		status = 204
		logger.WithField("status", status).Info("User deleted successfully")
	} else {
		logger.WithField("status", status).Warn("User deletion unsuccessful")
	}

	return &pb.DeleteUserResponse{
		Status:        int32(status),
		DeletedUserId: req.Id,
	}, nil
}
