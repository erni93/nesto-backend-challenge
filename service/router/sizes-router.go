package router

import (
	"backend-api/repository"
	response "backend-api/router/response"
	"log"
	"net/http"
)

type SizesRouter struct {
	Repository repository.SizeRepository
}

func (router *SizesRouter) GetSizesHandler(w http.ResponseWriter, r *http.Request) {
	sizes, err := router.Repository.GetSizes()
	if err != nil {
		log.Printf("Error retrieving sizes: %s", err)
		response.WriteGeneralError(w)
	} else {
		response.WriteJsonObject(w, sizes)
	}
}
