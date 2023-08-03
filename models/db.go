package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:temp@localhost/postgres?sslmode=disable")
	if err != nil {
		return nil, err
		// } else {
		// 	// Create model for our URL service
		// 	stmt, err := db.Prepare("CREATE TABLE WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
		// 	if err != nil {
		// 		log.Println(err)
		// 		return nil, err
		// 	}
		// 	res, err := stmt.Exec()
		// 	log.Println(res)
		// 	if err != nil {
		// 		log.Println(err)
		// 		return nil, err
		// 	}
		// 	return db, nil
		// }
	}

	return db, nil

}
