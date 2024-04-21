package project_crud

import (
	"app/db_connection"
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
	// Decode the request body into a project struct
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
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
	db := db_connection.GetDatabase()
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
		db.Close()
		return
	}
	status_code = models.StatusCode{StatusCode: http.StatusCreated, StatusCodeMessage: "Project created successfully"}
	json.NewEncoder(w).Encode(status_code)
}

func ReadProjects(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request method is GET
	if !utils.CheckMethod(w, r, "GET") {
		return
	}
	// Read the project from the database
	db := db_connection.GetDatabase()
	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		db.Close()
		return
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(
			&project.ProjectID,
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
			db.Close()
			return
		}
		projects = append(projects, project)
	}
	json.NewEncoder(w).Encode(projects)
	db.Close()
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
	db := db_connection.GetDatabase()
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
		project.ProjectID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		db.Close()
		return
	}
	status_code = models.StatusCode{StatusCode: http.StatusOK, StatusCodeMessage: "Project updated successfully"}
	json.NewEncoder(w).Encode(status_code)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	var status_code models.StatusCode
	// Set the response header
	utils.SetHeader(w)
	// Check if the request method is DELETE
	if !utils.CheckMethod(w, r, "DELETE") {
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
	if project.ProjectID == 0 {
		status_code = models.StatusCode{StatusCode: http.StatusBadRequest, StatusCodeMessage: "Project ID is required"}
		json.NewEncoder(w).Encode(status_code)
		log.Println("Project ID is required")
		return
	}
	// Delete the project from the database
	db := db_connection.GetDatabase()
	_, err := db.Exec("DELETE FROM projects WHERE project_id=$1", project.ProjectID)
	if err != nil {
		status_code = models.StatusCode{StatusCode: http.StatusInternalServerError, StatusCodeMessage: err.Error()}
		json.NewEncoder(w).Encode(status_code)
		log.Println(err)
		db.Close()
		return
	}
	status_code = models.StatusCode{StatusCode: http.StatusOK, StatusCodeMessage: "Project deleted successfully"}
	json.NewEncoder(w).Encode(status_code)
}
