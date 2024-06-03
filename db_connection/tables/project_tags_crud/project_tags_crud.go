package project_tags_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CreateProjectTagAPI creates a project tag
func CreateProjectTagAPI(w http.ResponseWriter, r *http.Request) {
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
	var project_tag models.ProjectTag
	// Decode the request body into a project tag struct
	err := json.NewDecoder(r.Body).Decode(&project_tag)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body. Must provide project tag data. as JSON object {\"project_id\": 1, \"tag_name\": \"tag_name\"}"})
		return
	}
	// Create a project tag
	status := createProjectTag(project_tag)
	// Return a response to the client indicating that the project tag was created
	json.NewEncoder(w).Encode(status)
}

// createProjectTag creates a project tag
func createProjectTag(project_tag_data models.ProjectTag) models.StatusCode {
	// Insert data into the project_tags table
	_, err := db_connection.Database.Exec(`
		INSERT INTO project_tags (project_id, tag_name) VALUES ($1, $2)`,
		project_tag_data.ProjectID, project_tag_data.TagName)
	if err != nil {
		panic(fmt.Sprintf("Error inserting data into project_tags table: project_tags %s", err))
	}
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Project tag created."}

}

// ReadProjectTagAPI reads a project tag
func ReadProjectTagAPI(w http.ResponseWriter, r *http.Request) {
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
	var project_tag models.ProjectTag
	// Decode the request body into a project tag struct
	err := json.NewDecoder(r.Body).Decode(&project_tag)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		return
	}
	if project_tag.ID == 0 {
		// Get all default project tags
		project_tags, err := readProjectTag(0, nil)
		if err != nil {
			log.Printf("Error getting project tag data: %v", err)
			// Return a response to the client indicating that there was an error getting the project tag data
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error getting project tag data. According to the project tag ID provided."})
			return
		}
		// Return a response to the client with all default project tags
		json.NewEncoder(w).Encode(project_tags)
		return
	}
	if project_tag.ProjectID != nil {
		// Get a project tag
		project_tags, er := readProjectTag(0, project_tag.ProjectID)
		if er != nil {
			log.Printf("Error getting project tag data: %v", er)
			// Return a response to the client indicating that there was an error getting the project tag data
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error getting project tag data. According to the project ID provided."})
			return
		}
		// Return a response to the client with the project tag data
		json.NewEncoder(w).Encode(project_tags)
		return
	}
	// Query the database to get the project tag data
	project_tags, er := readProjectTag(project_tag.ID, nil)
	if er != nil {
		log.Printf("Error getting project tag data: %v", er)
		// Return a response to the client indicating that there was an error getting the project tag data
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error getting project tag data."})
		return
	}
	json.NewEncoder(w).Encode(project_tags)
}

// readProjectTagByID reads a project tag by ID
func readProjectTag(tag_id int, project_id *int) (*[]models.ProjectTag, error) {
	db := db_connection.Database
	// Query the database to get the project tag data
	var project_tag models.ProjectTag
	var project_tags []models.ProjectTag
	if tag_id == 0 {
		rows, err := db.Query("SELECT * FROM project_tags")
		if err != nil {
			return nil, err
		}
		// Get the project tag data
		for rows.Next() {
			err = rows.Scan(&project_tag.ID, &project_tag.ProjectID, &project_tag.TagName)
			if err != nil {
				return nil, err
			}
			project_tags = append(project_tags, project_tag)
		}
		defer rows.Close()
	} else if project_id != nil {
		rows, err := db.Query("SELECT * FROM project_tags WHERE project_id=$1", project_id)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err = rows.Scan(&project_tag.ID, &project_tag.ProjectID, &project_tag.TagName)
			if err != nil {
				return nil, err
			}
			project_tags = append(project_tags, project_tag)
		}
		defer rows.Close()
	} else {
		rows, err := db.Query("SELECT * FROM project_tags WHERE id=$1, project_id=$2", tag_id, project_id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		// Get the project tag data
		for rows.Next() {
			err = rows.Scan(&project_tag.ID, &project_tag.ProjectID, &project_tag.TagName)
			if err != nil {
				return nil, err
			}
			project_tags = append(project_tags, project_tag)
		}
	}
	return &project_tags, nil
}

// UpdateProjectTagAPI updates a project tag

// func UpdateProjectTagAPI(w http.ResponseWriter, r *http.Request) {
// 	// Set the response header
// 	utils.SetHeader(w)
// 	// Check the request method
// 	method := utils.CheckMethod(w, r, "PUT")
// 	if !method {
// 		return
// 	}
// 	// Check the header
// 	header := utils.CheckHeader(w, r)
// 	if !header {
// 		return
// 	}
// 	// Verify the token
// 	verify := utils.VerifyToken(w, r)
// 	if !verify {
// 		return
// 	}
// 	var project_tag models.ProjectTag
// 	// Decode the request body into a project tag struct
// 	err := json.NewDecoder(r.Body).Decode(&project_tag)
// 	if err != nil {
// 		log.Printf("Error decoding request body: %v", err)
// 		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body. Must provide project tag data. as JSON object {\"id\": 1, \"project_id\": 1, \"tag_name\": \"tag_name\"}"})
// 		return
// 	}
// 	// Update a project tag
// 	status := updateProjectTag(project_tag)
// 	// Return a response to the client indicating that the project tag was updated
// 	json.NewEncoder(w).Encode(status)
// }

// // updateProjectTag updates a project tag
// func updateProjectTag(project_tag_data models.ProjectTag) models.StatusCode {
// 	// Update data in the project_tags table
// 	_, err := db_connection.Database.Exec(`
// 		UPDATE project_tags SET project_id=$1, tag_name=$2 WHERE id=$3`,
// 		project_tag_data.ProjectID, project_tag_data.TagName, project_tag_data.ID)
// 	if err != nil {
// 		panic(fmt.Sprintf("Error updating data in project_tags table: project_tags %s", err))
// 	}
// 	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Project tag updated."}
// }

// DeleteProjectTagAPI deletes a project tag
func DeleteProjectTagAPI(w http.ResponseWriter, r *http.Request) {
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
	var project_tag models.ProjectTag
	// Decode the request body into a project tag struct
	err := json.NewDecoder(r.Body).Decode(&project_tag)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body. Must provide project tag ID. as JSON object {\"id\": 1}"})
		return
	}
	// Delete a project tag
	status := deleteProjectTag(project_tag.ID)
	// Return a response to the client indicating that the project tag was deleted
	json.NewEncoder(w).Encode(status)
}

// deleteProjectTag deletes a project tag
func deleteProjectTag(project_tag_id int) models.StatusCode {
	// Delete data from the project_tags table
	_, err := db_connection.Database.Exec("DELETE FROM project_tags WHERE id=$1", project_tag_id)
	if err != nil {
		panic(fmt.Sprintf("Error deleting data from project_tags table: project_tags %s", err))
	}
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Project tag deleted."}
}
