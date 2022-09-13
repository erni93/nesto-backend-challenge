package database

import (
	"database/sql"
	"encoding/json"
)

// NullInt32 is an alias for sql.NullInt32 data type
type NullInt32 struct {
	sql.NullInt32
}

// MarshalJSON for NullInt32
func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

type IgnoreColumn struct{}

func (IgnoreColumn) Scan(value interface{}) error {
	return nil
}
