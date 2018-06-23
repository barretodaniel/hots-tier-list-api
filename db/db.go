package db

import (
	"database/sql"
)

var db *sql.DB

// Get provides the open db connection
func Get() *sql.DB {
	var err error
	if db == nil {
		connStr := "dbname=hots sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		err = db.Ping()

		if err != nil {
			panic(err)
		}
	}

	return db
}
