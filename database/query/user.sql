-- name: FindById :one
SELECT * FROM users WHERE id = @id;

-- name: FindByEmail :one
SELECT * FROM users WHERE email = @email LIMIT 1;

-- name: FindMany :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES (@name, @email, @password) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = @name, email = @email, password = @password
WHERE id = @id
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @id;
