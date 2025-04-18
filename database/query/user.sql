-- name: FindById :one
select * from users where id = $1;

-- name: FindByEmail :one
select * from users WHERE email = $1 limit 1;


-- name: FindMany :many
select * from users;

-- name: CreateUser :one
insert into users (name, email, password)
 values ($1, $2,$3) returning *;
