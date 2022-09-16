package repository

import (
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{
		"id", "title", "year_published", "rating", "pages", "genre_id", "author_id",
		"fk_genre_id", "fk_genre_title",
		"fk_author_id", "fk_author_first_name", "fk_author_last_name"},
	)
	rs.AddRow(1, "Alanna Saves the Day", 1972, 1.62, 169, 8, 6, 8, "Childrens", 6, "Bernard", "Hopf")
	rs.AddRow(2, "Adventures of Kaya", 1999, 2.13, 619, 1, 40, 1, "Young Adult", 40, "Wendell", "Stackhouse")

	mock.ExpectQuery(regexp.QuoteMeta(strings.TrimSpace(`
	SELECT * FROM book
	LEFT JOIN genre on (book.genre_id = genre.id)
	LEFT JOIN author on (book.author_id = author.id)
	ORDER BY book.rating DESC
	`))).WillReturnRows(rs)

	repository := BookRepositoryDB{DB: db}
	books, err := repository.GetBooks(GetBooksFilters{})
	if err != nil {
		t.Fatalf("error retrieving list of books %s", err)
	}
	if len(books) != 2 {
		t.Errorf("invalid books len: got %d want %d", len(books), 2)
	}

	var tests = []struct {
		got  any
		want any
	}{
		{books[0].Id, int64(1)},
		{books[0].Title, "Alanna Saves the Day"},
		{books[0].YearPublished, 1972},
		{books[0].Rating, float32(1.62)},
		{books[0].Pages, 169},
		{books[0].Genre.Id, int64(8)},
		{books[0].Genre.Title, "Childrens"},
		{books[0].Author.Id, int64(6)},
		{books[0].Author.FirstName, "Bernard"},
		{books[0].Author.LastName, "Hopf"},
		{books[1].Id, int64(2)},
		{books[1].Title, "Adventures of Kaya"},
		{books[1].YearPublished, 1999},
		{books[1].Rating, float32(2.13)},
		{books[1].Pages, 619},
		{books[1].Genre.Id, int64(1)},
		{books[1].Genre.Title, "Young Adult"},
		{books[1].Author.Id, int64(40)},
		{books[1].Author.FirstName, "Wendell"},
		{books[1].Author.LastName, "Stackhouse"},
	}

	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("got %v want %v", tt.got, tt.want)
		}
	}

}

func TestErrorGetBooks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rs := mock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT (.+) FROM book").WillReturnRows(rs)

	repository := BookRepositoryDB{DB: db}
	books, err := repository.GetBooks(GetBooksFilters{})
	if err == nil {
		t.Fatal("expected error to not be nil")
	}
	if books != nil {
		t.Fatalf("expected books to be nil, got %v", books)
	}
}
