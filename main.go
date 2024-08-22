package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

// User model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Active   bool   `json:"active"`
}

var db *sql.DB
var err error

func main() {
	// Define the connection string (replace with your own details)
	connString := "server=LAPTOP-BDSD6D7I\\SQLEXPRESS01;database=go_users;port=1433"

	// Connect to the database
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}

	fmt.Println("Connected to MS SQL Database!")

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{username}", getUser).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{username}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{username}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Get All Users (Read)
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT username, password, active FROM Users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.Password, &user.Active)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// Get Single User (Read)
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user User
	err := db.QueryRow("SELECT username, password, active FROM Users WHERE username = @p1", params["username"]).Scan(&user.Username, &user.Password, &user.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Create New User (Create)
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Exec("INSERT INTO Users (username, password, active) VALUES (@p1, @p2, @p3)", user.Username, user.Password, user.Active)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update User (Update)
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	_, err := db.Exec("UPDATE Users SET password = @p1, active = @p2 WHERE username = @p3", user.Password, user.Active, params["username"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Delete User (Delete)
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	_, err := db.Exec("DELETE FROM Users WHERE username = @p1", params["username"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}
