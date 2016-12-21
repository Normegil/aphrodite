package db

type DB string

func New(query string) *DB {
	db := DB(query)
	return &db
}
