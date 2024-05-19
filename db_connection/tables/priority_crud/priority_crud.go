package priority_crud

import (
	"app/db_connection"
	"app/models"
	"app/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ReadPriorityAPI reads a priority
func ReadPriorityAPI(w http.ResponseWriter, r *http.Request) {
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
	// Query the database to get the priority data
	var priority *models.Priority
	priority, er := readPriority()
	if er != nil {
		log.Printf("Error getting priority data: %v", er)
		// Return a response to the client indicating that there was an error getting the priority data
		json.NewEncoder(w).Encode(&models.StatusCode{StatusCode: 400, StatusCodeMessage: "Error getting priority data."})
		return
	}
	// Return the priority data to the client
	json.NewEncoder(w).Encode(priority)
}

// ReadPriorityByID reads a priority by ID
func readPriority() (*models.Priority, error) {
	db := db_connection.GetDatabase()
	// Query the database to get the priority data
	rows, err := db.Query("SELECT * FROM priority")
	if err != nil {
		print(rows)
		return nil, err
	}
	defer rows.Close()
	var priority models.Priority
	// Get the priority data
	for rows.Next() {
		err = rows.Scan(&priority.ID, &priority.PriorityName, &priority.PriorityDescription)
		if err != nil {
			return nil, err
		}
	}
	return &priority, nil
}
