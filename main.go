package main

import (
	apis "app/APIs"
	"app/db_connection"
	"log"
)

func main() {
	log.Println("Hello, Welcome to Project Sync!")
	// Start the database connection
	dbe := db_connection.Init()
	if dbe != nil {
		panic(dbe)
	}
	// Start the Auth API
	e := apis.StartServer()
	if e != nil {
		panic(e)
	}
}
