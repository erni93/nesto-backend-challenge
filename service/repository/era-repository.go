package repository

import (
	"backend-api/model"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type EraRepository interface {
	GetEras() ([]model.Era, error)
}

type EraRepositoryDB struct {
	DB *sql.DB
}

func (r *EraRepositoryDB) GetEras() ([]model.Era, error) {
	rows, err := sq.Select("*").From("era").RunWith(r.DB).Query()
	if err != nil {
		return nil, err
	}

	eras := make([]model.Era, 0)
	for rows.Next() {
		var era model.Era
		err = rows.Scan(&era.Id, &era.Title, &era.MinYear, &era.MaxYear)
		if err != nil {
			return nil, err
		}
		eras = append(eras, era)
	}

	return eras, nil
}
