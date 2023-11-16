-- name: CreateUser :one
INSERT INTO users ( email, username, hash )
VALUES (
    @email,
    COALESCE(sqlc.narg(username)::text, 'i'),
    @hash
)
RETURNING id, username, email;

-- name: GetUserInfo :one
SELECT username, email
FROM users
WHERE id = @user_id;

-- name: GetUserHash :one
SELECT hash
FROM users
WHERE id = @user_id;

-- name: UpdateUsername :exec
UPDATE users
SET username = @new_username
WHERE id = @user_id;

-- name: UpdateEmail :exec
UPDATE users
SET email = @new_email
WHERE id = @user_id;

-- name: UpdateHash :exec
UPDATE users
SET hash = @new_hash
WHERE id = @user_id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @user_id;