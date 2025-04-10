package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// InitDB initializes the database connection pool
func InitDB() {
	dsn := os.Getenv("DATABASE_URL") // Make sure this is set in your .env or system
	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

// CloseDB gracefully closes the DB connection pool
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
