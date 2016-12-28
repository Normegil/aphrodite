package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewConnection(dbType string, url string) (DB, error) {
	db, err := sql.Open(dbType, url)
	if err != nil {
		return DB{}, err
	}
	return DB{db}, nil
}
