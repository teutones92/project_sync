package task_status_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

// CreateTaskStatusAPI function is used to create a task status in the task_status table
func CreateTaskStatusAPI(w http.ResponseWriter, r *http.Request) {
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
	var taskStatus models.TaskStatus
	// Decode the request body into a task_status struct
	err := json.NewDecoder(r.Body).Decode(&taskStatus)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Insert data into the task_status table
	_, er := db.Exec(`
		INSERT INTO task_status (project_id, user_id, status_name, status_description)
			VALUES ($1, $2, $3, $4)`,
		taskStatus.ProjectId,
		taskStatus.UserID,
		taskStatus.StatusName,
		taskStatus.StatusDescription,
	)
	if er != nil {
		log.Printf("Error inserting data into task_status table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error creating task status."})

		return
	}

}

// ReadTaskStatusByProjectIDApi function is used to read task status by project id from the task_status table
func ReadTaskStatusByProjectIDApi(w http.ResponseWriter, r *http.Request) {
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
	// Get the project id from the request
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project.ID)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Validate project id
	if project.ID == 0 {
		log.Printf("Error validating project id: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID is required."})

		return
	}
	// Read data from the task_status table
	rows, er := db.Query(`SELECT * FROM task_status WHERE project_id = $1`, project.ID)
	if er != nil {
		log.Printf("Error reading data from task_status table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading task status."})

		return
	}
	var taskStatus []models.TaskStatus
	for rows.Next() {
		var ts models.TaskStatus
		er = rows.Scan(&ts.ID, &ts.ProjectId, &ts.UserID, &ts.StatusName, &ts.StatusDescription)
		if er != nil {
			log.Printf("Error scanning data from task_status table: %s", er)
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading task status."})

			return
		}
		taskStatus = append(taskStatus, ts)
	}
	json.NewEncoder(w).Encode(taskStatus)

}

// UpdateTaskStatusByProjectIdAPI function is used to update task status by project id in the task_status table
func UpdateTaskStatusByProjectIdAPI(w http.ResponseWriter, r *http.Request) {
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
	var taskStatus models.TaskStatus
	// Decode the request body into a task_status struct
	err := json.NewDecoder(r.Body).Decode(&taskStatus)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Validate project id and user id
	if taskStatus.ProjectId == 0 || taskStatus.UserID == 0 {
		log.Printf("Error validating project id and user id: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID and User ID are required."})

		return
	}
	// Update data in the task_status table
	_, er := db.Exec(`
		UPDATE task_status SET status_name = $1, status_description = $2
			WHERE project_id = $3 AND user_id = $4`,
		taskStatus.StatusName,
		taskStatus.StatusDescription,
		taskStatus.ProjectId,
		taskStatus.UserID,
	)
	if er != nil {
		log.Printf("Error updating data in task_status table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error updating task status."})

		return
	}

}

// DeleteTaskStatusByProjectIdAndUserIdAPI function is used to delete task status by project id and user id from the task_status table
func DeleteTaskStatusByProjectIdAndUserIdAPI(w http.ResponseWriter, r *http.Request) {
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
	var taskStatus models.TaskStatus
	// Decode the request body into a task_status struct
	err := json.NewDecoder(r.Body).Decode(&taskStatus)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})

		return
	}
	// Validate project id and user id
	if taskStatus.ProjectId == 0 || taskStatus.UserID == 0 {
		log.Printf("Project ID and User ID are required.")
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID and User ID are required."})

		return
	}
	// Delete data from the task_status table
	_, er := db.Exec(`DELETE FROM task_status WHERE project_id = $1 AND user_id = $2`, taskStatus.ProjectId, taskStatus.UserID)
	if er != nil {
		log.Printf("Error deleting data from task_status table: %s", er)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error deleting task status."})

		return
	}

}

func DeleteTaskStatusByProjectID(projectID int) {
	db := db_connection.Database
	_, er := db.Exec(`DELETE FROM task_status WHERE project_id = $1`, projectID)
	if er != nil {
		log.Printf("Error deleting data from task_status table: %s", er)
	}

}
