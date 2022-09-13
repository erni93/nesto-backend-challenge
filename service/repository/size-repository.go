package repository

import (
	"backend-api/model"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type SizeRepository struct {
	DB *sql.DB
}

func (r *SizeRepository) GetSizes() ([]model.Size, error) {
	rows, err := sq.Select("*").From("size").RunWith(r.DB).Query()
	if err != nil {
		return nil, err
	}

	sizes := make([]model.Size, 0)
	for rows.Next() {
		var size model.Size
		err = rows.Scan(&size.Id, &size.Title, &size.MinPages, &size.MaxPages)
		if err != nil {
			return nil, err
		}
		sizes = append(sizes, size)
	}

	return sizes, nil
}
