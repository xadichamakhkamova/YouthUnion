package service

import (
	"context"
	"notification-service/internal/repository"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
)

type NotifService struct {
	pb.UnimplementedNotificationServiceServer
	repo repository.INotifRepository
}

func NewNotifService(repo repository.INotifRepository) *NotifService {
	return &NotifService{
		repo: repo,
	}
}

func (s *NotifService) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	return s.repo.CreateNotification(ctx, req)
}