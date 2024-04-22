package user_crud

import (
	"app/db_connection"
	"app/db_connection/tables/session_crud"
	"app/models"
	"app/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(user_data models.User) models.StatusCode {
	// Insert data into the users table
	_, err := db_connection.GetDatabase().Exec(`
        INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`,
		user_data.Username, user_data.Email, user_data.PasswordHash)
	if err != nil {
		panic(fmt.Sprintf("Error inserting data into users table: users %s", err))
	}
	db_connection.GetDatabase().Close()
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "User created."}

}

func ReadUserAPI(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	utils.SetHeader(w)
	// Check the request method
	method := utils.CheckMethod(w, r, "PUT")
	if !method {
		return
	}
	// Check the header
	header := utils.CheckHeader(w, r)
	if !header {
		return
	}
	// Verify the token
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}
	var user models.User
	// Decode the request body into a user struct
	err := json.NewDecoder(r.Body).Decode(&user.ID)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
	}
	// Query the database to get the user data
	er := ReadUserByID(user.ID)
	if er != nil {
		log.Printf("Error getting user data: %v", er)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func ReadUserByID(user_id int) *models.User {

	db := db_connection.GetDatabase()
	var user models.User
	// Query the database to get the user data
	er := db.QueryRow("SELECT user_id, username, email, password_hash, user_avatar_path FROM users WHERE user_id = $1", user.ID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.UserAvatar,
	)
	if er != nil {
		log.Printf("Error getting user data: %s", er)
	}
	db.Close()
	return &user
}

func UpdateUser(user_data *models.User) models.StatusCode {
	// Update the user data in the database
	// email = $2, // user_data.Email, // no need to update email
	_, err := db_connection.GetDatabase().Exec(`
		UPDATE users SET username = $1, 
		user_avatar_path
		= $3 WHERE user_id = $4`,
		user_data.Username, user_data.UserAvatar, user_data.ID)
	if err != nil {
		panic(fmt.Sprintf("Error updating user data: %s", err))
	}
	db_connection.GetDatabase().Close()
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "User updated."}
}

func UpdateUserAPI(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	utils.SetHeader(w)
	// Check the request method
	method := utils.CheckMethod(w, r, "PUT")
	if !method {
		return
	}
	// Check the header
	header := utils.CheckHeader(w, r)
	if !header {
		return
	}
	// Verify the token
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}

	var user models.User
	// Decode the request body into a user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
	}
	// Update the user data in the database
	var status = UpdateUser(&user)
	json.NewEncoder(w).Encode(&status)

}

func DeleteUser(user_id int) models.StatusCode {
	// Delete the user from the database
	_, err := db_connection.GetDatabase().Exec("DELETE FROM users WHERE user_id = $1", user_id)
	if err != nil {
		panic(fmt.Sprintf("Error deleting user: %s", err))
	}
	db_connection.GetDatabase().Close()
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "User deleted."}
}

func ChangePasswordAPI(w http.ResponseWriter, r *http.Request) {
	// Set the response header
	utils.SetHeader(w)
	// Check the request method
	method := utils.CheckMethod(w, r, "POST")
	if !method {
		return
	}
	// Check the header
	header := utils.CheckHeader(w, r)
	if !header {
		return
	}
	// Verify the token
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}
	// Get the token from the request header
	token := r.Header.Get("Authorization")
	// Decode the request body into a user struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		status_code := models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Get the session data from the database
	session, err := session_crud.ReadSession(token)
	if err != nil {
		status_code := models.StatusCode{StatusCode: http.StatusUnauthorized, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err.Error())
		return
	}
	if session.ID > 0 {
		// Get the user data from the database
		var user_data = ReadUserByID(session.UserID)
		// Compare the password in the request body with the password in the database
		if user_data.PasswordHash == user.PasswordHash {
			// Update the user's password in the database
			status := UpdateUser(user_data)
			// Return a response to the client indicating that the password has been changed
			json.NewEncoder(w).Encode(status)
			log.Println("Password changed successfully")
		} else {
			// Return a response to the client indicating that the password is incorrect
			status_code := models.StatusCode{StatusCode: http.StatusUnauthorized, StatusCodeMessage: "Incorrect password"}
			json.NewEncoder(w).Encode(status_code)
			log.Println("Incorrect password")
		}
	}
}
