package model

import "backend-api/database"

type Size struct {
	Id       int64               `json:"id"`
	Title    string              `json:"title"`
	MinPages *database.NullInt32 `json:"minPages,omitempty"`
	MaxPages *database.NullInt32 `json:"maxPages,omitempty"`
}
