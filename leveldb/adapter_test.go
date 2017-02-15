package leveldb

import (
	"testing"

	"github.com/fortytw2/voodoo"
)

func TestLevelDB(t *testing.T) {
	voodoo.TestKeyValue(t, func() (voodoo.KeyValue, error) {
		return NewMem()
	})
}
