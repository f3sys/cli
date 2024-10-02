// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
	"net/netip"

	"github.com/jackc/pgx/v5/pgtype"
)

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

const createNode = `-- name: CreateNode :one
INSERT INTO nodes (key, name, ip, type)
VALUES ($1, $2, $3, $4)
RETURNING id, key, name, ip, type, created_at, updated_at
`

type CreateNodeParams struct {
	Key  pgtype.Text
	Name string
	Ip   *netip.Addr
	Type NodeType
}

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) (Node, error) {
	row := q.db.QueryRow(ctx, createNode,
		arg.Key,
		arg.Name,
		arg.Ip,
		arg.Type,
	)
	var i Node
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Name,
		&i.Ip,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
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
