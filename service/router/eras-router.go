package router

import (
	"backend-api/model"
	response "backend-api/router/response"
	"net/http"
)

type ErasRouter struct {
}

func (router *ErasRouter) GetErasHandler(w http.ResponseWriter, r *http.Request) {
	eras := []model.Era{{
		Id:      1,
		Title:   "eras-test",
		MinYear: 1950,
		MaxYear: 2000,
	}}

	response.WriteJsonObject(w, eras)
}
