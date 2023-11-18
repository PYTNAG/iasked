-- name: CreateComment :one
INSERT INTO comments ( author_id, rfc_id )
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: GetRFCComments :many
SELECT * FROM comments
WHERE rfc_id = $1
ORDER BY created_at ASC
LIMIT sqlc.arg('count') OFFSET sqlc.arg('offset');

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = @comment_id;

-- name: CountRFCComments :one
SELECT COUNT(*) FROM comments
WHERE rfc_id = $1;