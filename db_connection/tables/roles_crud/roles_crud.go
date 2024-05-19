package roles_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

func ReadUserRoleAPI(w http.ResponseWriter, r *http.Request) {
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
	rows, err := db.Query("SELECT * FROM roles")
	if err != nil {
		log.Printf("Error reading data from roles table: %s", err)
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading data from roles table."})
	}
	defer rows.Close()
	var roles []models.UserRole
	for rows.Next() {
		var role models.UserRole
		err := rows.Scan(&role.ID, &role.RoleName)
		if err != nil {
			log.Printf("Error scanning data from roles table: %s", err)
			json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error scanning data from roles table."})

		}
		roles = append(roles, role)
	}
	db_connection.Database.Close()
	log.Println("User roles read successfully.")
	json.NewEncoder(w).Encode(roles)
}
