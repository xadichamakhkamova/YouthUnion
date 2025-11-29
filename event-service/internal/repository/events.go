package repository

import (
	"context"
	"database/sql"
	"errors"
	"event-service/internal/storage"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	pb "github.com/xadichamakhkamova/YouthUnionContracts/genproto/eventpb"
)

type EventRepo struct {
	queries *storage.Queries
	log     *logrus.Logger
}

func NewEventStore(db *sql.DB, log *logrus.Logger) *EventRepo {
	return &EventRepo{
		queries: storage.New(db),
		log:     log,
	}
}

//! ------------------- Event Functions -------------------

func (r *EventRepo) CreateEvent(ctx context.Context, req *pb.CreateEventRequest) (*pb.Event, error) {

	start_time, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
	if err != nil {
		return nil, err
	}
	end_time, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
	if err != nil {
		return nil, err
	}
	created_by, err := uuid.Parse(req.CreatedBy)
	if err != nil {
		return nil, err
	}

	resp, err := r.queries.CreateEvent(ctx, storage.CreateEventParams{
		EventType:       storage.EventType(req.EventType.String()),
		Title:           req.Title,
		Description:     sql.NullString{String: req.Description, Valid: req.Description != ""},
		Location:        sql.NullString{String: req.Location, Valid: req.Location != ""},
		StartTime:       start_time,
		EndTime:         sql.NullTime{Time: end_time, Valid: true},
		CreatedBy:       created_by,
		MaxParticipants: sql.NullInt32{Int32: req.MaxParticipants},
	})
	if err != nil {
		return nil, err
	}

	return &pb.Event{
		Id:              resp.ID.String(),
		EventType:       pb.EventType(pb.EventType_value[string(resp.EventType)]),
		Title:           resp.Title,
		Description:     resp.Description.String,
		Location:        resp.Location.String,
		StartTime:       resp.StartTime.String(),
		EndTime:         resp.EndTime.Time.String(),
		CreatedBy:       resp.CreatedBy.String(),
		MaxParticipants: resp.MaxParticipants.Int32,
		Status:          pb.EventStatus(pb.EventStatus_value[resp.Status.String]),
		CreatedAt:       resp.CreatedAt.Time.String(),
		UpdatedAt:       resp.UpdatedAt.Time.String(),
	}, nil
}

func (r *EventRepo) UpdateEvent(ctx context.Context, req *pb.UpdateEventRequest) (*pb.Event, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	start_time, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
	if err != nil {
		return nil, err
	}
	end_time, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
	if err != nil {
		return nil, err
	}

	resp, err := r.queries.UpdateEvent(ctx, storage.UpdateEventParams{
		ID:              id,
		Title:           req.Title,
		Description:     sql.NullString{String: req.Description, Valid: req.Description != ""},
		Location:        sql.NullString{String: req.Location, Valid: req.Location != ""},
		StartTime:       start_time,
		EndTime:         sql.NullTime{Time: end_time, Valid: true},
		MaxParticipants: sql.NullInt32{Int32: req.MaxParticipants},
	})
	if err != nil {
		return nil, err
	}

	return &pb.Event{
		Id:              resp.ID.String(),
		EventType:       pb.EventType(pb.EventType_value[string(resp.EventType)]),
		Title:           resp.Title,
		Description:     resp.Description.String,
		Location:        resp.Location.String,
		StartTime:       resp.StartTime.String(),
		EndTime:         resp.EndTime.Time.String(),
		CreatedBy:       resp.CreatedBy.String(),
		MaxParticipants: resp.MaxParticipants.Int32,
		Status:          pb.EventStatus(pb.EventStatus_value[resp.Status.String]),
		CreatedAt:       resp.CreatedAt.Time.String(),
		UpdatedAt:       resp.UpdatedAt.Time.String(),
	}, nil
}

