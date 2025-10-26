-- name: CreateTeam :one
INSERT INTO teams (
    name,
    leader_id,
    event_id
)
VALUES ($1, $2, $3)
RETURNING
    id,
    name,
    leader_id,
    event_id,
    is_ready,
    created_at,
    updated_at;

-- name: UpdateTeam :one
UPDATE teams
SET 
    name = $2
WHERE id = $1
RETURNING
    id,
    name,
    leader_id,
    event_id,
    is_ready,
    created_at,
    updated_at;

-- name: GetTeamsByEvent :many
SELECT 
    id,
    name,
    leader_id,
    event_id,
    is_ready,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM teams
WHERE event_id = $1
ORDER BY created_at DESC
LIMIT $2 
OFFSET $3;

-- name: AddTeamMember :one
INSERT INTO team_members (
    team_id,
    user_id
)
VALUES ($1, $2)
RETURNING
    id,
    team_id,
    user_id,
    role,
    joined_at;

-- name: RemoveTeamMember :exec
DELETE FROM team_members
WHERE team_id = $1 AND user_id = $2;

-- name: GetTeamMembers :many
SELECT 
    id,
    team_id,
    user_id,
    role,
    joined_at,
    COUNT(*) OVER() AS total_count
FROM team_members
WHERE team_id = $1
ORDER BY joined_at DESC;