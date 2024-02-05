package controllers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    // Add more fields as needed
}

// Simulated user data stored in memory
var userData = map[int]User{
    1: {ID: 1, Name: "John"},
    2: {ID: 2, Name: "Jane"},
    // Add more users as needed
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handling GetUser request")

    // Extract user ID from request URL
    // For example, if the request URL is /users/1, ID would be 1
    // You need to implement the logic to extract ID based on your URL structure
    // Here, we assume the ID is provided as a query parameter named "id"
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "User ID is required", http.StatusBadRequest)
        return
    }

    // Parse ID to integer
    userID := 0
    _, err := fmt.Sscanf(id, "%d", &userID)
    if err != nil {
        http.Error(w, "Invalid User ID", http.StatusBadRequest)
        return
    }

    // Fetch user data from simulated data store
    user, ok := userData[userID]
    if !ok {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Convert user data to JSON
    jsonResponse, err := json.Marshal(user)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Handling UpdateUser request")

    // Read request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    // Parse request body into User struct
    var updatedUser User
    err = json.Unmarshal(body, &updatedUser)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Update user data in simulated data store
    // Here, we simply replace the existing user data with updated user data
    userData[updatedUser.ID] = updatedUser

    // Respond with a success message
    response := map[string]string{"message": "User updated successfully"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

