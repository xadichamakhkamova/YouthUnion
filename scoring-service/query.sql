-- name: GiveScore :one 
INSERT INTO scores (
    event_id,
    team_id,
    user_id,
    points,
    comment,
    scored_by_id,
    scored_by_type
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING 
    id,
    event_id,
    team_id,
    user_id,
    points,
    comment,
    scored_by_id,
    scored_by_type,
    created_at; 

-- name: GetScoreByEvent :many
SELECT 
    id,
    event_id,
    team_id,
    user_id,
    points,
    comment,
    scored_by_id,
    scored_by_type,
    created_at,
    COUNT(*) OVER() AS total_count
FROM scores 
WHERE event_id = $1
ORDER BY created_at DESC 
LIMIT $2 
OFFSET $3; 

-- name: GetScoreByUser :many
SELECT 
    id,
    event_id,
    team_id,
    user_id,
    points,
    comment,
    scored_by_id,
    scored_by_type,
    created_at,
    COUNT(*) OVER() AS total_count
FROM scores 
WHERE user_id = $1
ORDER BY created_at DESC 
LIMIT $2 
OFFSET $3; 

-- name: GetScoreByTeam :many
SELECT 
    id,
    event_id,
    team_id,
    user_id,
    points,
    comment,
    scored_by_id,
    scored_by_type,
    created_at,
    COUNT(*) OVER() AS total_count
FROM scores 
WHERE team_id = $1
ORDER BY created_at DESC 
LIMIT $2 
OFFSET $3; 

-- name: GetGlobalRanking :many
SELECT
    user_id,
    SUM(points) AS total_points
FROM scores
GROUP BY user_id
ORDER BY total_points DESC
LIMIT $1 
OFFSET $2;
