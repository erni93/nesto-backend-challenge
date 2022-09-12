package router

import (
	"encoding/json"
	"net/http"
)

func WriteJsonObject(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
