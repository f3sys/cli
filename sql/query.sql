-- name: CreateNode :one
INSERT INTO nodes (key, name, ip, type)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: CreateFood :one
INSERT INTO foods (name, price)
VALUES ($1, $2)
RETURNING *;