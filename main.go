package main

import (
	"log"
	"net/http"
	"next-go/db"
	"next-go/handlers"

	"github.com/joho/godotenv"	
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	defer db.CloseDB()

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UserByIDHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
