package router

import (
	"backend-api/database"
	"backend-api/model"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockSuccessfulEraRepository struct {
}

func (r *MockSuccessfulEraRepository) GetEras() ([]model.Era, error) {
	year := &database.NullInt32{NullInt32: sql.NullInt32{Int32: 2022, Valid: true}}
	return []model.Era{{Id: 1, Title: "title1", MinYear: year, MaxYear: year}}, nil
}

type MockErrorEraRepository struct {
}

func (r *MockErrorEraRepository) GetEras() ([]model.Era, error) {
	return nil, errors.New("error")
}

func TestSuccessfulGetErasHandler(t *testing.T) {
	ErasRouter := ErasRouter{
		Repository: &MockSuccessfulEraRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/eras", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ErasRouter.GetErasHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"title1","minYear":2022,"maxYear":2022}]`

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func TestErrorGetErasHandler(t *testing.T) {
	ErasRouter := ErasRouter{
		Repository: &MockErrorEraRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/eras", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ErasRouter.GetErasHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
