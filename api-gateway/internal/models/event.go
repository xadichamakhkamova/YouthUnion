package models

// ================= MODELS =================

type Event struct {
	ID                  string `json:"id"`
	EventType           int32  `json:"event_type"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	Location            string `json:"location"`
	StartTime           string `json:"start_time"`
	EndTime             string `json:"end_time"`
	CreatedBy           string `json:"created_by"`
	MaxParticipants     int32  `json:"max_participants"`
	CurrentParticipants int32  `json:"current_participants"`
	Status              int32  `json:"status"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

type EventParticipant struct {
	ID       string `json:"id"`
	EventID  string `json:"event_id"`
	UserID   string `json:"user_id"`
	Role     string `json:"role"`
	JoinedAt string `json:"joined_at"`
}

// ================= REQUESTS =================

type CreateEventRequest struct {
	EventType       int32  `json:"event_type" binding:"required"`
	Title           string `json:"title" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Location        string `json:"location" binding:"required"`
	StartTime       string `json:"start_time" binding:"required"`
	EndTime         string `json:"end_time" binding:"required"`
	CreatedBy       string `json:"created_by" binding:"required"`
	MaxParticipants int32  `json:"max_participants" binding:"required"`
}

type UpdateEventRequest struct {
	ID              string `json:"id" binding:"required"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Location        string `json:"location"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	MaxParticipants int32  `json:"max_participants"`
	Status          int32  `json:"status"`
}

type GetEventRequest struct {
	ID string `json:"id" binding:"required"`
}

type ListEventsRequest struct {
	Search    string `json:"search"`
	EventType int32  `json:"event_type"`
	Status    int32  `json:"status"`
	Limit     int32  `json:"limit"`
	Page      int32  `json:"page"`
}

type DeleteEventRequest struct {
	ID string `json:"id" binding:"required"`
}

type DeleteEventResponse struct {
	Status int32 `json:"status"`
}

// ========== EVENT REGISTRATION ==========

type RegisterEventRequest struct {
	EventID string `json:"event_id" binding:"required"`
	UserID  string `json:"user_id" binding:"required"` // individual user yoki leader
}

type RegisterTeamEventRequest struct {
	EventID string `json:"event_id" binding:"required"`
	TeamID  string `json:"team_id" binding:"required"`
}

type EventParticipantResponse struct {
	Participants []EventParticipant `json:"participants"`
	TotalCount   int32              `json:"total_count"`
}

// ================= RESPONSE FOR LIST =================

type ListEventsResponse struct {
	Events     []Event `json:"events"`
	TotalCount int32   `json:"total_count"`
}

