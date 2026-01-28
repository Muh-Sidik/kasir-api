package database

import (
	"database/sql" // Package for SQL database interactions
	"fmt"
	"log"
	"time"

	"github.com/Muh-Sidik/kasir-api/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(e *config.Env) *sql.DB {

	db, err := sql.Open("pgx", e.DB_URL)

	if err != nil {
		log.Fatalf("error open database: %v", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("error ping database: %v", err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS product (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(255),
		price REAL,
		stock INTEGER,
		category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS categories (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(255),
		description TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
	}

	fmt.Println("Database is connected")
	return db
}
