package leveldb

import (
	"github.com/fortytw2/voodoo"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// A DB is the core storage layer
type DB struct {
	level *leveldb.DB
}

// NewDB constructs a new LevelDB backed KV store
func NewDB(filepath string) (*DB, error) {
	db, err := leveldb.OpenFile(filepath, nil)
	if err != nil {
		return nil, err
	}

	return &DB{
		level: db,
	}, nil
}

// NewMem makes an in-mem levelDB
func NewMem() (*DB, error) {
	db, err := leveldb.Open(storage.NewMemStorage(), nil)
	if err != nil {
		return nil, err
	}

	return &DB{
		level: db,
	}, nil
}

// Get returns the value at a certain key
func (db *DB) Get(key []byte) ([]byte, error) {
	return db.level.Get(key, nil)
}

// Set sets the given key to a value
func (db *DB) Set(key []byte, val []byte) error {
	return db.level.Put(key, val, nil)
}

// Walk is used to traverse parts of the keyspace
func (db *DB) Walk(prefix []byte) (voodoo.Iterator, error) {
	i := db.level.NewIterator(util.BytesPrefix(prefix), nil)
	return &Iterator{i: i}, nil
}
