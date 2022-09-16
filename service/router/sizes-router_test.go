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

type MockSuccessfulSizeRepository struct {
}

func (r *MockSuccessfulSizeRepository) GetSizes() ([]model.Size, error) {
	pages := &database.NullInt32{NullInt32: sql.NullInt32{Int32: 100, Valid: true}}
	return []model.Size{{Id: 1, Title: "title1", MinPages: pages, MaxPages: pages}}, nil
}

type MockErrorSizeRepository struct {
}

func (r *MockErrorSizeRepository) GetSizes() ([]model.Size, error) {
	return nil, errors.New("error")
}

func TestSuccessfulGetSizesHandler(t *testing.T) {
	SizesRouter := SizesRouter{
		Repository: &MockSuccessfulSizeRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/sizes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SizesRouter.GetSizesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"title1","minPages":100,"maxPages":100}]`

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func TestErrorGetSizesHandler(t *testing.T) {
	SizesRouter := SizesRouter{
		Repository: &MockErrorSizeRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/sizes", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SizesRouter.GetSizesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
