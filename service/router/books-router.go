package router

import (
	"backend-api/model"
	response "backend-api/router/response"
	"net/http"
)

type BooksRouter struct {
}

func (router *BooksRouter) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := []model.Book{{
		Id:            1,
		YearPublished: 1950,
		Rating:        5,
		Pages:         200,
		Title:         "books-title",
		Genre: &model.Genre{
			Id:    1,
			Title: "genres-test",
		},
		Author: &model.Author{
			Id:        1,
			FirstName: "author-first-name",
			LastName:  "author-last-name",
		},
	}}

	response.WriteJsonObject(w, books)
}
