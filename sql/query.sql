-- name: CreateNode :one
INSERT INTO nodes (key, name, type)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateBattery :exec
INSERT INTO batteries (node_id)
VALUES ($1);
-- name: CreateFood :one
INSERT INTO foods (name, price, quantity)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateVisitor :one
INSERT INTO visitors (random)
VALUES ($1)
RETURNING id;
-- name: CreateStudent :one
INSERT INTO students (visitor_id, grade, class)
VALUES ($1, $2, $3)
RETURNING id;
-- name: GetLastVisitorId :one
SELECT (id) FROM visitors ORDER BY id DESC LIMIT 1;
-- name: CreateVisitors :copyfrom
INSERT INTO visitors (random)
VALUES ($1);
-- name: CreateStudents :copyfrom
INSERT INTO students (visitor_id, grade, class)
VALUES ($1, $2, $3);
-- name: GetNodes :many
SELECT * FROM nodes;
-- name: CreateOTP :one
UPDATE nodes SET otp = $1, updated_at = now() WHERE id = $2 RETURNING otp;
-- name: CreateOTPandDeleteKey :one
UPDATE nodes SET otp = $1, key = NULL, updated_at = now() WHERE id = $2 RETURNING otp;
-- name: CreateKey :one
UPDATE nodes SET key = $1, updated_at = now() WHERE id = $2 RETURNING key;