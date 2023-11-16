-- name: CreateRFC :one
INSERT INTO rfcs ( author_id, message )
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: DeleteRFC :exec
DELETE FROM rfcs
WHERE id = @rfc_id;