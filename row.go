package voodoo

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/fortytw2/voodoo/types"
)

type Row []interface{}

func (r *Row) Validate() bool {
	for _, v := range *r {
		switch v.(type) {
		case types.Float64, types.Int64, types.Text:
			continue
		default:
			return false
		}
	}
	return true
}

// MarshalBinary encodes the row
func (r *Row) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	for i, v := range *r {
		if i != 0 && i != len(*r) {
			// this totally breaks but fuck it
			b.WriteString("-?X?-")
		}
		f64, ok := v.(types.Float64)
		if ok {
			b.WriteString(fmt.Sprintf("%f", f64))
			continue
		}

		i64, ok := v.(types.Int64)
		if ok {
			b.WriteString(fmt.Sprintf("%d", i64))
			continue
		}

		text, ok := v.(types.Text)
		if ok {
			b.WriteString(fmt.Sprintf("%s", text))
		}
	}
	return b.Bytes(), nil
}

// UnmarshalBinary decodes the given row from a byte array
func (r *Row) UnmarshalBinary(buf []byte) error {
	o := make(Row, 0)

	for _, elem := range strings.Split(string(buf), "-?X?-") {
		i, err := strconv.Atoi(elem)
		if err == nil {
			o = append(o, types.Int64(i))
			continue
		}

		f64, err := strconv.ParseFloat(elem, 64)
		if err == nil {
			o = append(o, types.Float64(f64))
			continue
		}

		o = append(o, types.Text(elem))
	}

	*r = o
	return nil
}
