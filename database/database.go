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

	fmt.Println("Database is connected")
	return db
}
