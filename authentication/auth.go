package auth

import (
	"encoding/json"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gorilla/mux"
)

type _SignUpResponse struct {
	Message string `json:"message"`
}

type _SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignIn is a function that handles the sign in process
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//** TO DO/
	// Validate user data
	// Verify if the user already exists in the database
	// If the user does not exist, create a new user in the database
	// Send a response to the client

	// Parse the request body
	var req _SignUpRequest
	res := _SignUpResponse{Message: "User registered successfully"}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}

// LogIn is a function that handles the log in process
func LogIn(w http.ResponseWriter, r *http.Request) {
}

// LogOut is a function that handles the log out process
func LogOut(w http.ResponseWriter, r *http.Request) {
}
