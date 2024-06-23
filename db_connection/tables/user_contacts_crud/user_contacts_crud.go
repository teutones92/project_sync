package user_contacts_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

// Function to validate a user contact
func validateUserContact(user_contact models.UserContact, validate_to_delete *bool) *models.StatusCode {
	var status_code models.StatusCode
	if user_contact.UserID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "User ID is required"}
	}
	if validate_to_delete != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Contact ID is required"}
	}
	if user_contact.ContactEmail == "" && validate_to_delete == nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Contact email is required"}
	}
	return &status_code
}

// Function to create a user contact
func CreateUserContactAPI(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request header is application/json
	if !utils.CheckHeader(w, r) {
		return
	}
	// Check if the request method is POST
	if !utils.CheckMethod(w, r, "POST") {
		return
	}
	// Decode the request body into a user_contact struct
	var user_contact models.UserContact
	err := json.NewDecoder(r.Body).Decode(&user_contact)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Validate the user_contact struct
	status := validateUserContact(user_contact, nil)
	if status != nil {
		json.NewEncoder(w).Encode(status)
		log.Println(err)
		return
	}
	// Insert the user_contact struct into the database
	db := db_connection.Database
	_, err = db.Exec("INSERT INTO user_contacts "+
		"(user_id, contact_id) VALUES ($1, $2)",
		user_contact.UserID, user_contact.ID)
	if err != nil {
		log.Printf("Error inserting data into user_contacts table: %s", err)
		status_code = models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error inserting data into user_contacts table."}
		json.NewEncoder(w).Encode(status_code)
		return
	}

	log.Println("User contact created.")
	status_code = models.StatusCode{StatusCode: 200, StatusCodeMessage: "User contact created."}
	json.NewEncoder(w).Encode(status_code)
}

// Function to read user contacts by user ID
func ReadUserContactByUserIdAPI(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request method is GET
	if !utils.CheckMethod(w, r, "GET") {
		return
	}
	// Get the user ID from the URL
	var user_contact models.UserContact
	err := json.NewDecoder(r.Body).Decode(&user_contact.ID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}

	if user_contact.ID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "User ID is required"}
		json.NewEncoder(w).Encode(status_code)
		return
	}
	// Read data from the user_contacts table
	rows, err := db_connection.Database.Query("SELECT * FROM user_contacts WHERE user_id = $1", user_contact.ID)
	if err != nil {
		log.Printf("Error reading data from user_contacts table: %s", err)
		status_code = models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading data from user_contacts table."}
		json.NewEncoder(w).Encode(status_code)
		return
	}
	defer rows.Close()
	var user_contacts []models.UserContact
	for rows.Next() {
		var user_contact models.UserContact
		err := rows.Scan(&user_contact.UserID, &user_contact.ID, &user_contact.ContactName, &user_contact.ContactEmail)
		if err != nil {
			log.Printf("Error scanning data from user_contacts table: %s", err)
			status_code = models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error scanning data from user_contacts table."}
			json.NewEncoder(w).Encode(status_code)
			return
		}
		user_contacts = append(user_contacts, user_contact)
	}
	log.Println("User contacts read.")
	json.NewEncoder(w).Encode(user_contacts)
}

// Function to update a user contact
func UpdateUserContactAPI(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request header is application/json
	if !utils.CheckHeader(w, r) {
		return
	}
	// Check if the request method is PUT
	if !utils.CheckMethod(w, r, "PUT") {
		return
	}
	// Decode the request body into a user_contact struct
	var user_contact models.UserContact
	err := json.NewDecoder(r.Body).Decode(&user_contact)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Validate the user_contact struct
	status := validateUserContact(user_contact, nil)
	if status != nil {
		json.NewEncoder(w).Encode(status)
		log.Println(err)
		return
	}
	// Update the user_contact struct in the database
	db := db_connection.Database
	_, err = db.Exec("UPDATE user_contacts SET "+
		"contact_name=$1, contact_email=$2 WHERE user_id=$3",
		user_contact.ContactName, user_contact.ContactEmail, user_contact.UserID)
	if err != nil {
		log.Printf("Error updating data in user_contacts table: %s", err)
		status_code = models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error updating data in user_contacts table."}
		json.NewEncoder(w).Encode(status_code)
		return
	}

	log.Println("User contact updated.")
	status_code = models.StatusCode{StatusCode: 200, StatusCodeMessage: "User contact updated."}
	json.NewEncoder(w).Encode(status_code)
}

// Function to delete a user contact
func DeleteUserContactAPI(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request header is application/json
	if !utils.CheckHeader(w, r) {
		return
	}
	// Check if the request method is DELETE
	if !utils.CheckMethod(w, r, "DELETE") {
		return
	}
	// Decode the request body into a user_contact struct
	var user_contact models.UserContact
	err := json.NewDecoder(r.Body).Decode(&user_contact)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Validate the user_contact struct
	var validate = true
	status := validateUserContact(user_contact, &validate)
	if status != nil {
		json.NewEncoder(w).Encode(status)
		log.Println(err)
		return
	}
	// Delete the user_contact struct from the database
	db := db_connection.Database
	_, err = db.Exec("DELETE FROM user_contacts WHERE user_id=$1 AND contact_id=$2",
		user_contact.UserID, user_contact.ID)
	if err != nil {
		log.Printf("Error deleting data from user_contacts table: %s", err)
		status_code = models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error deleting data from user_contacts table."}
		json.NewEncoder(w).Encode(status_code)
		return
	}

	log.Println("User contact deleted.")
	status_code = models.StatusCode{StatusCode: 200, StatusCodeMessage: "User contact deleted."}
	json.NewEncoder(w).Encode(status_code)
}
