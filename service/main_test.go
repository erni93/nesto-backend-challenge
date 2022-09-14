package main

import (
	"backend-api/database"
	"backend-api/repository"
	"backend-api/router"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func readSampleFile(path string) (string, error) {
	outputSample, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, outputSample)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func TestGetAuthors(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	expected, err := readSampleFile("./samples/authors.json")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	authorsRouter := router.AuthorsRouter{
		Repository: &repository.AuthorRepository{DB: db},
	}

	req, err := http.NewRequest("GET", "/api/v1/authors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authorsRouter.GetAuthorsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}

}

func TestGetBooks(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	expected, err := readSampleFile("./samples/books.json")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	booksRouter := router.BooksRouter{
		Repository: &repository.BookRepository{DB: db},
	}

	req, err := http.NewRequest("GET", "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(booksRouter.GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}

}

func TestGetBooksWithFilters(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	expected, err := readSampleFile("./samples/books-with-filters.json")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	booksRouter := router.BooksRouter{
		Repository: &repository.BookRepository{DB: db},
	}

	req, err := http.NewRequest("GET", "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("authors", "4,1,21")
	q.Add("genres", "2")
	q.Add("min-pages", "500")
	q.Add("max-pages", "799")
	q.Add("limit", "10")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(booksRouter.GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}

}

func TestGetEras(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	expected, err := readSampleFile("./samples/eras.json")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	erasRouter := router.ErasRouter{
		Repository: &repository.EraRepository{DB: db},
	}

	req, err := http.NewRequest("GET", "/api/v1/eras", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(erasRouter.GetErasHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}

}

func TestGetGenres(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	expected, err := readSampleFile("./samples/genres.json")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	genresRouter := router.GenresRouter{
		Repository: &repository.GenreRepository{DB: db},
	}

	req, err := http.NewRequest("GET", "/api/v1/genres", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(genresRouter.GetGenresHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}

}

func TestGetSizes(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	expected, err := readSampleFile("./samples/sizes.json")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	sizesRouter := router.SizesRouter{
		Repository: &repository.SizeRepository{DB: db},
	}

	req, err := http.NewRequest("GET", "/api/v1/sizes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sizesRouter.GetSizesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}

}
