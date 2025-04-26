-- name: FindManyTasks :many
SELECT * FROM tasks;

-- name: FindManyTasksWithPagination :many
SELECT * FROM tasks
WHERE id > 0
ORDER BY id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountTasks :one
SELECT COUNT(*) FROM tasks;

-- name: FindTaskById :one
SELECT * FROM tasks WHERE id = @id;

-- name: FindTasksByProjectId :many
SELECT * FROM tasks WHERE project_id = @project_id;

-- name: FindTasksByAssignedTo :many
SELECT * FROM tasks WHERE assigned_to = @assigned_to;

-- name: FindTasksByStatus :many
SELECT * FROM tasks WHERE status = @status;

-- name: FindTasksByPriority :many
SELECT * FROM tasks WHERE priority = @priority;

-- name: CreateTask :one
INSERT INTO tasks (title, description, project_id, assigned_to, status, priority, due_date)
VALUES (@title, @description, @project_id, @assigned_to, @status, @priority, @due_date) RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = @title, description = @description, project_id = @project_id, assigned_to = @assigned_to,
    status = @status, priority = @priority, due_date = @due_date, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: CompleteTask :one
UPDATE tasks
SET status = 'completed', completed_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = @id;

-- name: FindTasksWithUsersPaginated :many
SELECT t.* FROM tasks t
WHERE t.id > 0
ORDER BY t.id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: FindUsersByTaskIds :many
SELECT u.* FROM users u
JOIN tasks t ON t.assigned_to = u.id
WHERE t.id IN (sqlc.slice('task_ids'))
ORDER BY u.id;
