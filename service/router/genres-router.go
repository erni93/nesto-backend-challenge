package router

import (
	"backend-api/model"
	response "backend-api/router/response"
	"net/http"
)

type GenresRouter struct {
}

func (router *GenresRouter) GetGenresHandler(w http.ResponseWriter, r *http.Request) {
	genres := []model.Genre{{
		Id:    1,
		Title: "genres-test",
	}}

	response.WriteJsonObject(w, genres)
}
