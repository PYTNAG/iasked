-- name: CreateComment :one
INSERT INTO comments ( author_id, rfc_id )
VALUES (
    $1,
    $2
)
RETURNING id;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = @comment_id;