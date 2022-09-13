package model

type Book struct {
	Id            int64   `json:"id"`
	YearPublished int     `json:"yearPublished"`
	Rating        float32 `json:"rating"`
	Pages         int     `json:"pages"`
	Title         string  `json:"title"`
	Genre         *Genre  `json:"genre"`
	Author        *Author `json:"author"`
}
