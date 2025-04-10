package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"next-go/db"
	"next-go/models"

	"next-go/utils"

	"github.com/lucsky/cuid"
)

// Get /users or POST /users

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w, r)

	if r.Method == http.MethodOptions {
		return
	}

	switch r.Method {
	case http.MethodGet:
		getAllUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	w.Header().Set("Content-Type", "application/json")
}

// Get /users/:id, PUT /users/:id, DELETE /users/:id

func UserByIDHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w, r)

	if r.Method == http.MethodOptions {
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getUserByID(w, r, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(r.Context(), `SELECT id, name, email FROM "user"`)
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "Error scanning user", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(users)
}

func getUserByID(w http.ResponseWriter, r *http.Request, id string) {	
	var u models.User
	err := db.DB.QueryRow(r.Context(),`SELECT id, name, email FROM "user" WHERE id=$1`, id).Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(u)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println("DECODE ERROR:", err)
		return
	}

	u.ID = cuid.New()

	_, err := db.DB.Exec(r.Context(),
	`INSERT INTO "user" (id, name, email, password) VALUES ($1, $2, $3, $4)`, u.ID, u.Name, u.Email, u.Password)
	if err != nil {
		http.Error(w, "Failed to create userrr", http.StatusInternalServerError)
		log.Println("DB INSERT ERROR:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
	"message": "User createdd",
	"id" : u.ID,
})
}

func updateUser(w http.ResponseWriter, r *http.Request, id string) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(r.Context(), `UPDATE "user" SET name=$1, email=$2 WHERE id=$3`, u.Name, u.Email, id)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated"})
}

func deleteUser(w http.ResponseWriter, r *http.Request, id string) {
	_, err := db.DB.Exec(r.Context(),`DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})

}