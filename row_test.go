package voodoo

import (
	"testing"

	"github.com/fortytw2/voodoo/types"
)

func TestRowMarshaling(t *testing.T) {
	exRow := Row{
		types.Int64(1),
		types.Float64(56.6),
		types.Text("I am a Ptoato AMA"),
	}

	buf, err := exRow.MarshalBinary()
	if err != nil {
		t.Error("could not marshal binary", err)
	}

	var r Row
	err = r.UnmarshalBinary(buf)
	if err != nil {
		t.Error("could not unmarshal binary", err)
	}

	if len(r) != 3 {
		t.Error("did not unmarshal binary correctly")
	}
}
