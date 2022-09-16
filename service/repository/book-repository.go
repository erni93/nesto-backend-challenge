package repository

import (
	"backend-api/database"
	"backend-api/model"
	"database/sql"
	"strconv"

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

type BookRepository interface {
	GetBooks(filters GetBooksFilters) ([]model.Book, error)
}

type BookRepositoryDB struct {
	DB *sql.DB
}

func (r *BookRepositoryDB) GetBooks(filters GetBooksFilters) ([]model.Book, error) {
	query := buildQuery(filters)

	rows, err := query.RunWith(r.DB).PlaceholderFormat(sq.Dollar).Query()
	if err != nil {
		return nil, err
	}

	return readRows(rows)
}

func buildQuery(filters GetBooksFilters) sq.SelectBuilder {
	query := sq.Select("*").From("book")
	query = query.LeftJoin("genre on (book.genre_id = genre.id)")
	query = query.LeftJoin("author on (book.author_id = author.id)")
	query = query.OrderBy("book.rating DESC")

	whereFilters := sq.And{}

	if filters.Genres != nil && len(filters.Genres) > 0 {
		whereFilters = append(whereFilters, sq.Eq{"book.genre_id": filters.Genres})
	}

	if filters.Authors != nil && len(filters.Authors) > 0 {
		whereFilters = append(whereFilters, sq.Eq{"book.author_id": filters.Authors})
	}

	if filters.MinPages != nil {
		whereFilters = append(whereFilters, sq.GtOrEq{"book.pages": strconv.Itoa(int(*filters.MinPages))})
	}

	if filters.MaxPages != nil {
		whereFilters = append(whereFilters, sq.LtOrEq{"book.pages": strconv.Itoa(int(*filters.MaxPages))})
	}

	if filters.MinYear != nil {
		whereFilters = append(whereFilters, sq.GtOrEq{"book.year_published": strconv.Itoa(int(*filters.MinYear))})
	}

	if filters.MaxYear != nil {
		whereFilters = append(whereFilters, sq.LtOrEq{"book.year_published": strconv.Itoa(int(*filters.MaxYear))})
	}

	if filters.Limit != nil {
		query = query.Limit(uint64(*filters.Limit))
	}

	if len(whereFilters) > 0 {
		query = query.Where(whereFilters)
	}

	return query
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
