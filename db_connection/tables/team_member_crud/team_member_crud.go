package team_member_crud

import (
	"app/db_connection"
	"app/models"
	"log"
)

func CreateTeamMember(team_member_data models.TeamMember) models.StatusCode {
	// Insert data into the team_members table
	_, err := db_connection.GetDatabase().Exec(`
		INSERT INTO team_members (role_id) VALUES ($1)`,
		team_member_data.RoleID)
	if err != nil {
		log.Printf("Error inserting data into team_members table: %s", err)
		return models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error inserting data into team_members table."}
	}
	db_connection.GetDatabase().Close()
	log.Println("Team member created.")
	return models.StatusCode{StatusCode: 200, StatusCodeMessage: "Team member created."}
}
