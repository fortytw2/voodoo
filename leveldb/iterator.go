package leveldb

import "github.com/syndtr/goleveldb/leveldb/iterator"

// An Iterator is used to walk the keyspace
type Iterator struct {
	i iterator.Iterator
}

// get the value at the iterator location
func (i *Iterator) Val() []byte {
	return i.i.Value()
}

// get the current key
func (i *Iterator) Key() []byte {
	return i.i.Key()
}

// move to the first key
func (i *Iterator) First() bool {
	return i.i.First()
}

// move to the last key
func (i *Iterator) Last() bool {
	return i.i.Last()
}

// move to the next key
func (i *Iterator) Next() bool {
	return i.i.Next()
}

// move to the previous key
func (i *Iterator) Prev() bool {
	return i.i.Prev()
}
