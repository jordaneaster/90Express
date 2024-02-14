package main

import (
	"fmt"
	"log"
	"net/http"
	"server"
)

func main() {
	// Setup HTTP server
	server.SetupServer()

	// Start the HTTP server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
