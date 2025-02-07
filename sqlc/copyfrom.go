// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: copyfrom.go

package sqlc

import (
	"context"
)

// iteratorForCreateStudents implements pgx.CopyFromSource.
type iteratorForCreateStudents struct {
	rows                 []CreateStudentsParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateStudents) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateStudents) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].VisitorID,
		r.rows[0].Grade,
		r.rows[0].Class,
	}, nil
}

func (r iteratorForCreateStudents) Err() error {
	return nil
}

func (q *Queries) CreateStudents(ctx context.Context, arg []CreateStudentsParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"students"}, []string{"visitor_id", "grade", "class"}, &iteratorForCreateStudents{rows: arg})
}

// iteratorForCreateVisitors implements pgx.CopyFromSource.
type iteratorForCreateVisitors struct {
	rows                 []int32
	skippedFirstNextCall bool
}

func (r *iteratorForCreateVisitors) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateVisitors) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0],
	}, nil
}

func (r iteratorForCreateVisitors) Err() error {
	return nil
}

func (q *Queries) CreateVisitors(ctx context.Context, random []int32) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"visitors"}, []string{"random"}, &iteratorForCreateVisitors{rows: random})
}
