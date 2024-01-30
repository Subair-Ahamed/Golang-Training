package main

import (
    "crypto/sha256"
    "crypto/subtle"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

type application struct {
    auth struct {
        username string
        password string
    }
}

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the protected handler")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the unprotected handler")
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			// Calculate SHA-256 hashes for the provided and expected usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			// Use the subtle.ConstantTimeCompare() function to check if the provided username and password hashes equal the 
			//expected username and password hashes
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})

func main() {
        app := new(application)
    
        app.auth.username = os.Getenv("AUTH_USERNAME")
        app.auth.password = os.Getenv("AUTH_PASSWORD")
    
        if app.auth.username == "" {
            log.Fatal("basic auth username must be provided")
        }
    
        if app.auth.password == "" {
            log.Fatal("basic auth password must be provided")
        }
    
        //http.NewServeMux() creates a new HTTP request multiplexer (ServeMux), 
        //which is used for routing incoming requests to the appropriate handler functions based on their paths.
        mux := http.NewServeMux()
        mux.HandleFunc("/unprotected", app.unprotectedHandler)
        mux.HandleFunc("/protected", app.basicAuth(app.protectedHandler))
    
        srv := &http.Server{
            Addr:         ":4000",
            Handler:      mux,
            IdleTimeout:  time.Minute,
            ReadTimeout:  10 * time.Second,
            WriteTimeout: 30 * time.Second,
        }
    
        log.Printf("starting server on https://localhost:4000/protected")
        err := srv.ListenAndServeTLS("./localhost.pem", "./localhost-key.pem")
        log.Fatal(err)
    }

}