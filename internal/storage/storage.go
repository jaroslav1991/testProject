package storage

import "database/sql"

type Db interface {
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
}

type Storage struct {
	Db Db
}

func NewStorage(db Db) *Storage {
	return &Storage{Db: db}
}
