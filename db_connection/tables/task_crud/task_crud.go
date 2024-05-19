package task_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

// CreateTask function is used to create a task in the tasks table
func CreateTaskAPI(w http.ResponseWriter, r *http.Request) {
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
	db := db_connection.Database
	var task models.Task
	// Decode the request body into a task struct
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Insert data into the tasks table
	_, er := db.Exec(`
		INSERT INTO tasks (project_id, task_name, description, status_id,
			priority, assigned_user, dead_line, image_path)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		task.ProjectID,
		task.TaskName,
		task.Description,
		task.StatusID,
		task.Priority,
		task.AssignedUser,
		task.Deadline,
		task.ImagePath)
	if er != nil {
		log.Printf("Error inserting data into tasks table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error creating task."})

		return
	}

	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Task created."})
}

// ReadTaskByProjectID function is used to read a task by project ID
func ReadTaskByProjectIDAndStatusIdAPI(w http.ResponseWriter, r *http.Request) {
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
	db := db_connection.Database
	// Decode the request body into a project struct
	newStruct := struct {
		ProjectID int `json:"project_id"`
		StatusID  int `json:"status_id"`
	}{}
	er := json.NewDecoder(r.Body).Decode(&newStruct)
	if er != nil {
		log.Printf("Error decoding request body: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Check if the project ID is 0
	if newStruct.ProjectID == 0 && newStruct.StatusID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID and Status ID not provided."})

		return
	}
	// Read data from the tasks table
	rows, err := db.Query("SELECT * FROM tasks WHERE project_id = $1 AND status_id = $2", newStruct.ProjectID, newStruct.StatusID)
	if err != nil {
		log.Printf("Error reading data from tasks table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading tasks."})

		return
	}
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		er := rows.Scan(
			&task.ID, &task.ProjectID,
			&task.TaskName, &task.Description,
			&task.StatusID, &task.Priority,
			&task.AssignedUser, &task.Deadline,
			&task.ImagePath,
		)
		if er != nil {
			log.Printf("Error scanning data from tasks table: %s", er)
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading tasks."})

			return
		}
	}

	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask function is used to update a task in the tasks table
func UpdateTaskAPI(w http.ResponseWriter, r *http.Request) {
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
	db := db_connection.Database
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if task.ID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Task ID not provided."})

		return
	}
	// Decode the request body into a task struct
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Update data in the tasks table
	_, er := db.Exec(`
		UPDATE tasks SET project_id = $1, task_name = $2, description = $3,
		status_id = $4, priority = $5, assigned_user = $6, dead_line = $7, image_path = $8
		WHERE task_id = $9`,
		task.ProjectID,
		task.TaskName,
		task.Description,
		task.StatusID,
		task.Priority,
		task.AssignedUser,
		task.Deadline,
		task.ImagePath,
		task.ID,
	)
	if er != nil {
		log.Printf("Error updating data in tasks table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error updating task."})

		return
	}

	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Task updated."})
}

// DeleteTask function is used to delete a task from the tasks table
func DeleteTaskAPI(w http.ResponseWriter, r *http.Request) {
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
	db := db_connection.Database
	// Decode the request body into a task struct
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if task.ID == 0 {
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Task ID not provided."})

		return
	}
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Delete data from the tasks table
	_, er := db.Exec("DELETE FROM tasks WHERE task_id = $1", task.ID)
	if er != nil {
		log.Printf("Error deleting data from tasks table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error deleting task."})

		return
	}

	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Task deleted."})
}

func DeleteTaskByProjectID(projectID int) {
	db := db_connection.Database
	_, er := db.Exec("DELETE FROM tasks WHERE project_id = $1", projectID)
	if er != nil {
		log.Printf("Error deleting data from tasks table: %s", er)

	}

}
