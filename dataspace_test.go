package voodoo

import (
	"strings"
	"testing"
)

func TestLoadDataspace(t *testing.T) {
	ds, err := LoadDataspace(strings.NewReader(exDS))
	if err != nil {
		t.Error(err)
	}

	if len(ds.Tables) != 1 {
		t.Error("did not load tables")
	}
}

func TestDataspaceValidation(t *testing.T) {
	_, err := LoadDataspace(strings.NewReader(invalidDS))
	if err == nil {
		t.Error("no error raised during loading of invalid dataspace")
	}
	if err != ErrDuplicateNames {
		t.Error("duplicate names allowed in dataspace")
	}
}

var exDS = `
{
    "name": "starwarsdb",
    "tables": [
        {
            "name": "planets",
            "columns": [
                {
                    "name": "id",
                    "type": "int64",
                    "primary_key": true
                },
                {
                    "name": "name",
                    "type": "text"
                },
                {
                    "name": "description",
                    "type": "text"
                }
            ]
        }
    ]
}
`

var invalidDS = `
{
    "name": "starwarsdb",
    "tables": [
        {
            "name": "planets"
        },
        {
            "name": "planets",
            "columns": [
                {
                    "name": "id",
                    "type": "int64",
                    "primary_key": true
                },
                {
                    "name": "name",
                    "type": "text"
                },
                {
                    "name": "description",
                    "type": "text"
                }
            ]
        }
    ]
}
`