func (r *EventRepo) GetEvent(ctx context.Context, req *pb.GetEventRequest) (*pb.Event, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := r.queries.GetEvent(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.Event{
		Id:              resp.ID.String(),
		EventType:       pb.EventType(pb.EventType_value[string(resp.EventType)]),
		Title:           resp.Title,
		Description:     resp.Description.String,
		Location:        resp.Location.String,
		StartTime:       resp.StartTime.String(),
		EndTime:         resp.EndTime.Time.String(),
		CreatedBy:       resp.CreatedBy.String(),
		MaxParticipants: resp.MaxParticipants.Int32,
		Status:          pb.EventStatus(pb.EventStatus_value[resp.Status.String]),
		CreatedAt:       resp.CreatedAt.Time.String(),
		UpdatedAt:       resp.UpdatedAt.Time.String(),
	}, nil
}

func (r *EventRepo) ListEvents(ctx context.Context, req *pb.ListEventsRequest) (*pb.ListEventsResponse, error) {
	offset := (req.Page - 1) * req.Limit

	params := storage.ListEventsParams{
		Column1: req.Search,
		Limit:   req.Limit,
		Offset:  offset,
	}

	resp, err := r.queries.ListEvents(ctx, params)
	if err != nil {
		return nil, err
	}

	var events []*pb.Event
	var total_count int64
	for _, r := range resp {
		events = append(events, &pb.Event{
			Id:              r.ID.String(),
			EventType:       pb.EventType(pb.EventType_value[string(r.EventType)]),
			Title:           r.Title,
			Description:     r.Description.String,
			Location:        r.Location.String,
			StartTime:       r.StartTime.String(),
			EndTime:         r.EndTime.Time.String(),
			CreatedBy:       r.CreatedBy.String(),
			MaxParticipants: r.MaxParticipants.Int32,
			Status:          pb.EventStatus(pb.EventStatus_value[r.Status.String]),
			CreatedAt:       r.CreatedAt.Time.String(),
			UpdatedAt:       r.UpdatedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.ListEventsResponse{
		Events:     events,
		TotalCount: int32(total_count),
	}, nil
}

func (r *EventRepo) ListParticipants(ctx context.Context, req *pb.EventParticipantRequest) (*pb.EventParticipantResponse, error) {
	offset := (req.Page - 1) * req.Limit

	event_id, err := uuid.Parse(req.EventId)
	if err != nil {
		return nil, err
	}
	resp, err := r.queries.ListParticipants(ctx, storage.ListParticipantsParams{
		EventID: event_id,
		Limit:   req.Limit,
		Offset:  offset,
	})
	if err != nil {
		return nil, err
	}
	var participants []*pb.EventParticipant
	var total_count int64
	for _, r := range resp {
		participants = append(participants, &pb.EventParticipant{
			Id:       r.ID.String(),
			EventId:  r.EventID.String(),
			UserId:   r.UserID.String(),
			Role:     r.Role.String,
			JoinedAt: r.JoinedAt.Time.String(),
		})
		total_count = r.TotalCount
	}
	return &pb.EventParticipantResponse{
		Participants: participants,
		TotalCount: int32(total_count),
	}, nil
}

func (r *EventRepo) DeleteEvent(ctx context.Context, req *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {

	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	message, err := r.queries.DeleteEvent(ctx, storage.DeleteEventParams{
		ID:        id,
		DeletedAt: sql.NullInt64{Int64: time.Now().Unix(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	status := 400
	if message == "deleted" {
		status = 204
	} else {
		return nil, errors.New("Event deletion unsuccessful")
	}

	return &pb.DeleteEventResponse{
		Status: int32(status),
	}, nil
}

func (r *EventRepo) RegisterEvent(ctx context.Context, req *pb.RegisterEventRequest) (*pb.EventParticipant, error) {

	event_id, err := uuid.Parse(req.EventId)
	if err != nil {
		return nil, err
	}
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	resp, err := r.queries.RegisterEvent(ctx, storage.RegisterEventParams{
		EventID: event_id,
		UserID:  user_id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.EventParticipant{
		Id:       resp.ID.String(),
		EventId:  resp.EventID.String(),
		UserId:   resp.UserID.String(),
		Role:     resp.Role.String,
		JoinedAt: resp.JoinedAt.Time.String(),
	}, nil
}