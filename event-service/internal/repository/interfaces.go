package repository

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/eventpb"
)

func NewEventRepository(repo *EventRepo, log *logrus.Logger) IEventRepository {
	return &EventRepo{
		queries: repo.queries,
		log:     log,
	}
}

type IEventRepository interface {
	//Event
	CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.Event, error)
	UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.Event, error)
	GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.Event, error)
	ListEvents(ctx context.Context, req *pb.ListEventsRequest) (*pb.ListEventsResponse, error)
	DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error)

	RegisterEvent(ctx context.Context, req *pb.RegisterEventRequest) (*pb.EventParticipant, error)
	ListParticipants(ctx context.Context, req *pb.EventParticipantRequest) (*pb.EventParticipantResponse, error)
}
