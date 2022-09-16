package router

import (
	"backend-api/model"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockSuccessfulAuthorRepository struct {
}

func (r *MockSuccessfulAuthorRepository) GetAuthors() ([]model.Author, error) {
	return []model.Author{{Id: 1, FirstName: "name1", LastName: "lastName1"}}, nil
}

type MockErrorAuthorRepository struct {
}

func (r *MockErrorAuthorRepository) GetAuthors() ([]model.Author, error) {
	return nil, errors.New("error")
}

func TestSuccessfulGetAuthorsHandler(t *testing.T) {
	authorsRouter := AuthorsRouter{
		Repository: &MockSuccessfulAuthorRepository{},
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

	expected := `[{"id":1,"firstName":"name1","lastName":"lastName1"}]`

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func TestErrorGetAuthorsHandler(t *testing.T) {
	authorsRouter := AuthorsRouter{
		Repository: &MockErrorAuthorRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/authors", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(authorsRouter.GetAuthorsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
