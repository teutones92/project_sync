package session_crud

import (
	"app/db_connection"
	"app/models"
	"fmt"
	"log"
)

func CreateSession(session models.Session) models.StatusCode {
	// Insert data into the sessions table
	db := db_connection.Database
	// Insert the session data into the database if the token does not exist
	_, err := db.Exec(`
		INSERT INTO sessions (user_id, token)
		VALUES ($1, $2)
		ON CONFLICT (token) DO NOTHING`,
		session.UserID, session.Token)
	if err != nil {
		log.Printf("Error inserting data into sessions table: %s", err)
	}

	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Session created."}
}

func ReadSession(token string) (models.Session, error) {
	db := db_connection.Database
	var session models.Session
	// Query the database to get the session data
	er := db.QueryRow("SELECT id, user_id, token, last_activity_timestamp, expiration_minutes"+
		" FROM sessions WHERE token = $1",
		token).Scan(
		&session.ID,
		&session.UserID,
		&session.Token,
		&session.LastActivityTimestamp,
		&session.ExpirationMinutes,
	)
	if er != nil {
		log.Printf("Error reading session data: %s", er)

		return session, er
	}

	return session, nil
}

func UpdateSession(session models.Session) models.StatusCode {
	// Update the session in the database
	db := db_connection.Database
	_, err := db.Exec(`
		UPDATE sessions SET last_activity_timestamp = CURRENT_TIMESTAMP WHERE token user_id = $1 $2`,
		session.Token, session.UserID)
	if err != nil {
		panic(fmt.Sprintf("Error updating session: %s", err))
	}

	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Session updated."}
}

func DeleteSession(token string) models.StatusCode {
	// Delete the session from the database
	db := db_connection.Database
	_, err := db.Exec("DELETE FROM sessions WHERE token = $1", token)
	if err != nil {
		log.Printf("Error deleting session: %s", err)

		return models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error deleting session."}
	}

	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Session deleted."}
}
