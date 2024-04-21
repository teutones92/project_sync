package apis

import (
	auth "app/authentication"
	project "app/db_connection/tables/project_crud"
	"log"

	"net/http"
)

var Host = "localhost"
var Port = ":8080"

func AuthApis() error {
	// Auth API
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/signup", auth.SignUp)
	mux.HandleFunc("/auth/login", auth.LogIn)
	mux.HandleFunc("/auth/logout", auth.LogOut)
	mux.HandleFunc("/auth/delete_account", auth.DeleteAccount)
	log.Printf("Server running on %s%s", Host, Port)
	err := http.ListenAndServe(Port, mux)
	return err
}

func ProjectsApis() error {
	// Projects API
	mux := http.NewServeMux()
	mux.HandleFunc("/projects/create", project.CreateProject)
	// mux.HandleFunc("/projects/read", project.ReadProject)
	// mux.HandleFunc("/projects/update", project.UpdateProject)
	// mux.HandleFunc("/projects/delete", project.DeleteProject)
	err := http.ListenAndServe(Port, mux)
	return err
}
