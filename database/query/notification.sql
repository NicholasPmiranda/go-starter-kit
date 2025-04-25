-- name: FindManyNotifications :many
SELECT * FROM notifications;

-- name: FindNotificationById :one
SELECT * FROM notifications WHERE id = @id;

-- name: FindNotificationsByUserId :many
SELECT * FROM notifications WHERE user_id = @user_id;

-- name: FindUnreadNotificationsByUserId :many
SELECT * FROM notifications WHERE user_id = @user_id AND read = false;

-- name: FindNotificationsByNotifiable :many
SELECT * FROM notifications WHERE notifiable_type = @notifiable_type AND notifiable_id = @notifiable_id;

-- name: CreateNotification :one
INSERT INTO notifications (user_id, title, content, type, notifiable_type, notifiable_id)
VALUES (@user_id, @title, @content, @type, @notifiable_type, @notifiable_id) RETURNING *;

-- name: UpdateNotification :one
UPDATE notifications
SET title = @title, content = @content, type = @type, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: MarkNotificationAsRead :one
UPDATE notifications
SET read = true, read_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: MarkAllNotificationsAsRead :many
UPDATE notifications
SET read = true, read_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE user_id = @user_id AND read = false
RETURNING *;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE id = @id;
