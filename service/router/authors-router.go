package router

import (
	"backend-api/model"
	response "backend-api/router/response"
	"net/http"
)

type AuthorsRouter struct {
}

func (router *AuthorsRouter) GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	author := []model.Author{{
		Id:        1,
		FirstName: "author-first-name",
		LastName:  "author-last-name",
	}}

	response.WriteJsonObject(w, author)
}
