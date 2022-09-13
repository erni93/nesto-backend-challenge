package router

import (
	"backend-api/repository"
	response "backend-api/router/response"
	"log"
	"net/http"
)

type ErasRouter struct {
	Repository *repository.EraRepository
}

func (router *ErasRouter) GetErasHandler(w http.ResponseWriter, r *http.Request) {
	eras, err := router.Repository.GetEras()
	if err != nil {
		log.Printf("Error retrieving eras: %s", err)
		response.WriteGeneralError(w)
	} else {
		response.WriteJsonObject(w, eras)
	}
}
