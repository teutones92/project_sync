package project_crud

import (
	"app/db_connection"
	"app/db_connection/tables/task_crud"
	"app/db_connection/tables/task_status_crud"
	"app/db_connection/tables/user_crud"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

func validateProject(project models.Project) *models.StatusCode {
	var status_code models.StatusCode
	if project.ProjectName == "" {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Project name is required"}
	}
	if project.Description == "" {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Description is required"}
	}
	if project.ProjectLeadID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Project lead ID is required"}
	}
	return &status_code
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
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
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}
	// Decode the request body into a project struct
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Validate the project struct
	status := validateProject(project)
	if status != nil {
		json.NewEncoder(w).Encode(status)
		log.Println(err)
		return
	}
	// Insert the project struct into the database
	db := db_connection.Database
	_, err = db.Exec("INSERT INTO projects "+
		"(project_name, description, start_date, "+
		"end_date, project_lead_id,"+
		"image_path) VALUES ($1, $2, $3, $4, $5, $6)",
		project.ProjectName,
		project.Description,
		project.StartDate,
		project.EndDate,
		project.ProjectLeadID,
		project.ImagePath)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)

		return
	}
	// Read User by ID
	user_crud.ReadUserByID(project.ProjectLeadID)
	// Create a task status for the project and User Id
	task_status_crud.DeleteTaskStatusByProjectID(project.ProjectLeadID)
	status_code = models.StatusCode{StatusCode: http.StatusCreated, StatusCodeMessage: "Project created successfully"}
	json.NewEncoder(w).Encode(status_code)
}

// Function to read all projects from the database
func ReadProjects(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request method is GET
	if !utils.CheckMethod(w, r, "GET") {
		return
	}
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}
	// Read the project from the database
	db := db_connection.Database
	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)

		return
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.ID,
			&project.ProjectName,
			&project.Description,
			&project.StartDate,
			&project.EndDate,
			&project.ProjectLeadID,
			&project.ImagePath)
		if err != nil {
			status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
			json.NewEncoder(w).Encode(status_code)
			log.Println(err)

			return
		}
		projects = append(projects, project)
	}
	json.NewEncoder(w).Encode(projects)

}

// Function to read a project by ID from the database
func ReadProjectByID(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	// Set the response header
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
	// Get the project ID from the URL
	projectID := models.Project{}.ID
	err := json.NewDecoder(r.Body).Decode(&projectID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	if projectID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Project ID is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Project ID is required")
		return
	}
	// Read the project from the database
	db := db_connection.Database
	rows, err := db.Query("SELECT * FROM projects WHERE project_id=$1", projectID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)

		return
	}
	defer rows.Close()
	var project models.Project
	for rows.Next() {
		err := rows.Scan(
			&project.ID,
			&project.ProjectName,
			&project.Description,
			&project.StartDate,
			&project.EndDate,
			&project.ProjectLeadID,
			&project.ImagePath)
		if err != nil {
			status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
			json.NewEncoder(w).Encode(status_code)
			log.Println(err)

			return
		}
	}
	json.NewEncoder(w).Encode(project)

}

// Function to read a project by project lead ID from the database
//
//	Parameters: project_lead_id as Json in the request body
func ReadProjectByProjectLeadID(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request method is GET
	if !utils.CheckMethod(w, r, "GET") {
		return
	}
	// Verify the token
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}
	// Get the project lead ID from the URL
	projectLeadID := models.Project{}.ProjectLeadID
	err := json.NewDecoder(r.Body).Decode(&projectLeadID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Check if the project lead ID is empty
	if projectLeadID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Project lead ID is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Project lead ID is required")
		return
	}
	// Read the project from the database
	db := db_connection.Database
	rows, err := db.Query("SELECT * FROM projects WHERE project_lead_id=$1", projectLeadID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)

		return
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.ID,
			&project.ProjectName,
			&project.Description,
			&project.StartDate,
			&project.EndDate,
			&project.ProjectLeadID,
			&project.ImagePath)
		if err != nil {
			status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
			json.NewEncoder(w).Encode(status_code)
			log.Println(err)

			return
		}
		projects = append(projects, project)
	}
	json.NewEncoder(w).Encode(projects)

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
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
	// Verify the token
	verify := utils.VerifyToken(w, r)
	if !verify {
		return
	}
	// Decode the request body into a project struct
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		return
	}
	// Validate the project struct
	status := validateProject(project)
	if status != nil {
		json.NewEncoder(w).Encode(status)
		log.Println(err)
		return
	}
	// Update the project struct in the database
	db := db_connection.Database
	_, err = db.Exec("UPDATE projects SET "+
		"project_name=$1, description=$2, start_date=$3, "+
		"end_date=$4, project_lead_id=$5, image_path=$6 "+
		"WHERE project_id=$7",
		project.ProjectName,
		project.Description,
		project.StartDate,
		project.EndDate,
		project.ProjectLeadID,
		project.ImagePath,
		project.ID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)

		return
	}
	status_code = models.StatusCode{StatusCode: http.StatusOK, StatusCodeMessage: "Project updated successfully"}
	json.NewEncoder(w).Encode(status_code)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	// Set the response header
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
	// Get the project ID from the URL
	var project models.Project
	e := json.NewDecoder(r.Body).Decode(&project)
	if e != nil {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: e.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(e)
		return
	}
	if project.ID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Project ID is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Project ID is required")
		return
	}
	// Delete the project from the database
	db := db_connection.Database
	_, err := db.Exec("DELETE FROM projects WHERE project_id=$1", project.ID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)

		return
	}
	task_crud.DeleteTaskByProjectID(project.ID)
	task_status_crud.DeleteTaskStatusByProjectID(project.ID)
	status_code = models.StatusCode{StatusCode: http.StatusOK, StatusCodeMessage: "Project deleted successfully"}
	json.NewEncoder(w).Encode(status_code)
}
