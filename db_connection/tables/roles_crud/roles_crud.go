package roles_crud

import (
	"app/db_connection"
	"app/models"
	"log"
)

func ReadUserRole() interface{} {
	// Read data from the roles table
	rows, err := db_connection.GetDatabase().Query("SELECT * FROM roles")
	if err != nil {
		log.Printf("Error reading data from roles table: %s", err)
		return models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error reading data from roles table."}
	}
	defer rows.Close()
	var roles []models.UserRole
	for rows.Next() {
		var role models.UserRole
		err := rows.Scan(&role.RoleID, &role.RoleName)
		if err != nil {
			log.Printf("Error scanning data from roles table: %s", err)
			return models.StatusCode{StatusCode: 500, StatusCodeMessage: "Error scanning data from roles table."}
		}
		roles = append(roles, role)
	}
	db_connection.GetDatabase().Close()
	log.Println("User roles red.")
	return roles
}
