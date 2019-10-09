package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// NewSQLConnection creates sql connection
func NewSQLConnection() *sql.DB {
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
