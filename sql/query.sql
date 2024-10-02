-- name: CreateNode :one
INSERT INTO nodes (key, name, ip, type)
VALUES ($1, $2, $3, $4)
RETURNING *;
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
-- name: CreateVisitors :copyfrom
INSERT INTO visitors (random)
VALUES ($1);
-- name: CreateStudents :copyfrom
INSERT INTO students (visitor_id, grade, class)
VALUES ($1, $2, $3);