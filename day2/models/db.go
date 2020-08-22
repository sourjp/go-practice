package models

import (
	"database/sql"
	"fmt"
)

func NewDB() (*sql.DB, error) {
	var (
		driver   = "postgres"
		host     = "localhost"
		user     = "root"
		password = "root"
		dbname   = "todo"
		sslmode  = "disable"

		params = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, sslmode)
	)
	db, err := sql.Open(driver, params)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
