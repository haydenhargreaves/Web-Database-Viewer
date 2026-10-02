package model

import "database/sql"

// Column data structure
type Column struct {
	Name       string
	Type       string
	MaxLength  sql.NullInt64
	Nullable   string
	PrimaryKey bool
	ForeignKey ForeignKey
	Unique     bool
}

type ForeignKey struct {
	Column        string
	ForeignTable  string
	ForeignColumn string
}

// Routine is a stored procedure or function that can be displayed in the tree.
type Routine struct {
	ID        string
	Schema    string
	Name      string
	Signature string
	Kind      string
}
