package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetGenres(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id", "title"})
	rs.AddRow(1, "Young Adult")
	rs.AddRow(2, "SciFi/Fantasy")

	mock.ExpectQuery("SELECT (.+) FROM genre").WillReturnRows(rs)

	repository := GenreRepositoryDB{DB: db}
	genres, err := repository.GetGenres()
	if err != nil {
		t.Fatalf("error retrieving list of Genre %s", err)
	}
	if len(genres) != 2 {
		t.Errorf("invalid genres len: got %d want %d", len(genres), 2)
	}

	var tests = []struct {
		got  any
		want any
	}{
		{genres[0].Id, int64(1)},
		{genres[0].Title, "Young Adult"},
		{genres[1].Id, int64(2)},
		{genres[1].Title, "SciFi/Fantasy"},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %v want %v", tt.got, tt.want)
		}
	}

}

func TestErrorGetGenres(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT (.+) FROM genre").WillReturnRows(rs)

	repository := GenreRepositoryDB{DB: db}
	genres, err := repository.GetGenres()
	if err == nil {
		t.Fatal("expected error to not be nil")
	}
	if genres != nil {
		t.Fatalf("expected genres to be nil, got %v", genres)
	}
}
