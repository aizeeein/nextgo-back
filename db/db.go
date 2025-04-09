package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres password=Poiuyt123 dbname=authjs sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	log.Println(" Connected to DB")
}

