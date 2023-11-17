-- name: CreateRFC :one
INSERT INTO rfcs ( author_id, message )
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: GetLastRFCs :many
SELECT * FROM rfcs
ORDER BY rfcs.created_at DESC
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');


-- name: DeleteRFC :exec
DELETE FROM rfcs
WHERE id = @rfc_id;