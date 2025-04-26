-- name: FindManyProjects :many
SELECT *
FROM projects;

-- name: FindManyProjectsWithPagination :many
SELECT *
FROM projects
WHERE id > 0
ORDER BY id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountProjects :one
SELECT COUNT(*)
FROM projects;

-- name: FindProjectById :one
SELECT *
FROM projects
WHERE id = @id;

-- name: FindProjectsByClientId :many
SELECT *
FROM projects
WHERE client_id = @client_id;

-- name: FindProjectsByUser :many
select *
from user_project
where user_id = @user_id;


-- name: FinUsersByProject :many
select *
from user_project
where project_id = @project_id;

-- name: CreateUserProject :exec
insert into user_project (user_id, project_id)
values (@user_id, @project_id);


-- name: DeleteUserProject :exec
delete from user_project
where user_id = @user_id and project_id = @project_id;


-- name: CreateProject :one
INSERT INTO projects (name, description, client_id, status, start_date, end_date)
VALUES (@name, @description, @client_id, @status, @start_date, @end_date)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET name        = @name,
    description = @description,
    client_id   = @client_id,
    status      = @status,
    start_date  = @start_date,
    end_date    = @end_date,
    updated_at  = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = @id;
