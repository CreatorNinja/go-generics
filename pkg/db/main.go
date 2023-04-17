package db

import (
	"database/sql"
	"fmt"
	"log"

	// postgresql driver
	"os"

	_ "github.com/lib/pq"
)

// Database main db pointer
var Database *sql.DB

// InitializeDB initializes db
func InitializeDB() {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatalf("ERROR: Failed connecting to Database:\n %v\n", err)
		os.Exit(0)
	}

	Database = db
}
