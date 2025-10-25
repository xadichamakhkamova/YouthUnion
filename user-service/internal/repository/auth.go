package repository

import (
	"context"
	"database/sql"
	"time"
	"user-service/internal/hash"
	"user-service/internal/storage"

	"github.com/google/uuid"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//! ------------------- Authorization -------------------

func (q *UserREPO) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	q.log.Info("CreateUser started")

	passwordHash, err := hash.HashPassword(req.PasswordHash)
	if err != nil {
		q.log.WithError(err).Error("Failed to hash password")
		return nil, err
	}

	user, err := q.queries.CreateUser(ctx, storage.CreateUserParams{
		Identifier:   req.Identifier,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		PasswordHash: passwordHash,
		Faculty:      sql.NullString{String: req.Faculty, Valid: req.Faculty != ""},
		Course:       sql.NullInt16{Int16: int16(req.Course), Valid: req.Course != 0},
		BirthDate:    req.BirthDate,
		Gender:       storage.GenderEnum(req.Gender),
	})
	if err != nil {
		q.log.WithError(err).Error("Failed to create user in database")
		return nil, err
	}

	q.log.WithField("user_id", user.ID.String()).Info("User created successfully")

	return &pb.User{
		Id:          user.ID.String(),
		Identifier:  user.Identifier,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Faculty:     user.Faculty.String,
		Course:      int32(user.Course.Int16),
		BirthDate:   user.BirthDate,
		Gender:      string(user.Gender),
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}, nil
}

func (q *UserREPO) GetUserByIdentifier(ctx context.Context, req *pb.GetUserByIdentifierRequest) (*pb.GetUserByIdentifierResponse, error) {
	q.log.Info("Authentication attempt started")

	user, err := q.queries.GetUserByIdentifier(ctx, req.Identifier)
	if err != nil {
		q.log.WithError(err).Error("User not found in database")
		return nil, err
	}

	// Parolni tekshirish
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.PasswordHash)); err != nil {
		q.log.Warn("Invalid credentials provided")
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	q.log.WithField("user_id", user.ID.String()).Info("Authentication successful")

	return &pb.GetUserByIdentifierResponse{
		Status: 200,
	}, nil
}

func (q *UserREPO) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	q.log.Info("Password change process started")

	id, err := uuid.Parse(req.UserId)
	if err != nil {
		q.log.WithError(err).Error("Invalid UUID format")
		return nil, err
	}

	newPasswordHash, err := hash.HashPassword(req.NewPassword)
	if err != nil {
		q.log.WithError(err).Error("Failed to hash new password")
		return nil, err
	}

	message, err := q.queries.ChangePassword(ctx, storage.ChangePasswordParams{
		ID:           id,
		PasswordHash: newPasswordHash,
	})
	if err != nil {
		q.log.WithError(err).Error("Database error while changing password")
		return nil, err
	}

	status := 400
	if message == "changed" {
		status = 204
		q.log.WithField("status", status).Info("Password changed successfully")
	} else {
		q.log.WithField("status", status).Warn("Password not changed")
	}

	return &pb.ChangePasswordResponse{
		Status:    int32(status),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
