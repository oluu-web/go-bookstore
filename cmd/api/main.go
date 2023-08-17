// cmd/api/main.go

package main

import (
	"log"
	"net/http"

	"bookstore/cmd/api/routes" // Import your routes package
)

func main() {
	router := routes.InitRoutes() // Call the InitRoutes function
	port := "8080"
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
