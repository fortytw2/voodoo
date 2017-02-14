package voodoo

import (
	"encoding/json"
	"errors"
	"io"
)

// Errors returned during dataspace parsing/validation
var (
	ErrDuplicateNames = errors.New("voodoo/dataspace: conflicting names within dataspace")
)

// Dataspace is the overall db metadata
type Dataspace struct {
	Name    string  `json:"name"`
	Tables  []Table `json:"tables"`
	Indexes []Index `json:"indexes"`
}

// LoadDataspace parses and validates a given dataspace
// from a json file
func LoadDataspace(r io.Reader) (*Dataspace, error) {
	var ds Dataspace
	err := json.NewDecoder(r).Decode(&ds)
	if err != nil {
		return nil, err
	}

	// ensure no names are duplicated
	allNames := make(map[string]bool)
	for _, t := range ds.Tables {
		if !allNames[t.Name] {
			allNames[t.Name] = true
		} else {
			return nil, ErrDuplicateNames
		}
	}

	return &ds, nil
}
