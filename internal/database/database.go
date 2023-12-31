package database

import (
	"database/sql"
	"ec/internal/config"
	"fmt"
)

func Connect() (*sql.DB, error) {
	conf := config.GlobalConfig

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		conf.Store["database.user"],
		conf.Store["database.password"],
		conf.Store["database.dbname"],
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
