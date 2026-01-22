package database

import (
	"database/sql" // Package for SQL database interactions
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite", "./kasir.db")

	if err != nil {
		log.Fatalf("error open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error ping database: %v", err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS produk (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		nama TEXT,
		harga REAL,
		stok INTEGER
	);
	
	CREATE TABLE IF NOT EXISTS categories (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}

	fmt.Println("Database is connected")
	return db
}
