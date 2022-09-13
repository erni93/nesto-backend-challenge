package router

import (
	"backend-api/repository"
	response "backend-api/router/response"
	"log"
	"net/http"
)

type GenresRouter struct {
	Repository *repository.GenreRepository
}

func (router *GenresRouter) GetGenresHandler(w http.ResponseWriter, r *http.Request) {
	genres, err := router.Repository.GetGenres()
	if err != nil {
		log.Printf("Error retrieving genres: %s", err)
		response.WriteGeneralError(w)
	} else {
		response.WriteJsonObject(w, genres)
	}
}
