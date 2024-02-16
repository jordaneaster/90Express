package server

import (
	"encoding/json"
	"fileprocessor"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"authenticate"
)

// SetupServer initializes the HTTP server and sets up the routes.
func SetupServer() {
	http.HandleFunc("/api/process", handleProcess)
	http.HandleFunc("/api/authenticate", handleAuthentication)
	http.HandleFunc("/dashboard", handleDashboard)
	http.HandleFunc("/profile", handleProfile)                  // Add route for the profile page
	http.HandleFunc("/api/check_profile", handleCheckProfile)   // Endpoint to check profile existence
	http.HandleFunc("/api/create_profile", handleCreateProfile) // Endpoint to create a profile
}

type Profile struct {
	UserID   int
	Username string
	Email    string
	FullName string
	Company  string
}

func handleCheckProfile(w http.ResponseWriter, r *http.Request) {
	// Extract the username from the query parameters
	username := r.URL.Query().Get("username")

	// Check if the profile exists
	profileExists := authenticate.CheckUserProfileExists(username)

	// Return the result as JSON
	response := struct {
		ProfileExists bool `json:"profileExists"`
	}{
		ProfileExists: profileExists,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling JSON: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing JSON response: %v\n", err)
	}
}

func handleCreateProfile(w http.ResponseWriter, r *http.Request) {
	// Get the username from the request
	username, err := getUsernameFromRequest(r)
	if err != nil {
		log.Printf("Error getting username from request: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a new profile for the user
	err = authenticate.CreateUserProfile(authenticate.Profile{Username: username})
	if err != nil {
		log.Printf("Error creating profile: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write a success response
	w.WriteHeader(http.StatusOK)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	// Serve the dashboard or profile page HTML file
	http.ServeFile(w, r, "go/server/templates/dashboard.html")
}

func getUsernameFromRequest(r *http.Request) (string, error) {
	// Parse the request body to get the username
	var requestBody struct {
		Username string `json:"username"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		return "", err
	}

	return requestBody.Username, nil
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
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Log the received username
	log.Printf("Received authentication request for username: %s\n", requestBody.Username)

	// Create a new AuthDB instance
	authDB, err := authenticate.NewAuthDB()
	if err != nil {
		log.Printf("Error creating AuthDB: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Perform authentication logic (e.g., check credentials against a database)
	response, err := authDB.AuthenticateUserDB(requestBody.Username, requestBody.Password)
	if err != nil {
		log.Printf("Error authenticating user: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Add the username to the response
	response.Username = requestBody.Username

	// Marshal the response into JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling JSON: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing JSON response: %v\n", err)
	}
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	// Get the username from the query parameters
	username := r.URL.Query().Get("username")

	// Render the profile template with the username
	tpl, err := template.ParseFiles("go/server/templates/profile.html")
	if err != nil {
		log.Printf("Error parsing profile template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
	}{
		Username: username,
	}

	if err := tpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
