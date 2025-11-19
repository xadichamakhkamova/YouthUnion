package repository

import (
	"context"
	"database/sql"
	"notification-service/internal/storage"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
)

type NotifRepo struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewNotifStore(db *sql.DB, log *logrus.Logger) *NotifRepo {
	return &NotifRepo{
		queries: storage.New(db),
		log:     log,
	}
}

func (r *NotifRepo) CreateNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {

	sender_id, err := uuid.Parse(req.SenderId)
	if err != nil {
		return nil, err
	}
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	resp, err := r.queries.CreateNotification(ctx, storage.CreateNotificationParams{
		SenderID:   sender_id,
		SenderType: storage.SenderType(req.SenderType.String()),
		UserID:     uuid.NullUUID{UUID: user_id, Valid: true},
		IsPublic:   req.IsPublic,
		Title:      req.Title,
		Body:       req.Body,
		Type:       storage.NotificationType(req.Type.String()),
	})
	if err != nil {
		return nil, err
	}
	return &pb.SendNotificationResponse{
		Notification: &pb.Notification{
			Id:         resp.ID.String(),
			SenderId:   resp.SenderID.String(),
			SenderType: pb.SenderType(pb.SenderType_value[string(resp.SenderType)]),
			UserId:     resp.UserID.UUID.String(),
			IsPublic:   resp.IsPublic,
			Title:      resp.Title,
			Body:       resp.Body,
			Type:       pb.NotificationType(pb.NotificationType_value[string(resp.Type)]),
			CreatedAt:  resp.CreatedAt.Time.String(),
		},
	}, nil
}
