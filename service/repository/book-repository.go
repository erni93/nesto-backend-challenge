package repository

import (
	"backend-api/database"
	"backend-api/model"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) GetBooks() ([]model.Book, error) {
	query := sq.Select("*").From("book")
	query = query.LeftJoin("genre on (book.genre_id = genre.id)")
	query = query.LeftJoin("author on (book.author_id = author.id)")
	rows, err := query.RunWith(r.DB).Query()
	if err != nil {
		return nil, err
	}

	books := make([]model.Book, 0)
	for rows.Next() {
		var book model.Book
		var genre model.Genre
		var author model.Author
		var ignoreColumn database.IgnoreColumn
		err = rows.Scan(
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
