// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ( email, username, hash )
VALUES (
    $1,
    COALESCE($2::text, 'i'),
    $3
)
RETURNING id, username, email
`

type CreateUserParams struct {
	Email    string         `json:"email"`
	Username sql.NullString `json:"username"`
	Hash     []byte         `json:"hash"`
}

type CreateUserRow struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.Username, arg.Hash)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Username, &i.Email)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getUserHash = `-- name: GetUserHash :one
SELECT hash
FROM users
WHERE id = $1
`

func (q *Queries) GetUserHash(ctx context.Context, userID int32) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, getUserHash, userID)
	var hash []byte
	err := row.Scan(&hash)
	return hash, err
}

const getUserInfo = `-- name: GetUserInfo :one
SELECT username, email
FROM users
WHERE id = $1
`

type GetUserInfoRow struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (q *Queries) GetUserInfo(ctx context.Context, userID int32) (GetUserInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getUserInfo, userID)
	var i GetUserInfoRow
	err := row.Scan(&i.Username, &i.Email)
	return i, err
}

const updateEmail = `-- name: UpdateEmail :exec
UPDATE users
SET email = $1
WHERE id = $2
`

type UpdateEmailParams struct {
	NewEmail string `json:"new_email"`
	UserID   int32  `json:"user_id"`
}

func (q *Queries) UpdateEmail(ctx context.Context, arg UpdateEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateEmail, arg.NewEmail, arg.UserID)
	return err
}

const updateHash = `-- name: UpdateHash :exec
UPDATE users
SET hash = $1
WHERE id = $2
`

type UpdateHashParams struct {
	NewHash []byte `json:"new_hash"`
	UserID  int32  `json:"user_id"`
}

func (q *Queries) UpdateHash(ctx context.Context, arg UpdateHashParams) error {
	_, err := q.db.ExecContext(ctx, updateHash, arg.NewHash, arg.UserID)
	return err
}

const updateUsername = `-- name: UpdateUsername :exec
UPDATE users
SET username = $1
WHERE id = $2
`

type UpdateUsernameParams struct {
	NewUsername string `json:"new_username"`
	UserID      int32  `json:"user_id"`
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) error {
	_, err := q.db.ExecContext(ctx, updateUsername, arg.NewUsername, arg.UserID)
	return err
}
