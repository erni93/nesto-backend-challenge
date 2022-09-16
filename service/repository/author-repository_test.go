package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id", "first_name", "last_name"})
	rs.AddRow(1, "Wendell", "Stackhouse")
	rs.AddRow(2, "Amelia", "Wangerin, Jr.")

	mock.ExpectQuery("SELECT (.+) FROM author").WillReturnRows(rs)

	repository := AuthorRepositoryDB{DB: db}
	authors, err := repository.GetAuthors()
	if err != nil {
		t.Fatalf("error retrieving list of authors %s", err)
	}
	if len(authors) != 2 {
		t.Errorf("invalid authors len: got %d want %d", len(authors), 2)
	}

	var tests = []struct {
		got  any
		want any
	}{
		{authors[0].Id, int64(1)},
		{authors[0].FirstName, "Wendell"},
		{authors[0].LastName, "Stackhouse"},
		{authors[1].Id, int64(2)},
		{authors[1].FirstName, "Amelia"},
		{authors[1].LastName, "Wangerin, Jr."},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %v want %v", tt.got, tt.want)
		}
	}

}

func TestErrorGetAuthors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT (.+) FROM author").WillReturnRows(rs)

	repository := AuthorRepositoryDB{DB: db}
	authors, err := repository.GetAuthors()
	if err == nil {
		t.Fatal("expected error to not be nil")
	}
	if authors != nil {
		t.Fatalf("expected authors to be nil, got %v", authors)
	}
}
