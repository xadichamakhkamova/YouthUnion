package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/notificationpb"
)

func NewNotifRepository(repo *NotifRepo, log *logrus.Logger) INotifRepository {
	return &NotifRepo{
		queries: repo.queries,
		log:     log,
	}
}

type INotifRepository interface {
	CreateNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) 
}