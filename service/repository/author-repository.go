package repository

import (
	"backend-api/model"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type AuthorRepository interface {
	GetAuthors() ([]model.Author, error)
}

type AuthorRepositoryDB struct {
	DB *sql.DB
}

func (r *AuthorRepositoryDB) GetAuthors() ([]model.Author, error) {
	rows, err := sq.Select("*").From("author").RunWith(r.DB).Query()
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)
	for rows.Next() {
		var author model.Author
		err = rows.Scan(&author.Id, &author.FirstName, &author.LastName)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}
