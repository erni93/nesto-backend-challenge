package repository

import (
	"backend-api/model"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type GenreRepository interface {
	GetGenres() ([]model.Genre, error)
}

type GenreRepositoryDB struct {
	DB *sql.DB
}

func (r *GenreRepositoryDB) GetGenres() ([]model.Genre, error) {
	rows, err := sq.Select("*").From("genre").RunWith(r.DB).Query()
	if err != nil {
		return nil, err
	}

	genres := make([]model.Genre, 0)
	for rows.Next() {
		var genre model.Genre
		err = rows.Scan(&genre.Id, &genre.Title)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}
