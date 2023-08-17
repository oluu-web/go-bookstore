// cmd/api/main.go

package main

import (
	"log"
	"net/http"

	"bookstore/cmd/api/models"
	"bookstore/cmd/api/routes" // Import your routes package
)

func main() {
	err := models.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	router := routes.InitRoutes() // Call the InitRoutes function
	port := "4000"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
