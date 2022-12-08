package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// NewMySQLConnection creates mysql connection
func NewMySQLConnection() *sql.DB {
	dbDriver := os.Getenv("DBDriverName")
	dbSource := os.Getenv("DBDataSourceName")

	db, err := sql.Open(dbDriver, dbSource)
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
