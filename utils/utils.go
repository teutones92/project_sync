package utils

import (
	"app/db_connection/tables/session_crud"
	"app/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Authorization", "")
}

func CheckHeader(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		// http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		res := models.StatusCode{StatusCode: http.StatusUnsupportedMediaType, StatusCodeMessage: "Content-Type must be application/json"}
		json.NewEncoder(w).Encode(res)
		log.Println("Content-Type must be application/json")
		return false
	}
	return true
}

func CheckMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		// http.Error(w, "Method not allowed asd", http.StatusMethodNotAllowed)
		var status_code = models.StatusCode{StatusCode: http.StatusMethodNotAllowed, StatusCodeMessage: "Method not allowed"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Method not allowed")
		return false
	}
	return true
}

func VerifyToken(w http.ResponseWriter, r *http.Request) bool {
	token := r.Header.Get("Authorization")
	if token == "" {
		status_code := models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Token is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Token is required")
		return false
	}
	session, err := session_crud.ReadSession(token)
	if err != nil {
		status_code := models.StatusCode{StatusCode: http.StatusUnauthorized, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err.Error())
		return false
	}
	if session.ID == 0 {
		status_code := models.StatusCode{StatusCode: http.StatusUnauthorized, StatusCodeMessage: "Invalid token"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Invalid token")
		return false
	}
	return true
}

func Validate(u models.User, from_login bool) error {
	if u.Username == "" && !from_login {
		return errors.New("username is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.DarkMode == nil && !from_login {
		return errors.New("dark mode is required as boolean")
	}
	return nil
}
