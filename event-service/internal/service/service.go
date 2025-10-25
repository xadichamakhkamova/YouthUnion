package service

import (
	"context"
	"event-service/internal/repository"

	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/eventpb"
)

type EventService struct {
	pb.UnimplementedEventServiceServer
	repo repository.IEventRepository
}

func NewEventService (repo repository.IEventRepository) *EventService {
	return &EventService{
		repo: repo,
	}
}

//! ------------------- Event -------------------

func (s *EventService) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.Event, error) {
	return s.repo.CreateEvent(ctx, req)
}

func (s *EventService) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.Event, error) {
	return s.repo.UpdateEvent(ctx, req)
}

func (s *EventService) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.Event, error) {
	return s.repo.GetEvent(ctx, req)
}

func (s *EventService) ListEvents(ctx context.Context, req *pb.ListEventsRequest) (*pb.ListEventsResponse, error) {
	return s.repo.ListEvents(ctx, req)
}

func (s *EventService) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	return s.repo.DeleteEvent(ctx, req)
}
