package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:temp@localhost/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		// Create database if not exists, to store URLs
		createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS URL_SHORTENER(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Execute create table statement
		_, err = createStatement.Exec()
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}
}
