-- name: FindManyComments :many
SELECT * FROM comments;

-- name: FindManyCommentsWithPagination :many
SELECT * FROM comments
WHERE id > 0
ORDER BY id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountComments :one
SELECT COUNT(*) FROM comments;

-- name: FindCommentById :one
SELECT * FROM comments WHERE id = @id;

-- name: FindCommentsByUserId :many
SELECT * FROM comments WHERE user_id = @user_id;

-- name: FindCommentsByCommentable :many
SELECT * FROM comments WHERE commentable_type = @commentable_type AND commentable_id = @commentable_id;

-- name: CreateComment :one
INSERT INTO comments (content, user_id, commentable_type, commentable_id)
VALUES (@content, @user_id, @commentable_type, @commentable_id) RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET content = @content, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = @id;
