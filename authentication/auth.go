package auth

import (
	"app/db_connection"
	"app/db_connection/session_crud"
	"app/db_connection/user_crud"
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func checkHeader(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		// http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		res := models.StatusCode{StatusCode: http.StatusUnsupportedMediaType, StatusCodeMessage: "Content-Type must be application/json"}
		json.NewEncoder(w).Encode(res)
		log.Println("Content-Type must be application/json")
		return false
	}
	return true
}

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		// http.Error(w, "Method not allowed asd", http.StatusMethodNotAllowed)
		var status_code = models.StatusCode{StatusCode: http.StatusMethodNotAllowed, StatusCodeMessage: "Method not allowed"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Method not allowed")
		return false
	}
	return true
}

func validate(u models.User, from_login bool) error {
	if u.Username == "" && !from_login {
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

// ** [SignUp] is a function that handles the sign up process
func SignUp(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	setHeader(w)
	// Check the request method
	checkMethod(w, r, "POST")
	// Check the header
	checkHeader(w, r)
	var req models.User
	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		return
	}
	// Validate the user data
	e := validate(req, false)
	if e != nil {
		// http.Error(w, e.Error(), http.StatusBadRequest)
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: e.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(e)
		return
	}
	// Verify if the user already exists in the database
	user_exist := checkUserExists(req.Email)
	if user_exist {
		status_code = models.StatusCode{StatusCode: http.StatusConflict, StatusCodeMessage: "User already exists"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("User already exists")
		return
	}
	// Generate a hash for the password
	hashedPassword, _ := hasGenerator(req.Password)
	req.PasswordHash = hashedPassword
	resp := user_crud.CreateUser(req)
	// Return a response to the client indicating that the user has been created
	json.NewEncoder(w).Encode(resp)
	log.Println("User created successfully")
}

// ** [LogIn] is a function that handles the log in process
func LogIn(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	setHeader(w)
	// Check the request method
	checkMethod(w, r, "POST")
	// Check the header
	checkHeader(w, r)
	var req models.User

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		return
	}

	// Validate the user data
	e := validate(req, true)
	if e != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: e.Error()}
		json.NewEncoder(w).Encode(status_code)
		return
	}

	// Verify if the user exists in the database
	userExists, userID := checkUserCredentials(req.Email, req.Password)
	if !userExists {
		status_code = models.StatusCode{StatusCode: http.StatusUnauthorized, StatusCodeMessage: "Invalid credentials"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Invalid credentials")
		return
	}

	// Get the user data from the database
	user := user_crud.ReadUser(userID)
	// Generate a JWT token
	token, err := generateJWT(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()})
		return
	}
	session := models.Session{
		UserID: userID,
		Token:  token,
	}
	// Save the session in the database
	session_crud.CreateSession(session)
	// Return the JWT token to the client
	session = session_crud.ReadSession(token)
	json.NewEncoder(w).Encode(session)
}

// ** [LogOut] is a function that handles the log out process
func LogOut(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	setHeader(w)
	// Check the request method
	method := checkMethod(w, r, "POST")
	if !method {
		return
	}
	// Check the header
	header := checkHeader(w, r)
	if !header {
		return
	}
	// Get the token from the request header
	token := r.Header.Get("Authorization")
	if token == "" {
		// http.Error(w, "Token is required", http.StatusBadRequest)
		status_code := models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Token is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Token is required")
		return
	}
	// Delete the session from the database
	status := session_crud.DeleteSession(token)
	// Return a response to the client indicating that the user has been logged out
	json.NewEncoder(w).Encode(status)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	setHeader(w)
	// Check the request method
	method := checkMethod(w, r, "DELETE")
	if !method {
		return
	}
	// Check the header
	header := checkHeader(w, r)
	if !header {
		return
	}
	// Get the token from the request header
	token := r.Header.Get("Authorization")
	if token == "" {
		// http.Error(w, "Token is required", http.StatusBadRequest)
		status_code := models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Token is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Token is required")
		return
	}
	// Get the session data from the database
	session := session_crud.ReadSession(token)
	// Delete the user from the database
	status := user_crud.DeleteUser(session.UserID)
	// Delete the session from the database
	session_crud.DeleteSession(token)
	// Return a response to the client indicating that the user has been deleted
	json.NewEncoder(w).Encode(status)
}

// Local function to check if the user credentials are valid
func checkUserCredentials(email, password string) (bool, int) {
	db := db_connection.GetDatabase()
	var userID int
	var hashedPassword string
	// Query the database to get the user ID and hashed password
	er := db.QueryRow("SELECT user_id, password_hash FROM users WHERE email = $1", email).Scan(&userID, &hashedPassword)
	if er != nil {
		return false, 0
	}
	// Compare the hashed password with the password provided by the user
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// If the passwords do not match, return false
		return false, 0
	}
	db.Close()
	// If the passwords match, return true
	return true, userID

}

// Local function to check if the user already exists in the database
func checkUserExists(email string) bool {
	db := db_connection.GetDatabase()
	var count int
	// er := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1 OR email = $2", user_data.Username, user_data.Email).Scan(&count)
	er := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", email).Scan(&count)
	if er != nil {
		panic(fmt.Sprintf("Error checking if user exists: %s", er))
	}
	db.Close()
	// If count is greater than zero, indicate that the user already exists
	return count > 0
}

// Local function to generate a hash for the password
func hasGenerator(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error generating hash: ", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func generateJWT(user models.User) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)
	// Create a new claim
	claims := token.Claims.(jwt.MapClaims)
	// Set the claim
	claims["user_id"] = user.UserID
	// Generate the token
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Println("Error generating JWT: ", err)
		return "", err
	}
	return tokenString, nil
}
