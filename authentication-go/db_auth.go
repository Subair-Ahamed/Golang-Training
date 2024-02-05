package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "123"
	dbName     = "authentication_db"
)

var db *sql.DB

func main() {
	// Connect to the database
	var err error
	//sql.Open() opens a new database connection
	db, err = sql.Open(dbDriver, fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Serve signup page and handle signing up requests
	http.HandleFunc("/signup", signupHandler)
	
	// Serve login page and handle login requests
	http.HandleFunc("/login", loginHandler)

	// Serve profile page and handle profile requests
	http.HandleFunc("/profile", profileHandler)

	// Start server
	fmt.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed : Change to POST method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data from the request.
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get username, password and email from form data
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")

	// Insert the new user into the database
	_, err = db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, password, email)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		log.Println("Error inserting user:", err)
		return
	}

	// Registration successful
	fmt.Fprintf(w, "Registration successful for user: %s", username)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed: Change to POST method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data from the request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// This retrieves the username and password from the form data
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// Query the database for the user
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		log.Println("Error querying user:", err)
		return
	}

	//This compares the stored password with the provided password. 
	//If they don't match, an "Invalid password" error is returned.
	if storedPassword != password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Authentication successful
	fmt.Fprintf(w, "Welcome, %s!", username)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    // Assume the user is already authenticated (you might want to add authentication logic here)
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed : Change to GET method", http.StatusMethodNotAllowed)
		return
	}

    
    // Get the username from the request
    username := r.URL.Query().Get("username")
    
    // Query the database for additional user profile information
    var email string
    err := db.QueryRow("SELECT email FROM users WHERE username = ?", username).Scan(&email)
    if err != nil {
        http.Error(w, "Failed to retrieve user profile", http.StatusInternalServerError)
        log.Println("Error querying user profile:", err)
        return
    }
    
    // Display user profile information
	fmt.Fprintf(w,"Welcome %s to the profile page!\n", username)
    fmt.Fprintf(w, "Username: %s\n", username)
    fmt.Fprintf(w, "Email: %s\n", email)
    // Add more profile information if needed
}



