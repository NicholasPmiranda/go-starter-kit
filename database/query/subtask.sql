-- name: FindManySubtasks :many
SELECT * FROM subtasks;

-- name: FindManySubtasksWithPagination :many
SELECT * FROM subtasks
WHERE id > 0
ORDER BY id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountSubtasks :one
SELECT COUNT(*) FROM subtasks;

-- name: FindSubtaskById :one
SELECT * FROM subtasks WHERE id = @id;

-- name: FindSubtasksByTaskId :many
SELECT * FROM subtasks WHERE task_id = @task_id;

-- name: FindSubtasksByAssignedTo :many
SELECT * FROM subtasks WHERE assigned_to = @assigned_to;

-- name: FindSubtasksByStatus :many
SELECT * FROM subtasks WHERE status = @status;

-- name: CreateSubtask :one
INSERT INTO subtasks (title, description, task_id, assigned_to, status, due_date)
VALUES (@title, @description, @task_id, @assigned_to, @status, @due_date) RETURNING *;

-- name: UpdateSubtask :one
UPDATE subtasks
SET title = @title, description = @description, task_id = @task_id, assigned_to = @assigned_to,
    status = @status, due_date = @due_date, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: CompleteSubtask :one
UPDATE subtasks
SET status = 'completed', completed_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteSubtask :exec
DELETE FROM subtasks
WHERE id = @id;
