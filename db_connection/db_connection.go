package db_connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver (required for database/sql)
)

func Init() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "psadmin"
		password = "Calibre92*"
		db_name  = "psdb"
	)
	// Create Connection()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	// Open the connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verify the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create the tables in the database
	fmt.Println("The connection to the database was successful.")
	_CreateTables(db)
}
func _CreateTables(db *sql.DB) {
	// Define a map containing SQL queries to create the tables
	createQueries := map[string]string{
		"users": `
            CREATE TABLE IF NOT EXISTS users (
                user_id SERIAL PRIMARY KEY,
                username VARCHAR(50) NOT NULL,
                email VARCHAR(100) NOT NULL UNIQUE,
                password_hash VARCHAR(100) NOT NULL
            );`,
		"user_roles": `
            CREATE TABLE IF NOT EXISTS user_roles (
                role_id SERIAL PRIMARY KEY,
                role_name VARCHAR(50) NOT NULL UNIQUE,
                role_description TEXT
            );`,
		"projects": `
            CREATE TABLE IF NOT EXISTS projects (
                project_id SERIAL PRIMARY KEY,
                project_name VARCHAR(100) NOT NULL,
                description TEXT,
                start_date DATE,
                end_date DATE,
                project_lead_id INT
            );`,
		"tasks": `
            CREATE TABLE IF NOT EXISTS tasks (
                task_id SERIAL PRIMARY KEY,
                project_id INT,
                task_name VARCHAR(100) NOT NULL,
                description TEXT,
                status VARCHAR(20),
                priority VARCHAR(20),
                assigned_user INT,
                deadline DATE,
                FOREIGN KEY (project_id) REFERENCES projects(project_id),
                FOREIGN KEY (assigned_user) REFERENCES users(user_id)
            );`,
		"team_members": `
            CREATE TABLE IF NOT EXISTS team_members (
                team_member_id SERIAL PRIMARY KEY,
                project_id INT,
                user_id INT,
                role_id INT,
                FOREIGN KEY (project_id) REFERENCES projects(project_id),
                FOREIGN KEY (user_id) REFERENCES users(user_id),
                FOREIGN KEY (role_id) REFERENCES user_roles(role_id)
            );`,
		"comments": `
            CREATE TABLE IF NOT EXISTS comments (
                comment_id SERIAL PRIMARY KEY,
                project_id INT,
                task_id INT,
                user_id INT,
                timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                comment_text TEXT,
                FOREIGN KEY (project_id) REFERENCES projects(project_id),
                FOREIGN KEY (task_id) REFERENCES tasks(task_id),
                FOREIGN KEY (user_id) REFERENCES users(user_id)
            );`,
		"sessions": `
            CREATE TABLE IF NOT EXISTS sessions (
                session_id SERIAL PRIMARY KEY,
                user_id INT,
                expiration_timestamp TIMESTAMP,
                FOREIGN KEY (user_id) REFERENCES users(user_id)
            );`,
	}

	// Iterate over the map and execute each SQL query
	for tableName, query := range createQueries {
		_, err := db.Exec(query)
		if err != nil {
			panic(fmt.Sprintf("Error creating table %s: %s", tableName, err))
		}
	}
	fmt.Println("Tables have been created successfully.")

	// Insert data into the user_roles table
	fmt.Println("Inserting data into user_roles table...")
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_roles").Scan(&count)
	if err != nil {
		panic(fmt.Sprintf("Error counting rows in user_roles table: %s", err))
	}
	// If count is greater than zero, print a message indicating that the data already exists in the user_roles table
	if count > 0 {
		fmt.Println("Data already exists in user_roles table. Skipping insertion.")
	} else {
		// If the count is zero, insert the data into the user_roles table
		_, err := db.Exec(`
        INSERT INTO user_roles (role_name, role_description) VALUES
        ('Project Lead', 'Responsible for leading and managing projects.'),
        ('Project Manager', 'Responsible for overseeing project planning, execution, and delivery.'),
        ('Team Member', 'Member of the project team responsible for user interface design and implementation.'),
        ('Designer', 'Responsible for graphic design and user interface (UI) design.'),
        ('Main Developer', 'Lead developer responsible for overall software architecture and design.'),
        ('Developer', 'Member of the development team responsible for coding and implementation.'),
        ('Junior Developer', 'Entry-level developer learning and contributing to development tasks.'),
        ('Guest', 'Limited access user with read-only permissions.')`)
		if err != nil {
			panic(fmt.Sprintf("Error inserting data into user_roles table: %s", err))
		}
		fmt.Println("Data inserted successfully into user_roles table.")
	}

}
