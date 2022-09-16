package repository

import (
	"backend-api/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetSizes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id", "title", "min_pages", "max_pages"})
	rs.AddRow(1, "any", nil, nil)
	rs.AddRow(2, "Short story – up to 35 pages", 35, 84)

	mock.ExpectQuery("SELECT (.+) FROM size").WillReturnRows(rs)

	repository := SizeRepositoryDB{DB: db}
	sizes, err := repository.GetSizes()
	if err != nil {
		t.Fatalf("error retrieving list of sizes %s", err)
	}
	if len(sizes) != 2 {
		t.Errorf("invalid sizes len: got %d want %d", len(sizes), 2)
	}

	var tests = []struct {
		got  any
		want any
	}{
		{sizes[0].Id, int64(1)},
		{sizes[0].Title, "any"},
		{sizes[0].MinPages, (*database.NullInt32)(nil)},
		{sizes[0].MaxPages, (*database.NullInt32)(nil)},
		{sizes[1].Id, int64(2)},
		{sizes[1].Title, "Short story – up to 35 pages"},
		{sizes[1].MinPages.Int32, int32(35)},
		{sizes[1].MaxPages.Int32, int32(84)},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %v want %v", tt.got, tt.want)
		}
	}

}

func TestErrorGetSizes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT (.+) FROM size").WillReturnRows(rs)

	repository := SizeRepositoryDB{DB: db}
	sizes, err := repository.GetSizes()
	if err == nil {
		t.Fatal("expected error to not be nil")
	}
	if sizes != nil {
		t.Fatalf("expected sizes to be nil, got %v", sizes)
	}
}
