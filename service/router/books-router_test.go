package router

import (
	"backend-api/model"
	"backend-api/repository"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var GetBooksJsonMockup = `[
	{"id":1,"yearPublished":2000,"rating":4.5,"pages":200,"title":"title1",
	"genre":{"id":1,"title":"title1"},
	"author":{"id":1,"firstName":"name1","lastName":"lastName1"}},
	{"id":2,"yearPublished":2005,"rating":5,"pages":500,"title":"title2",
	"genre":{"id":2,"title":"title2"},
	"author":{"id":2,"firstName":"name2","lastName":"lastName2"}},
	{"id":3,"yearPublished":2010,"rating":1,"pages":700,"title":"title3",
	"genre":{"id":3,"title":"title3"},
	"author":{"id":3,"firstName":"name3","lastName":"lastName3"}}]`

type MockSuccessfulBookRepository struct {
}

func (r *MockSuccessfulBookRepository) GetBooks(filters repository.GetBooksFilters) ([]model.Book, error) {
	return []model.Book{
		{
			Id: 1, YearPublished: 2000, Rating: 4.5, Pages: 200, Title: "title1",
			Genre:  &model.Genre{Id: 1, Title: "title1"},
			Author: &model.Author{Id: 1, FirstName: "name1", LastName: "lastName1"},
		},
		{
			Id: 2, YearPublished: 2005, Rating: 5, Pages: 500, Title: "title2",
			Genre:  &model.Genre{Id: 2, Title: "title2"},
			Author: &model.Author{Id: 2, FirstName: "name2", LastName: "lastName2"},
		},
		{
			Id: 3, YearPublished: 2010, Rating: 1, Pages: 700, Title: "title3",
			Genre:  &model.Genre{Id: 3, Title: "title3"},
			Author: &model.Author{Id: 3, FirstName: "name3", LastName: "lastName3"},
		},
	}, nil
}

type MockErrorBookRepository struct {
}

func (r *MockErrorBookRepository) GetBooks(filters repository.GetBooksFilters) ([]model.Book, error) {
	return nil, errors.New("error")
}

func compactJson(t *testing.T, content string) string {
	buffer := new(bytes.Buffer)
	err := json.Compact(buffer, []byte(content))
	if err != nil {
		t.Errorf("error compacting json: %v", err)
		return ""
	}
	return buffer.String()
}

func TestSuccessfulGetBooksHandler(t *testing.T) {
	BooksRouter := BooksRouter{
		Repository: &MockSuccessfulBookRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksRouter.GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := compactJson(t, GetBooksJsonMockup)

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func TestErrorGetBooksHandler(t *testing.T) {
	BooksRouter := BooksRouter{
		Repository: &MockErrorBookRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksRouter.GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func testSuccessfulFilter(t *testing.T, key string, value string) {
	BooksRouter := BooksRouter{
		Repository: &MockSuccessfulBookRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add(key, value)
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksRouter.GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := compactJson(t, GetBooksJsonMockup)

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func testErrorFilter(t *testing.T, key string, value string) {
	BooksRouter := BooksRouter{
		Repository: &MockSuccessfulBookRepository{},
	}

	req, err := http.NewRequest("GET", "/api/v1/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add(key, value)
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BooksRouter.GetBooksHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := compactJson(t, `{"message":"Some filters are not valid"}`)

	if want := strings.TrimSpace(rr.Body.String()); want != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", want, expected)
	}
}

func TestSuccessfulAuthorsFilter(t *testing.T) {
	testSuccessfulFilter(t, "authors", "1,3")
}

func TestErrorAuthorsFilter(t *testing.T) {
	testErrorFilter(t, "authors", "demo")
}

func TestSuccessfulGenresFilter(t *testing.T) {
	testSuccessfulFilter(t, "genres", "3,2")
}

func TestErrorGenresFilter(t *testing.T) {
	testErrorFilter(t, "genres", "demo")
}

func TestSuccessfulMinPagesFilter(t *testing.T) {
	testSuccessfulFilter(t, "min-pages", "200")
}

func TestErrorMinPagesFilter(t *testing.T) {
	testErrorFilter(t, "min-pages", "demo")
}

func TestSuccessfulMaxPagesFilter(t *testing.T) {
	testSuccessfulFilter(t, "max-pages", "200")
}

func TestErrorMaxPagesFilter(t *testing.T) {
	testErrorFilter(t, "max-pages", "demo")
}
func TestSuccessfulMinYearFilter(t *testing.T) {
	testSuccessfulFilter(t, "min-year", "2000")
}

func TestErrorMinYearFilter(t *testing.T) {
	testErrorFilter(t, "min-year", "demo")
}

func TestSuccessfulMaxYearFilter(t *testing.T) {
	testSuccessfulFilter(t, "max-year", "2010")
}

func TestErrorMaxYearFilter(t *testing.T) {
	testErrorFilter(t, "max-year", "demo")
}

func TestSuccessfulLimitFilter(t *testing.T) {
	testSuccessfulFilter(t, "limit", "10")
}

func TestErrorLimitFilter(t *testing.T) {
	testErrorFilter(t, "limit", "demo")
}
