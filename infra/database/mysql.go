package database

import (
	"database/sql"
	"fmt"
	"log"
)

// NewMySQLConnection creates mysql connection
func NewMySQLConnection(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB Connected")

	return db
}
