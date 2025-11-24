package service

import (
	"context"
	"notification-service/internal/repository"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
)

type NotifService struct {
	pb.UnimplementedNotificationServiceServer
	repo repository.INotifRepository
	hub  WebSocketHub
}

func NewNotifService(repo repository.INotifRepository, hub WebSocketHub) *NotifService {
	return &NotifService{
		repo: repo,
		hub:  hub,
	}
}

func (s *NotifService) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	// Bazaga yozamiz
	notif, err := s.repo.CreateNotification(ctx, req)
	if err != nil {
		return nil, err
	}
	// WebSocketga yuborish
	if req.IsPublic {
		// hamma userlarga
		s.hub.BroadcastPublic(notif.Notification)
	} else {
		// faqat 1 ta userga
		s.hub.SendToUser(req.UserId, notif.Notification)
	}
	return notif, nil
}

type WebSocketHub interface {
	SendToUser(userID string, msg interface{})
	BroadcastPublic(msg interface{})
}
