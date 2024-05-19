package team_member_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

// CreateTeamMemberAPI creates a new team member
func CreateTeamMemberAPI(w http.ResponseWriter, r *http.Request) {
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
	// Decode the request body into a team member struct
	var team_member_data models.TeamMember
	err := json.NewDecoder(r.Body).Decode(&team_member_data)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		return
	}
	if team_member_data.ProjectID == 0 && team_member_data.RoleID == 0 && team_member_data.UserID == 0 {
		log.Println("Project ID, Role ID and User ID not provided.")
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID, Role ID and User ID required."})
		return
	}
	// Get the database connection
	db := db_connection.Database
	// Insert data into the team_members table
	_, err = db.Exec(`
		INSERT INTO team_members (
			project_id,
			user_id,
			role_id
		) VALUES ($1, $2, $3)`,
		team_member_data.ProjectID,
		team_member_data.UserID,
		team_member_data.RoleID,
	)
	if err != nil {
		log.Printf("Error inserting data into team_members table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error inserting data into team_members table."})

		return
	}

	log.Println("Team member created.")
	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Team member created."})
}

// ReadTeamMemberAPI reads team members
func ReadTeamMembersAPI(w http.ResponseWriter, r *http.Request) {
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
	// Decode the request body into a team member struct
	var team_member_data models.TeamMember
	err := json.NewDecoder(r.Body).Decode(&team_member_data)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		return
	}
	// Check if the project ID, role ID and user ID are provided
	if team_member_data.ProjectID == 0 && team_member_data.RoleID == 0 && team_member_data.UserID == 0 {
		log.Println("Project ID, Role ID and User ID not provided.")
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID, Role ID and User ID required."})
		return
	}
	// Get the database connection
	db := db_connection.Database
	rows, err := db.Query("SELECT * FROM team_members WHERE project_id = $1 AND user_id = $2 AND role_id = $3", team_member_data.ProjectID, team_member_data.UserID, team_member_data.RoleID)
	if err != nil {
		log.Printf("Error reading data from team_members table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading data from team_members table."})

		return
	}
	var team_members []models.TeamMember
	for rows.Next() {
		var team_member models.TeamMember
		err := rows.Scan(&team_member.ID, &team_member.ProjectID, &team_member.UserID, &team_member.RoleID)
		if err != nil {
			log.Printf("Error scanning data from team_members table: %s", err)
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error scanning data from team_members table."})

			return
		}
		team_members = append(team_members, team_member)
	}

	log.Println("Team members read successfully.")
	json.NewEncoder(w).Encode(team_members)
}

// UpdateTeamMemberAPI updates a team member in the database
func UpdateTeamMemberAPI(w http.ResponseWriter, r *http.Request) {
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
	// Decode the request body into a team member struct
	var team_member_data models.TeamMember
	err := json.NewDecoder(r.Body).Decode(&team_member_data)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		return
	}
	// Check if the project ID, role ID and user ID are provided
	if team_member_data.ProjectID == 0 && team_member_data.RoleID == 0 && team_member_data.UserID == 0 {
		log.Println("Project ID, Role ID and User ID not provided.")
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID, Role ID and User ID required."})
		return
	}
	// Get the database connection
	db := db_connection.Database
	// Update the team member
	_, err = db.Exec(`
		UPDATE team_members SET role_id = $1 WHERE project_id = $2 AND user_id = $3`,
		team_member_data.RoleID,
		team_member_data.ProjectID,
		team_member_data.UserID,
	)
	if err != nil {
		log.Printf("Error updating data in team_members table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error updating data in team_members table."})

		return
	}

	log.Println("Team member updated.")
	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Team member updated."})
}

// DeleteTeamMemberAPI deletes a team member from the database
func DeleteTeamMemberAPI(w http.ResponseWriter, r *http.Request) {
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
	// Decode the request body into a team member struct
	var team_member_data models.TeamMember
	err := json.NewDecoder(r.Body).Decode(&team_member_data)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error decoding request body."})
		return
	}
	if team_member_data.ProjectID == 0 && team_member_data.RoleID == 0 && team_member_data.UserID == 0 {
		log.Println("Project ID, Role ID and User ID not provided.")
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Project ID, Role ID and User ID required."})
		return
	}
	// Get the database connection
	db := db_connection.Database
	// Delete the team member
	_, err = db.Exec("DELETE FROM team_members WHERE project_id = $1 AND user_id = $2 AND role_id = $3", team_member_data.ProjectID, team_member_data.UserID, team_member_data.RoleID)
	if err != nil {
		log.Printf("Error deleting data from team_members table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error deleting data from team_members table."})

		return
	}

	log.Println("Team member deleted.")
	json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 200, StatusCodeMessage: "Team member deleted."})
}
