package main

import (
	"backend-api/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	authorsRouter := router.AuthorsRouter{}
	booksRouter := router.BooksRouter{}
	erasRouter := router.ErasRouter{}
	genresRouter := router.GenresRouter{}
	sizesRouter := router.SizesRouter{}

	port := ":5001"
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.HandleFunc("/api/v1/authors", authorsRouter.GetAuthorsHandler).Methods("GET")
	router.HandleFunc("/api/v1/books", booksRouter.GetBooksHandler).Methods("GET")
	router.HandleFunc("/api/v1/eras", erasRouter.GetErasHandler).Methods("GET")
	router.HandleFunc("/api/v1/genres", genresRouter.GetGenresHandler).Methods("GET")
	router.HandleFunc("/api/v1/sizes", sizesRouter.GetSizesHandler).Methods("GET")

	http.Handle("/", router)
	log.Printf("Application listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
