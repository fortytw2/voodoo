package voodoo

// KeyValue is the primitive interface
// on top of which SQL is mapped.
//
// /columnspace/pkey = row
type KeyValue interface {
	Get(key []byte) ([]byte, error)
	Set(key []byte, val []byte) error
	// Walk is used to traverse parts of the keyspace
	Walk(prefix []byte) (Iterator, error)
}

// An Iterator is used to traverse the keyspace
type Iterator interface {
	// get the value at the iterator location
	Val() []byte
	// get the current key
	Key() []byte
	// move to the first key
	First() bool
	// move to the last key
	Last() bool
	// move to the next key
	Next() bool
	// move to the previous key
	Prev() bool
}
