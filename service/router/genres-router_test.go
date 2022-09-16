package router

import (
	"backend-api/model"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockSuccessfulGenreRepository struct {
}

func (r *MockSuccessfulGenreRepository) GetGenres() ([]model.Genre, error) {
	return []model.Genre{{Id: 1, Title: "title1"}}, nil
}

type MockErrorGenreRepository struct {
}

func (r *MockErrorGenreRepository) GetGenres() ([]model.Genre, error) {
	return nil, errors.New("error")
}

func TestSuccessfulGetGenresHandler(t *testing.T) {
	GenresRouter := GenresRouter{
		Repository: &MockSuccessfulGenreRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/genres", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenresRouter.GetGenresHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"title1"}]`

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func TestErrorGetGenresHandler(t *testing.T) {
	GenresRouter := GenresRouter{
		Repository: &MockErrorGenreRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/genres", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenresRouter.GetGenresHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
