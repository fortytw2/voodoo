package voodoo

// A ResultSet is returned from a query
type ResultSet struct {
	// types of each column, in order
	ColumnTypes []ColumnType
	// names of each column in the resultset
	ColumnNames []string
	// all rows returned
	Rows []Row
	// an error, if any
	Err error
}
