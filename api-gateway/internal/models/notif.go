package models 

type Notification struct {
	ID         string `json:"id"`
	SenderID   string `json:"sender_id"`
	SenderType int32  `json:"sender_type"` 
	UserID     string `json:"user_id,omitempty"`
	IsPublic   bool   `json:"is_public"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Type       int32  `json:"type"`       
	CreatedAt  string `json:"created_at"` // ISO format
}

type SendNotificationRequest struct {
	SenderID   string `json:"sender_id" binding:"required"`
	SenderType int32  `json:"sender_type" binding:"required"` 
	UserID     string `json:"user_id,omitempty"`
	IsPublic   bool   `json:"is_public"`
	Title      string `json:"title" binding:"required"`
	Body       string `json:"body" binding:"required"`
	Type       int32  `json:"type" binding:"required"`        // int32 instead of enum
}

type SendNotificationResponse struct {
	Notification Notification `json:"notification"`
}
