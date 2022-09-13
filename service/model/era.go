package model

import "backend-api/database"

type Era struct {
	Id      int64               `json:"id"`
	Title   string              `json:"title"`
	MinYear *database.NullInt32 `json:"minYear,omitempty"`
	MaxYear *database.NullInt32 `json:"maxYear,omitempty"`
}
