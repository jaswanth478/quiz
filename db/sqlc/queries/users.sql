-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    username,
    user_type,
    metadata
) VALUES (
    sqlc.narg(email),
    sqlc.narg(password_hash),
    sqlc.narg(username),
    sqlc.narg(user_type),
    COALESCE(sqlc.narg(metadata), '{}'::jsonb)
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserLogin :exec
UPDATE users
SET last_login_at = NOW(),
    updated_at = NOW()
WHERE user_id = $1;

-- name: PromoteGuestToRegistered :one
UPDATE users
SET email = sqlc.arg(email),
    password_hash = sqlc.arg(password_hash),
    username = sqlc.arg(username),
    user_type = 'registered',
    updated_at = NOW()
WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = sqlc.arg(password_hash),
    updated_at = NOW()
WHERE user_id = sqlc.arg(user_id);

-- name: UpdateUsername :one
UPDATE users
SET username = sqlc.arg(username),
    updated_at = NOW()
WHERE user_id = sqlc.arg(user_id) AND username IS NULL
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

