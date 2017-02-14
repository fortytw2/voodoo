package voodoo

// A Table is a table in the dataspace
type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
}

// A ColumnType is the datatype of a column
type ColumnType string

var (
	Float64   ColumnType = "float64"
	Int64     ColumnType = "int64"
	Timezone  ColumnType = "timezone"
	Timestamp ColumnType = "timestamp"
	Date      ColumnType = "date"
	Text      ColumnType = "text"
	Bool      ColumnType = "bool"
)

// A Column is a column of a table
type Column struct {
	Name       string     `json:"name"`
	Type       ColumnType `json:"type"`
	PrimaryKey bool       `json:"primary_key,omitempty"`
}
