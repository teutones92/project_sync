package main

import (
	apis "app/APIs"
	"app/db_connection"
	"fmt"
)

func main() {
	fmt.Println("Hello, Welcome to the Project Sync!")
	// Start the database connection
	db_connection.Init()
	// Start the server on port 8080 and handle requests
	e := apis.HandleAuthentication()
	if e != nil {
		fmt.Println("Error: ", e)
	}
	fmt.Println("Server is running on port 8080")
	// auth.Init()
}
