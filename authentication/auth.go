package auth

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/gorilla/mux"
)

func checkHeader(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		// http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		res := models.StatusCode{StatusCode: http.StatusUnsupportedMediaType, StatusCodeMessage: "Content-Type must be application/json"}
		json.NewEncoder(w).Encode(res)
		fmt.Println("Content-Type must be application/json")
		return
	}

}

func validate(u models.User) error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// SignUp is a function that handles the sign up process
func SignUp(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	w.Header().Set("Content-Type", "application/json")
	// Check the request method
	if r.Method != "POST" {
		// http.Error(w, "Method not allowed asd", http.StatusMethodNotAllowed)
		status_code = models.StatusCode{StatusCode: http.StatusMethodNotAllowed, StatusCodeMessage: "Method not allowed"}
		json.NewEncoder(w).Encode(status_code)
		fmt.Println("Method not allowed")
		return
	}
	// Check the header
	checkHeader(w, r)
	// Parse the request body
	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		return
	}
	// Validate the user data
	e := validate(req)
	if e != nil {
		// http.Error(w, e.Error(), http.StatusBadRequest)
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: e.Error()}
		json.NewEncoder(w).Encode(status_code)
		fmt.Println(e)
		return
	}
	//** TO DO/
	// Verify if the user already exists in the database
	// If the user exists, send a response to the client indicating that the user already exists
	// If the user does not exist, create a new user in the database and generate a JWT token
	// Return the JWT token to the client
	hashedPassword, err := hasGenerator("password123")
	if err != nil {
		fmt.Println("Error generating hash: ", err)
		return
	}
	fmt.Println(hashedPassword)
	fmt.Println(req)
	// fmt.Println("User registered successfully")
	// Return a response to the client indicating that the user has been created
	status_code = models.StatusCode{StatusCode: http.StatusAccepted, StatusCodeMessage: "User created"}
	json.NewEncoder(w).Encode(status_code)

}

// Generate a hash for the password

// LogIn is a function that handles the log in process
func LogIn(w http.ResponseWriter, r *http.Request) {
}

// LogOut is a function that handles the log out process
func LogOut(w http.ResponseWriter, r *http.Request) {
}

func hasGenerator(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
