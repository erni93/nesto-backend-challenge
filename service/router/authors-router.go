package router

import (
	"backend-api/repository"
	response "backend-api/router/response"
	"log"
	"net/http"
)

type AuthorsRouter struct {
	Repository repository.AuthorRepository
}

// Returns a list of authors
func (router *AuthorsRouter) GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	authors, err := router.Repository.GetAuthors()
	if err != nil {
		log.Printf("Error retrieving authors: %s", err)
		response.WriteGeneralError(w)
	} else {
		response.WriteJsonObject(w, authors)
	}
}
