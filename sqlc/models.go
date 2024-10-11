// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql/driver"
	"fmt"
	"net/netip"

	"github.com/jackc/pgx/v5/pgtype"
)

type EntryLogsType string

const (
	EntryLogsTypeENTERED EntryLogsType = "ENTERED"
	EntryLogsTypeLEFT    EntryLogsType = "LEFT"
)

func (e *EntryLogsType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EntryLogsType(s)
	case string:
		*e = EntryLogsType(s)
	default:
		return fmt.Errorf("unsupported scan type for EntryLogsType: %T", src)
	}
	return nil
}

type NullEntryLogsType struct {
	EntryLogsType EntryLogsType
	Valid         bool // Valid is true if EntryLogsType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullEntryLogsType) Scan(value interface{}) error {
	if value == nil {
		ns.EntryLogsType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.EntryLogsType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullEntryLogsType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.EntryLogsType), nil
}

type NodeType string

const (
	NodeTypeENTRY      NodeType = "ENTRY"
	NodeTypeFOODSTALL  NodeType = "FOODSTALL"
	NodeTypeEXHIBITION NodeType = "EXHIBITION"
)

func (e *NodeType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NodeType(s)
	case string:
		*e = NodeType(s)
	default:
		return fmt.Errorf("unsupported scan type for NodeType: %T", src)
	}
	return nil
}

type NullNodeType struct {
	NodeType NodeType
	Valid    bool // Valid is true if NodeType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNodeType) Scan(value interface{}) error {
	if value == nil {
		ns.NodeType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NodeType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNodeType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.NodeType), nil
}

type Battery struct {
	ID              int64
	NodeID          int64
	Level           pgtype.Int4
	ChargingTime    pgtype.Int4
	DischargingTime pgtype.Int4
	Charging        pgtype.Bool
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
}

type EntryLog struct {
	ID        int64
	NodeID    int64
	VisitorID int64
	Type      EntryLogsType
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type ExhibitionLog struct {
	ID        int64
	NodeID    int64
	VisitorID int64
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Food struct {
	ID        int64
	Name      string
	Price     int32
	Quantity  int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type FoodStallLog struct {
	ID         int64
	NodeFoodID int64
	VisitorID  int64
	Quantity   int32
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

type Model struct {
	ID        int64
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Node struct {
	ID        int64
	Key       pgtype.Text
	Name      string
	Ip        *netip.Addr
	Type      NodeType
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type NodeFood struct {
	ID        int64
	NodeID    int64
	FoodID    int64
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Student struct {
	ID        int64
	VisitorID int64
	Grade     int32
	Class     int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Visitor struct {
	ID        int64
	ModelID   pgtype.Int8
	Ip        *netip.Addr
	Random    int32
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
