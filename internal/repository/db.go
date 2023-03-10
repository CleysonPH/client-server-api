package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const createCambiumTable = `
	CREATE TABLE IF NOT EXISTS cambium (
		id TEXT NOT NULL PRIMARY KEY,
		code TEXT NOT NULL,
		codein TEXT NOT NULL,
		name TEXT NOT NULL,
		high TEXT NOT NULL,
		low TEXT NOT NULL,
		var_bid TEXT NOT NULL,
		pct_change TEXT NOT NULL,
		bid TEXT NOT NULL,
		ask TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		create_date TEXT NOT NULL
	)
`

var (
	db  *sql.DB
	err error
)

func InitDb(dsn string) error {
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	if _, err := db.Exec(createCambiumTable); err != nil {
		return err
	}
	return nil
}

func CloseDb() error {
	return db.Close()
}
