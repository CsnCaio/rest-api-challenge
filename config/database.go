package config

import (
	"database/sql"
	"fmt"
	"os"
)

var DbConn *sql.DB

// This function will make a connection to the database only once.
func InitDatabase() {
	var err error

	connStr := os.Getenv("DATABASE_URL")
	databaseType := os.Getenv("DATABASE_TYPE")

	// open the connection to the database
	DbConn, err = sql.Open(databaseType, connStr)
	if err != nil {
		panic(err)
	}

	// check if the connection to the database is successful
	if err = DbConn.Ping(); err != nil {
		panic(err)
	}

	// this will be printed in the terminal, confirming the connection to the database
	fmt.Printf("%s: Successfully connected!\n", databaseType)
}
