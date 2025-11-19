-- name: CreateNotification :one
INSERT INTO notifications (
    sender_id,
    sender_type,
    user_id,
    is_public,
    title,
    body,
    type
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    id,
    sender_id,
    sender_type,
    user_id,
    is_public,
    title,
    body,
    type,
    created_at; 

