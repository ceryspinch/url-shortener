package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbUser     = "postgres"
	dbPassword = "temp"
	dbName     = "postgres"
)

func InitDB() (*sql.DB, error) {
	connectionCredentials := fmt.Sprintf("user=%s password=%s dbname= %s sslmode=disable", dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connectionCredentials)
	if err != nil {
		return nil, err
	}

	// Create database if not exists, to store URLs
	createStatement, err := db.Prepare("CREATE TABLE IF NOT EXISTS urls(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
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
