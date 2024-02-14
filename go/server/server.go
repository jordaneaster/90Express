package server

import (
	"encoding/json"
	"fileprocessor"
	"fmt"
	"net/http"
	"structs"

	"authenticate"
)

// SetupServer initializes the HTTP server and sets up the routes.
func SetupServer() {
	http.HandleFunc("/api/process", handleProcess)
	http.HandleFunc("/api/authenticate", handleAuthentication)
	http.HandleFunc("/dashboard", handleDashboard)
	http.HandleFunc("/api/profile", handleUserProfile) // Add this line for the profile endpoint
}

type Profile struct {
	UserID   int
	Username string
	Email    string
	FullName string
	Company  string
}

func handleUserProfile(w http.ResponseWriter, r *http.Request) {
	// Get the username from the request
	username, err := getUsernameFromRequest(r)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create the user profile
	profile := structs.Profile{
		Username: username,
	}

	// Convert structs.Profile to authenticate.Profile
	authProfile := authenticate.Profile(profile)

	// Create the user profile
	if err := authenticate.CreateUserProfile(authProfile); err != nil {
		http.Error(w, "Failed to create user profile", http.StatusInternalServerError)
		return
	}

	// Return a success response after creating the profile
	fmt.Fprintf(w, "User profile created for user: %s", username)
}

func getUsernameFromRequest(r *http.Request) (string, error) {
	// Parse the request body to get the username
	var requestBody struct {
		Username string `json:"username"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return "", err
	}

	return requestBody.Username, nil
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	// Serve the dashboard or profile page HTML file
	http.ServeFile(w, r, "go/server/templates/dashboard.html")
}

func handleProcess(w http.ResponseWriter, r *http.Request) {
	// Process files
	features, err := fileprocessor.ProcessFiles()
	if err != nil {
		fmt.Printf("Error processing files: %v\n", err) // Log the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Marshal the features into JSON
	jsonResponse, err := json.Marshal(features)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err) // Log the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate that the response body contains JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Printf("Error writing JSON response: %v\n", err) // Log the error
	}
}

func handleAuthentication(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the username and password
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Printf("Error decoding request body: %v\n", err) // Log the error
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Create a new AuthDB instance
	authDB, err := authenticate.NewAuthDB()
	if err != nil {
		fmt.Printf("Error creating AuthDB: %v\n", err) // Log the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Perform authentication logic (e.g., check credentials against a database)
	response, err := authDB.AuthenticateUserDB(requestBody.Username, requestBody.Password)
	if err != nil {
		fmt.Printf("Error authenticating user: %v\n", err) // Log the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Marshal the response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err) // Log the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Printf("Error writing JSON response: %v\n", err) // Log the error
	}
}
