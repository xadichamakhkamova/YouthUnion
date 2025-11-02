package models

// ================= ENUMS =================

type MemberRole string

const (
	RoleLeader MemberRole = "LEADER"
	RoleMember MemberRole = "MEMBER"
)

type InviteStatus string

const (
	InvitePending  InviteStatus = "PENDING"
	InviteAccepted InviteStatus = "ACCEPTED"
	InviteRejected InviteStatus = "REJECTED"
)

// ================= MODELS =================

type Team struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LeaderID  string `json:"leader_id"`
	EventID   string `json:"event_id"`
	IsReady   bool   `json:"is_ready"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TeamMember struct {
	ID       string     `json:"id"`
	TeamID   string     `json:"team_id"`
	UserID   string     `json:"user_id"`
	Role     MemberRole `json:"role"`
	JoinedAt string     `json:"joined_at"`
}

// ================= REQUESTS =================

// --- Team CRUD ---

type CreateTeamRequest struct {
	Name     string `json:"name" binding:"required"`
	LeaderID string `json:"leader_id" binding:"required"`
	EventID  string `json:"event_id" binding:"required"`
}

type UpdateTeamRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type GetTeamRequest struct {
	TeamID string `json:"team_id" binding:"required"`
}

type GetTeamsByEventRequest struct {
	EventID string `json:"event_id" binding:"required"`
	Limit   int32  `json:"limit"`
	Page    int32  `json:"page"`
}

type RemoveTeamMemberRequest struct {
	TeamID string `json:"team_id" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
}

// --- Invitations & Membership ---

type InviteMemberRequest struct {
	TeamID        string `json:"team_id" binding:"required"`
	InviterID     string `json:"inviter_id" binding:"required"`
	InvitedUserID string `json:"invited_user_id" binding:"required"`
}

type RespondInviteRequest struct {
	TeamID        string       `json:"team_id" binding:"required"`
	InvitedUserID string       `json:"invited_user_id" binding:"required"`
	Status        InviteStatus `json:"status" binding:"required"`
}

// ================= RESPONSES =================

type TeamList struct {
	Teams      []Team `json:"teams"`
	TotalCount int32  `json:"total_count"`
}

type MemberList struct {
	Members    []TeamMember `json:"members"`
	TotalCount int32        `json:"total_count"`
}

type StatusResponse struct {
	Status int32 `json:"status"`
}
