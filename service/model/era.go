package model

type Era struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	MinYear int    `json:"minYear,omitempty"`
	MaxYear int    `json:"maxYear,omitempty"`
}
