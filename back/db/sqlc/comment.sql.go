// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: comment.sql

package db

import (
	"context"
	"database/sql"
)

const countRFCComments = `-- name: CountRFCComments :one
SELECT COUNT(*) FROM comments
WHERE rfc_id = $1
`

func (q *Queries) CountRFCComments(ctx context.Context, rfcID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, countRFCComments, rfcID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createComment = `-- name: CreateComment :one
INSERT INTO comments ( author_id, rfc_id )
VALUES (
    $1,
    $2
)
RETURNING id
`

type CreateCommentParams struct {
	AuthorID sql.NullInt32 `json:"author_id"`
	RfcID    int32         `json:"rfc_id"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.AuthorID, arg.RfcID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, commentID int32) error {
	_, err := q.db.ExecContext(ctx, deleteComment, commentID)
	return err
}

const getRFCComments = `-- name: GetRFCComments :many
SELECT id, author_id, rfc_id, created_at FROM comments
WHERE rfc_id = $1
ORDER BY created_at ASC
LIMIT $3 OFFSET $2
`

type GetRFCCommentsParams struct {
	RfcID  int32 `json:"rfc_id"`
	Offset int32 `json:"offset"`
	Count  int32 `json:"count"`
}

func (q *Queries) GetRFCComments(ctx context.Context, arg GetRFCCommentsParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getRFCComments, arg.RfcID, arg.Offset, arg.Count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.RfcID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
