package router

import (
	"backend-api/repository"
	response "backend-api/router/response"
	"log"
	"net/http"
)

type BooksRouter struct {
	Repository *repository.BookRepository
}

func (router *BooksRouter) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := router.Repository.GetBooks()
	if err != nil {
		log.Printf("Error retrieving books: %s", err)
		response.WriteGeneralError(w)
	} else {
		response.WriteJsonObject(w, books)
	}
}
