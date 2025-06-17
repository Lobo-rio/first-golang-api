package database

import (
	"database/sql"
	"modules/src/config"

	_ "github.com/go-sql-driver/mysql"
)

// Connect abre a conexão com o banco de dados e a retorna
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
