package router

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Message: err})
}

func WriteGeneralError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Message: "Something went wrong. Please try again later"})
}
