package apis

import (
	auth "app/authentication"

	"net/http"
)

func HandleAuthentication() error {
	// Auth API
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/signup", auth.SignUp)
	mux.HandleFunc("/auth/login", auth.LogIn)
	mux.HandleFunc("/auth/logout", auth.LogOut)
	err := http.ListenAndServe(":8080", mux)
	return err
}
