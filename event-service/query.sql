-- name: CreateEvent :one 
INSERT INTO events (
    event_type,
    title,
    description,
    location,
    start_time,
    end_time,
    created_by,
    max_participants
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING 
    id,
    event_type,
    title,
    description,
    location,
    start_time,
    end_time,
    created_by,
    max_participants,
    status,
    created_at,
    updated_at; 

-- name: UpdateEvent :one 
UPDATE events 
SET
    title = $2,
    description = $3,
    location = $4,
    start_time = $5,
    end_time = $6,
    max_participants = $7,
    status = $8
WHERE id = $1 AND deleted_at = 0
RETURNING
    id,
    event_type,
    title,
    description,
    location,
    start_time,
    end_time,
    created_by,
    max_participants,
    status,
    created_at,
    updated_at; 

-- name: GetEvent :one 
SELECT 
    id,
    event_type,
    title,
    description,
    location,
    start_time,
    end_time,
    created_by,
    max_participants,
    status,
    created_at,
    updated_at 
FROM events 
WHERE id = $1 AND deleted_at = 0; 
 
-- name: ListEvents :many
SELECT 
    id,
    event_type,
    title,
    description,
    location,
    start_time,
    end_time,
    created_by,
    max_participants,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM events 
WHERE (
        $1::text = ''
        OR LOWER(title) LIKE LOWER(CONCAT('%', $1::text, '%'))  -- name bo‘yicha qidirish
        OR LOWER(event_type::text) LIKE LOWER(CONCAT('%', $1::text, '%'))
        OR LOWER(status) LIKE LOWER(CONCAT('%', $1::text, '%'))
    ) 
AND deleted_at = 0
ORDER BY created_at DESC 
LIMIT $2 
OFFSET $3;

-- name: DeleteEvent :one
UPDATE events
SET deleted_at = $2
WHERE id = $1
RETURNING 'deleted' AS message;

