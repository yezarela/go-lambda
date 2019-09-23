package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// OpenSQLConnection creates sql connection
func OpenSQLConnection() *sql.DB {
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
