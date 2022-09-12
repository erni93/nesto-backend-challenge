package model

type Size struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	MinPages int    `json:"minPages,omitempty"`
	MaxPages int    `json:"maxPages,omitempty"`
}
