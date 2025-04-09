package main

import (
	"log"
	"net/http"
	"next-go/db"
	"next-go/handlers"
)

func main() {
	db.InitDB()

	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/users/", handlers.UserByIDHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

