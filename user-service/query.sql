-- name: CreateUser :one 
INSERT INTO users (
    identifier,
    first_name,
    last_name,
    phone_number,
    password_hash,
    faculty,
    course,
    birth_date,
    gender
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING 
    id,
    identifier,
    first_name,
    last_name,
    phone_number,
    faculty,
    course,
    birth_date,
    gender,
    created_at,
    updated_at;
    
-- name: GetUserByIdentifier :one
SELECT 
    id,
    identifier,
    first_name,
    last_name,
    phone_number,
    password_hash -- bu faqat backendda solishtirish uchun kerak
FROM users
WHERE identifier = $1;

-- name: GetUserById :one 
SELECT 
    id,
    identifier,
    first_name,
    last_name,
    phone_number,
    faculty,
    course,
    birth_date,
    gender,
    created_at,
    updated_at 
FROM users
WHERE id = $1;

-- name: UpdateUser :one 
UPDATE users 
SET 
    first_name = $2,
    last_name = $3,
    phone_number = $4,
    faculty = $5,
    course = $6,
    birth_date = $7,
    gender = $8,
    updated_at = NOW()
WHERE id = $1
RETURNING
    id,
    identifier,
    first_name,
    last_name,
    phone_number,
    faculty,
    course,
    birth_date,
    gender,
    created_at,
    updated_at;

-- name: ChangePassword :one 
UPDATE users
SET password_hash = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING 'changed' AS message;

-- name: ListUsers :many
SELECT
    id,
    identifier,
    first_name,
    last_name,
    phone_number,
    faculty,
    course,
    birth_date,
    gender,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM users
ORDER BY created_at DESC
LIMIT $2 
OFFSET ($1 - 1) * $2;

-- name: DeleteUser :one
UPDATE users
SET deleted_at = $2
WHERE id = $1
RETURNING 'deleted' AS message;


-- name: CreateSession :one
INSERT INTO user_sessions 
    (
        user_id, 
        refresh_token, 
        expires_at
    )
VALUES ($1, $2, $3)
RETURNING
    id,
    user_id,
    refresh_token,
    expires_at,
    created_at;

-- name: GetSessionByToken :one
SELECT
    id,
    user_id,
    refresh_token,
    expires_at,
    created_at
FROM user_sessions
WHERE refresh_token = $1;

-- name: CreateRole :one
INSERT INTO roles_type
    (
        name, 
        description
    )
VALUES ($1, $2)
RETURNING
    id,
    name,
    description,
    created_at,
    updated_at;

-- name: GetRoleByID :one
SELECT
    id,
    name,
    description,
    created_at,
    updated_at
FROM roles_type
WHERE id = $1;

-- name: UpdateRole :one
UPDATE roles_type
SET name = $2,
    description = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING
    id,
    name,
    description,
    created_at,
    updated_at;

-- name: ListRoles :many
SELECT
    id,
    name,
    description,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM roles_type
ORDER BY created_at DESC
LIMIT $2 
OFFSET ($1 - 1) * $2;

-- name: DeleteRole :one
DELETE FROM roles_type
WHERE id = $1
RETURNING 'deleted' AS message;;


-- name: AssignRoleToUser :one
INSERT INTO user_roles 
    (
        user_id, 
        role_id
    )
VALUES ($1, $2)
RETURNING
    id,
    user_id,
    role_id,
    assigned_at;

-- name: RemoveRoleFromUser :one
DELETE FROM user_roles
WHERE id = $1 
RETURNING 'removed' AS message;

-- name: ListUserRoles :many
SELECT
    ur.id,
    ur.user_id,
    rt.name AS role_name,
    ur.assigned_at,
    COUNT(*) OVER() AS total_count
FROM user_roles ur
JOIN roles_type rt ON ur.role_id = rt.id
WHERE ur.user_id = $1
LIMIT $3 
OFFSET ($2 - 1) * $3;
