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
from project_user
where user_id = @user_id;

-- name: FinUsersByProject :many
select *
from project_user
where project_id = @project_id;

-- name: CreateUserProject :exec
insert into project_user (user_id, project_id)
values (@user_id, @project_id);

-- name: DeleteUserProject :exec
delete from project_user
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

-- name: FindProjectWithUsers :one
SELECT
    p.id,
    p.name,
    p.description,
    p.client_id,
    p.status,
    p.start_date,
    p.end_date,
    p.created_at,
    p.updated_at,
    COALESCE(
        json_agg(
            json_build_object(
                'id', u.id,
                'name', u.name,
                'email', u.email
            )
        ) FILTER (WHERE u.id IS NOT NULL),
        '[]'::json
    ) as users
FROM
    projects p
LEFT JOIN
    project_user up ON p.id = up.project_id
LEFT JOIN
    users u ON up.user_id = u.id
WHERE
    p.id = @project_id
GROUP BY
    p.id;

-- name: FindManyProjectsWithUsers :many
SELECT
    p.id,
    p.name,
    p.description,
    p.client_id,
    p.status,
    p.start_date,
    p.end_date,
    p.created_at,
    p.updated_at,
    COALESCE(
        json_agg(
            json_build_object(
                'id', u.id,
                'name', u.name,
                'email', u.email
            )
        ) FILTER (WHERE u.id IS NOT NULL),
        '[]'::json
    ) as users
FROM
    projects p
LEFT JOIN
    project_user up ON p.id = up.project_id
LEFT JOIN
    users u ON up.user_id = u.id
GROUP BY
    p.id
ORDER BY
    p.id;
