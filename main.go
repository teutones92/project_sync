package main

import (
	"fmt"
	// auth "app/authentication"
	"app/db_connection"
)

func main() {
	fmt.Println("Hello, Welcome to the Project Sync!")
	// Start the database connection
	db_connection.Init()
	// auth.Init()
}
