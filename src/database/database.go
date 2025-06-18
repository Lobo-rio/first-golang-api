package database

import (
	"database/sql"
	"modules/src/config"

	_ "github.com/go-sql-driver/mysql"
)

// Connect function that opens the connection to the database and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnectionDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
