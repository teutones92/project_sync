package user_crud

import (
	"app/db_connection"
	"app/models"
	"fmt"
	"log"
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

func ReadUser(user_id int) models.User {
	db := db_connection.GetDatabase()
	var user models.User
	// Query the database to get the user data
	er := db.QueryRow("SELECT user_id, username, email, user_avatar_path FROM users WHERE user_id = $1", user_id).Scan(&user.UserID, &user.Username, &user.Email, &user.UserAvatar)
	if er != nil {
		log.Printf("Error getting user data: %s", er)
	}
	db.Close()
	return user
}

func UpdateUser(user_data models.User) models.StatusCode {
	// Update the user data in the database
	// email = $2, // user_data.Email, // no need to update email
	_, err := db_connection.GetDatabase().Exec(`
		UPDATE users SET username = $1, 
		user_avatar_path
		= $3 WHERE user_id = $4`,
		user_data.Username, user_data.UserAvatar, user_data.UserID)
	if err != nil {
		panic(fmt.Sprintf("Error updating user data: %s", err))
	}
	db_connection.GetDatabase().Close()
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "User updated."}
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
