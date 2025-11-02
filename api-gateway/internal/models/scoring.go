package models

// ================ ENUMS =================

type ScoredByType string

const (
	ScoredByOrganizer ScoredByType = "ORGANIZER"
	ScoredByAdmin     ScoredByType = "ADMIN"
)

// ================ MODELS =================

type Score struct {
	ID           string       `json:"id"`
	EventID      string       `json:"event_id"`
	TeamID       string      `json:"team_id,omitempty"` // nullable
	UserID       string       `json:"user_id"`
	Points       int32        `json:"points"`
	Comment      string      `json:"comment,omitempty"` // nullable
	ScoredByID   string       `json:"scored_by_id"`
	ScoredByType ScoredByType `json:"scored_by_type"`
	CreatedAt    string       `json:"created_at"`
}

// ================ REQUESTS =================

type GiveScoreRequest struct {
	EventID      string       `json:"event_id" binding:"required"`
	TeamID       string      `json:"team_id,omitempty"`
	UserID       string       `json:"user_id" binding:"required"`
	Points       int32        `json:"points" binding:"required"`
	Comment      string      `json:"comment,omitempty"`
	ScoredByID   string       `json:"scored_by_id" binding:"required"`
	ScoredByType ScoredByType `json:"scored_by_type" binding:"required"`
}

type GetScoresByEventRequest struct {
	EventID string `json:"event_id" binding:"required"`
}

type GetScoresByUserRequest struct {
	UserID string `json:"user_id" binding:"required"`
}

type GetScoresByTeamRequest struct {
	TeamID string `json:"team_id" binding:"required"`
}

type GetGlobalRankingRequest struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}

// ================ RESPONSES =================

type ScoreList struct {
	Scores     []Score `json:"scores"`
	TotalCount int32   `json:"total_count"`
}

type Ranking struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	TotalPoints int32  `json:"total_points"`
	Rank        int32  `json:"rank"`
}

type RankingList struct {
	Rankings   []Ranking `json:"rankings"`
	TotalCount int32     `json:"total_count"`
}
