package repository

import (
	"backend-api/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetEras(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id", "title", "min_year", "max_year"})
	rs.AddRow(1, "Any", nil, nil)
	rs.AddRow(2, "Classic", nil, 1969)

	mock.ExpectQuery("SELECT (.+) FROM era").WillReturnRows(rs)

	repository := EraRepositoryDB{DB: db}
	eras, err := repository.GetEras()
	if err != nil {
		t.Fatalf("error retrieving list of eras %s", err)
	}
	if len(eras) != 2 {
		t.Errorf("invalid eras len: got %d want %d", len(eras), 2)
	}

	var tests = []struct {
		got  any
		want any
	}{
		{eras[0].Id, int64(1)},
		{eras[0].Title, "Any"},
		{eras[0].MinYear, (*database.NullInt32)(nil)},
		{eras[0].MaxYear, (*database.NullInt32)(nil)},
		{eras[1].Id, int64(2)},
		{eras[1].Title, "Classic"},
		{eras[1].MinYear, (*database.NullInt32)(nil)},
		{eras[1].MaxYear.Int32, int32(1969)},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %v want %v", tt.got, tt.want)
		}
	}

}

func TestErrorGetEras(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT (.+) FROM era").WillReturnRows(rs)

	repository := EraRepositoryDB{DB: db}
	eras, err := repository.GetEras()
	if err == nil {
		t.Fatal("expected error to not be nil")
	}
	if eras != nil {
		t.Fatalf("expected eras to be nil, got %v", eras)
	}
}
