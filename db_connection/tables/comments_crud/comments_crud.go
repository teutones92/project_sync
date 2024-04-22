package comments_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

// CreateCommentAPI creates a new comment
func CreateCommentAPI(w http.ResponseWriter, r *http.Request) {
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
	// Get the database connection
	db := db_connection.GetDatabase()
	// Decode the request body into a comment struct
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		db.Close()
		return
	}
	// Check if the task ID, project ID and user ID are provided
	if comment.TaskID == 0 && comment.ProjectID == 0 && comment.UserID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Task ID, Project ID and User ID not provided."})
		db.Close()
		return
	}
	_, err = db.Exec("INSERT INTO comments ("+
		"task_id, "+
		"project_id, "+
		"user_id, "+
		"comment_text "+
		") VALUES ($1, $2, $3, $4)",
		comment.TaskID,
		comment.ProjectID,
		comment.UserID,
		comment.CommentText)
	if err != nil {
		log.Printf("Error inserting data into comments table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error creating comment."})
		db.Close()
		return
	}
	json.NewEncoder(w).Encode(models.StatusCode{StatusCode: 200, StatusCodeMessage: "Comment created successfully."})
}

// ReadComments reads a comment from the database
func ReadCommentsAPI(w http.ResponseWriter, r *http.Request) {
	utils.SetHeader(w)
	// Check the request method
	method := utils.CheckMethod(w, r, "GET")
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
	// Get the database connection
	db := db_connection.GetDatabase()
	// Get the task ID from the query parameter
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		db.Close()
		return
	}
	if comment.TaskID == 0 && comment.ProjectID == 0 && comment.UserID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Task ID, Project ID and User ID required."})
		db.Close()
		return
	}

	// Query the database
	// rows, err := db.Query("SELECT * FROM comments WHERE task_id = $1", taskID)
	rows, err := db.Query("SELECT * FROM comments WHERE task_id = $1 AND project_id = $2 AND user_id = $3", comment.TaskID, comment.ProjectID, comment.UserID)
	if err != nil {
		log.Printf("Error querying the database: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading comment."})
		db.Close()
		return
	}
	// Create a slice of comments
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ID, &comment.TaskID, &comment.ProjectID, &comment.UserID, &comment.CommentText, &comment.Timestamp)
		if err != nil {
			log.Printf("Error scanning rows: %s", err)
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading comment."})
			db.Close()
			return
		}
		comments = append(comments, comment)
	}
	json.NewEncoder(w).Encode(comments)
}

// UpdateCommentAPI updates a comment in the database
func UpdateCommentAPI(w http.ResponseWriter, r *http.Request) {
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
	// Get the database connection
	db := db_connection.GetDatabase()
	// Decode the request body into a comment struct
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		db.Close()
		return
	}
	// Check if the comment ID is provided
	if comment.ID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Comment ID not provided."})
		db.Close()
		return
	}
	_, err = db.Exec("UPDATE comments SET "+
		"comment_text = $1 "+
		"WHERE id = $2",
		comment.CommentText,
		comment.ID)
	if err != nil {
		log.Printf("Error updating data in comments table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error updating comment."})
		db.Close()
		return
	}
	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Comment updated successfully."})
}

// DeleteCommentAPI deletes a comment from the database
func DeleteCommentAPI(w http.ResponseWriter, r *http.Request) {
	utils.SetHeader(w)
	// Check the request method
	method := utils.CheckMethod(w, r, "DELETE")
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
	// Get the database connection
	db := db_connection.GetDatabase()
	// Decode the request body into a comment struct
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		db.Close()
		return
	}
	// Check if the comment ID is provided
	if comment.ID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Comment ID not provided."})
		db.Close()
		return
	}
	_, err = db.Exec("DELETE FROM comments WHERE id = $1", comment.ID)
	if err != nil {
		log.Printf("Error deleting data from comments table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error deleting comment."})
		db.Close()
		return
	}
	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Comment deleted successfully."})
}
