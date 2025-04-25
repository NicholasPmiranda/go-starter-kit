-- name: FindManyClients :many
SELECT * FROM clients;

-- name: FindManyClientsWithPagination :many
SELECT * FROM clients
WHERE id > 0
ORDER BY id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountClients :one
SELECT COUNT(*) FROM clients;

-- name: FindClientById :one
SELECT * FROM clients WHERE id = @id;

-- name: CreateClient :one
INSERT INTO clients (name, email, phone, address)
VALUES (@name, @email, @phone, @address) RETURNING *;

-- name: UpdateClient :one
UPDATE clients
SET name = @name, email = @email, phone = @phone, address = @address, updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;

-- name: DeleteClient :exec
DELETE FROM clients
WHERE id = @id;
