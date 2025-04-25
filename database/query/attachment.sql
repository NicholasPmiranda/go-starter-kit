-- name: FindManyAttachments :many
SELECT * FROM attachments;

-- name: FindAttachmentById :one
SELECT * FROM attachments WHERE id = @id;

-- name: FindAttachmentsByUserId :many
SELECT * FROM attachments WHERE user_id = @user_id;

-- name: FindAttachmentsByAttachable :many
SELECT * FROM attachments WHERE attachable_type = @attachable_type AND attachable_id = @attachable_id;

-- name: CreateAttachment :one
INSERT INTO attachments (filename, filepath, filesize, filetype, user_id, attachable_type, attachable_id)
VALUES (@filename, @filepath, @filesize, @filetype, @user_id, @attachable_type, @attachable_id) RETURNING *;

-- name: UpdateAttachment :one
UPDATE attachments
SET filename = @filename, filepath = @filepath, filesize = @filesize, filetype = @filetype, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteAttachment :exec
DELETE FROM attachments
WHERE id = @id;
