-- name: FindById :one
SELECT *
FROM users
WHERE id = @id;

-- name: FindManyUserIds :many
SELECT id, name, email, password
FROM users
WHERE id = ANY(@ids::bigint[]);

-- name: FindByEmail :one
SELECT *
FROM users
WHERE email = @email
LIMIT 1;

-- name: FindMany :many
SELECT *
FROM users;

-- name: FindManyWithPagination :many
SELECT *
FROM users
WHERE id > 0
ORDER BY id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountUsers :one
SELECT COUNT(*)
FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES (@name, @email, @password)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name     = @name,
    email    = @email,
    password = @password
WHERE id = @id
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = @id;
