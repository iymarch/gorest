package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type dataBaseStorage struct {
	db *sql.DB
}

func newDataBaseStorage() *dataBaseStorage {
	db, err := sql.Open("sqlite3", "test.db")
	fmt.Print(err)
	return &dataBaseStorage{
		db: db,
	}
}

func (dbs *dataBaseStorage) closeDataBaseStorage() {
	dbs.db.Close()
}
