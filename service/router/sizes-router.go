package router

import (
	"backend-api/model"
	response "backend-api/router/response"
	"net/http"
)

type SizesRouter struct {
}

func (router *SizesRouter) GetSizesHandler(w http.ResponseWriter, r *http.Request) {
	sizes := []model.Size{{
		Id:       1,
		Title:    "sizes-test",
		MinPages: 1,
		MaxPages: 100,
	}}

	response.WriteJsonObject(w, sizes)
}
