package utils

import (
	"app/models"
	"encoding/json"
	"log"
	"net/http"
)

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
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
