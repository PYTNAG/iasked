// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: rfc.sql

package db

import (
	"context"
	"database/sql"
)

const createRFC = `-- name: CreateRFC :one
INSERT INTO rfcs ( author_id, message )
VALUES (
    $1,
    $2
)
RETURNING id
`

type CreateRFCParams struct {
	AuthorID sql.NullInt32 `json:"author_id"`
	Message  string        `json:"message"`
}

func (q *Queries) CreateRFC(ctx context.Context, arg CreateRFCParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createRFC, arg.AuthorID, arg.Message)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteRFC = `-- name: DeleteRFC :exec
DELETE FROM rfcs
WHERE id = $1
`

func (q *Queries) DeleteRFC(ctx context.Context, rfcID int32) error {
	_, err := q.db.ExecContext(ctx, deleteRFC, rfcID)
	return err
}

const getLastRFCs = `-- name: GetLastRFCs :many
SELECT id, author_id, message, created_at, archived FROM rfcs
ORDER BY rfcs.created_at DESC
LIMIT $2 OFFSET $1
`

type GetLastRFCsParams struct {
	Offset int32 `json:"offset"`
	Count  int32 `json:"count"`
}

func (q *Queries) GetLastRFCs(ctx context.Context, arg GetLastRFCsParams) ([]Rfc, error) {
	rows, err := q.db.QueryContext(ctx, getLastRFCs, arg.Offset, arg.Count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Rfc{}
	for rows.Next() {
		var i Rfc
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Message,
			&i.CreatedAt,
			&i.Archived,
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
