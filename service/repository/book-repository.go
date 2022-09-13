package repository

import (
	"backend-api/database"
	"backend-api/model"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type GetBooksFilters struct {
	Authors  []int64
	Genres   []int64
	MinPages *int32
	MaxPages *int32
	MinYear  *int32
	MaxYear  *int32
	Limit    *int32
}

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) GetBooks(filters GetBooksFilters) ([]model.Book, error) {
	query := sq.Select("*").From("book")
	query = query.LeftJoin("genre on (book.genre_id = genre.id)")
	query = query.LeftJoin("author on (book.author_id = author.id)")
	rows, err := query.RunWith(r.DB).Query()
	if err != nil {
		return nil, err
	}

	return readRows(rows)
}

func readRows(rows *sql.Rows) ([]model.Book, error) {
	books := make([]model.Book, 0)
	for rows.Next() {
		var book model.Book
		var genre model.Genre
		var author model.Author
		var ignoreColumn database.IgnoreColumn
		err := rows.Scan(
			&book.Id, &book.Title, &book.YearPublished, &book.Rating, &book.Pages, ignoreColumn, ignoreColumn,
			&genre.Id, &genre.Title,
			&author.Id, &author.FirstName, &author.LastName,
		)
		if err != nil {
			return nil, err
		}
		book.Genre = &genre
		book.Author = &author
		books = append(books, book)
	}

	return books, nil
}
