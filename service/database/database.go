package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("PSQL-HOST")
	port     = os.Getenv("PSQL-PORT")
	user     = os.Getenv("PSQL-USER")
	password = os.Getenv("PSQL-PASSWORD")
	dbname   = os.Getenv("PSQL-DBNAME")
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
