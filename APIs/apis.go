package apis

import (
	auth "app/authentication"
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
