package server

import (
	"encoding/json"
)

// IDriver opens IContext.
type IDriver interface {
	// OpenCtx opens an IContext with client capability, collation and dbname.
	OpenCtx(capability uint32, collation uint8, dbname string) (IContext, error)
}

// IContext is the interface to execute commant.
type IContext interface {
	// Status returns server status code.
	Status() uint16

	// LastInsertID returns last inserted ID.
	LastInsertID() uint64

	// AffectedRows returns affected rows of last executed command.
	AffectedRows() uint64

	// WarningCount returns warning count of last executed command.
	WarningCount() uint16

	// CurrentDB returns current DB.
	CurrentDB() string

	// Execute executes a SQL statement.
	Execute(sql string) (*ResultSet, error)

	// Prepare prepares a statement.
	Prepare(sql string) (statement IStatement, columns, params []*ColumnInfo, err error)

	// GetStatement get IStatement by statement ID.
	GetStatement(stmtID int) IStatement

	// FieldList returns columns of a table.
	FieldList(tableName string) (columns []*ColumnInfo, err error)

	// Close closes the IContext.
	Close() error
}

// IStatement is the interface to use a prepared statement.
type IStatement interface {
	// ID returns statement ID
	ID() int

	// Execute executes the statement.
	Execute(args ...interface{}) (*ResultSet, error)

	// AppendParam appends parameter to the statement.
	AppendParam(paramID int, data []byte) error

	// NumParams returns number of parameters.
	NumParams() int

	// BoundParams returns bound parameters.
	BoundParams() [][]byte

	// Reset remove all bound parameters.
	Reset()

	// Close closes the statement.
	Close() error
}

// ResultSet is the result set of an query.
type ResultSet struct {
	Columns []*ColumnInfo
	Rows    [][]interface{}
}

// String implements fmt.Stringer
func (res *ResultSet) String() string {
	b, _ := json.MarshalIndent(res, "", "\t")
	return string(b)
}

// AddRow appends a row to the result set.
func (res *ResultSet) AddRow(values ...interface{}) *ResultSet {
	res.Rows = append(res.Rows, values)
	return res
}
