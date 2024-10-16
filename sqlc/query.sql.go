// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBattery = `-- name: CreateBattery :exec
INSERT INTO batteries (node_id)
VALUES ($1)
`

func (q *Queries) CreateBattery(ctx context.Context, nodeID pgtype.Int8) error {
	_, err := q.db.Exec(ctx, createBattery, nodeID)
	return err
}

const createFood = `-- name: CreateFood :one
INSERT INTO foods (name, price, quantity)
VALUES ($1, $2, $3)
RETURNING id, name, price, quantity, created_at, updated_at
`

type CreateFoodParams struct {
	Name     string
	Price    int32
	Quantity int32
}

func (q *Queries) CreateFood(ctx context.Context, arg CreateFoodParams) (Food, error) {
	row := q.db.QueryRow(ctx, createFood, arg.Name, arg.Price, arg.Quantity)
	var i Food
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		&i.Quantity,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createKey = `-- name: CreateKey :one
UPDATE nodes SET key = $1, updated_at = now() WHERE id = $2 RETURNING key
`

type CreateKeyParams struct {
	Key pgtype.Text
	ID  int64
}

func (q *Queries) CreateKey(ctx context.Context, arg CreateKeyParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, createKey, arg.Key, arg.ID)
	var key pgtype.Text
	err := row.Scan(&key)
	return key, err
}

const createNode = `-- name: CreateNode :one
INSERT INTO nodes (key, name, type)
VALUES ($1, $2, $3)
RETURNING id, key, otp, name, ip, type, created_at, updated_at
`

type CreateNodeParams struct {
	Key  pgtype.Text
	Name string
	Type NodeType
}

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) (Node, error) {
	row := q.db.QueryRow(ctx, createNode, arg.Key, arg.Name, arg.Type)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Otp,
		&i.Name,
		&i.Ip,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createOTP = `-- name: CreateOTP :one
UPDATE nodes SET otp = $1, updated_at = now() WHERE id = $2 RETURNING otp
`

type CreateOTPParams struct {
	Otp pgtype.Text
	ID  int64
}

func (q *Queries) CreateOTP(ctx context.Context, arg CreateOTPParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, createOTP, arg.Otp, arg.ID)
	var otp pgtype.Text
	err := row.Scan(&otp)
	return otp, err
}

const createOTPandDeleteKey = `-- name: CreateOTPandDeleteKey :one
UPDATE nodes SET otp = $1, key = NULL, updated_at = now() WHERE id = $2 RETURNING otp
`

type CreateOTPandDeleteKeyParams struct {
	Otp pgtype.Text
	ID  int64
}

func (q *Queries) CreateOTPandDeleteKey(ctx context.Context, arg CreateOTPandDeleteKeyParams) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, createOTPandDeleteKey, arg.Otp, arg.ID)
	var otp pgtype.Text
	err := row.Scan(&otp)
	return otp, err
}

const createStudent = `-- name: CreateStudent :one
INSERT INTO students (visitor_id, grade, class)
VALUES ($1, $2, $3)
RETURNING id
`

type CreateStudentParams struct {
	VisitorID int64
	Grade     int32
	Class     int32
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (int64, error) {
	row := q.db.QueryRow(ctx, createStudent, arg.VisitorID, arg.Grade, arg.Class)
	var id int64
	err := row.Scan(&id)
	return id, err
}

type CreateStudentsParams struct {
	VisitorID int64
	Grade     int32
	Class     int32
}

const createVisitor = `-- name: CreateVisitor :one
INSERT INTO visitors (random)
VALUES ($1)
RETURNING id
`

func (q *Queries) CreateVisitor(ctx context.Context, random int32) (int64, error) {
	row := q.db.QueryRow(ctx, createVisitor, random)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getLastVisitorId = `-- name: GetLastVisitorId :one
SELECT (id) FROM visitors ORDER BY id DESC LIMIT 1
`

func (q *Queries) GetLastVisitorId(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, getLastVisitorId)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getNodes = `-- name: GetNodes :many
SELECT id, key, otp, name, ip, type, created_at, updated_at FROM nodes
`

func (q *Queries) GetNodes(ctx context.Context) ([]Node, error) {
	rows, err := q.db.Query(ctx, getNodes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.ID,
			&i.Key,
			&i.Otp,
			&i.Name,
			&i.Ip,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
