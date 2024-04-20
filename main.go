package main

import (
	apis "app/APIs"
	"app/db_connection"
	"log"
)

func main() {
	log.Println("Hello, Welcome to Project Sync!")
	// Start the database connection
	db_connection.Init()
	// Start the server on port 8080 and handle requests
	e := apis.HandleAuthentication()
	if e != nil {
		log.Println("Error: ", e)
	}
	log.Println("Server is running on port 8080")
	// auth.Init()
}
