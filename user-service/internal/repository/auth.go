package repository

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/hash"
	"user-service/internal/storage"

	"github.com/google/uuid"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/userpb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		Course:       sql.NullInt32{Int32: int32(req.Course), Valid: req.Course != 0},
		BirthDate:    req.BirthDate,
		Gender:       storage.Gender(req.Gender),
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
		Course:      user.Course.Int32,
		BirthDate:   user.BirthDate,
		Gender:      string(user.Gender),
		CreatedAt:   user.CreatedAt.Time.String(),
		UpdatedAt:   user.UpdatedAt.Time.String(),
	}, nil
}

func (q *UserREPO) GetUserByIdentifier(ctx context.Context, req *pb.GetUserByIdentifierRequest) (*pb.LoginResponse, error) {
	q.log.Info("Authentication attempt started")

	user, err := q.queries.GetUserByIdentifier(ctx, req.Identifier)
	if err != nil {
		q.log.WithError(err).Error("User not found in database")
		return nil, err
	}

	// Parolni tekshirish
	fmt.Println(req.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		q.log.Warn("Invalid credentials provided")
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	q.log.WithField("user_id", user.ID.String()).Info("Authentication successful")

	return &pb.LoginResponse{
		User: &pb.User{
			Id:          user.ID.String(),
			Identifier:  user.Identifier,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
		},
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

	resp, err := q.queries.ChangePassword(ctx, storage.ChangePasswordParams{
		ID:           id,
		PasswordHash: newPasswordHash,
	})
	if err != nil {
		q.log.WithError(err).Error("Database error while changing password")
		return nil, err
	}

	is_success := false
	if resp.Message == "changed" {
		is_success = true
		q.log.Info("Password changed successfully")
	} else {
		q.log.Warn("Password not changed")
	}

	return &pb.ChangePasswordResponse{
		Success:   is_success,
		Message:   resp.Message,
		UpdatedAt: resp.UpdatedAt.Time.String(),
	}, nil
}
