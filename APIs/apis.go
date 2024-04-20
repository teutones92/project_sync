package apis

import (
	auth "app/authentication"

	"net/http"
)

var port = ":8080"

func HandleAuthentication() error {
	// Auth API
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/signup", auth.SignUp)
	mux.HandleFunc("/auth/login", auth.LogIn)
	mux.HandleFunc("/auth/logout", auth.LogOut)
	mux.HandleFunc("/auth/delete_account", auth.DeleteAccount)
	err := http.ListenAndServe(port, mux)
	return err
}
