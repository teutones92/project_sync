package db_connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver (required for database/sql)
)

var database *sql.DB

const (
	// Host and port to connect to the PostgreSQL server
	postgresqlHost = "localhost"
	postgresqlPort = 5432
	// Username and password to connect to the PostgreSQL server
	serverUserName = "postgres"
	serverPassword = "rfv/789*-+"
	// Username and password to connect to the database
	user     = "psadmin"
	password = "Calibre92*"
	db_name  = "psdb"
)

func GetDatabase() *sql.DB {
	// Create Connection()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresqlHost, postgresqlPort, user, password, db_name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Verify the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// Function to create a database in postgresql server if it does not exist
func _CreateDataBaseIfNotExists() bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", postgresqlHost, postgresqlPort, serverUserName, serverPassword)
	// Connect to PostgreSQL server
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error opening database connection: %s", err)
		return false
	}
	defer db.Close()

	// Check if the database already exists
	var dbExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'psdb')").Scan(&dbExists)
	if err != nil {
		log.Printf("Error checking if database exists: %s", err)
		return false
	}

	// If the database does not exist, create it
	if !dbExists {
		_, err := db.Exec("CREATE DATABASE psdb")
		if err != nil {
			log.Printf("Error creating database: %s", err)
			return false
		}
		log.Println("Database created successfully.")
	} else {
		log.Println("Database already exists.")
	}
	return true
}

// Function to create a user and password in postgresql server if they do not exist
func _CreateUserAndPasswordIfNotExists() bool {
	var username string = "postgres"
	var password string = "rfv/789*-+"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", postgresqlHost, postgresqlPort, username, password)
	// Connect to PostgreSQL server
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error opening database connection: %s", err)
		return false
	}
	defer db.Close()

	// Check if the user already exists
	var userExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname = 'psadmin')").Scan(&userExists)
	if err != nil {
		log.Printf("Error checking if user exists: %s", err)
		return false
	}

	// If the user does not exist, create it
	if !userExists {
		_, err := db.Exec("CREATE ROLE psadmin WITH LOGIN PASSWORD 'Calibre92*' SUPERUSER CREATEDB CREATEROLE;")
		if err != nil {
			log.Printf("Error creating user: %s", err)
			return false
		}
		log.Println("User created successfully.")
	} else {
		log.Println("User already exists.")
	}
	return true
}

func Init() error {
	// Create channels for synchronization
	var dbCreated bool
	var userCreated bool
	// dbCreated := make(chan bool)
	// userCreated := make(chan bool)
	// databaseConnected := make(chan bool)
	// Create the database if it does not exist (goroutine)
	respDb := _CreateDataBaseIfNotExists()
	dbCreated = respDb
	// Create a user and password if they do not exist (goroutine)
	respUser := _CreateUserAndPasswordIfNotExists()
	userCreated = respUser
	// Database connection
	database = GetDatabase()
	// Create tables in the database
	if dbCreated && userCreated {
		terror := _CreateTables()
		if terror != nil {
			return terror
		}
	}
	// Close the database connection
	defer database.Close()
	return nil
}

func _CreateTables() error {
	// Define a map containing SQL queries to create the tables
	createQueries := map[string]string{
		"users": `
            CREATE TABLE IF NOT EXISTS users (
                id SERIAL PRIMARY KEY,
                username VARCHAR(50) NOT NULL,
                email VARCHAR(100) NOT NULL UNIQUE,
                password_hash VARCHAR(100) NOT NULL,
				user_avatar_path TEXT
				-- user_avatar_path TEXT DEFAULT ''
            );`,
		"user_roles": `
            CREATE TABLE IF NOT EXISTS user_roles (
                id SERIAL PRIMARY KEY,
                role_name VARCHAR(50) NOT NULL UNIQUE,
                role_description TEXT
            );`,
		"projects": `
            CREATE TABLE IF NOT EXISTS projects (
                id SERIAL PRIMARY KEY,
                project_name VARCHAR(100) NOT NULL,
                description TEXT,
                start_date DATE,
                end_date DATE,
                project_lead_id INT,
				image_path TEXT
            );`,
		"task_status": `
            CREATE TABLE IF NOT EXISTS task_status (
                id SERIAL PRIMARY KEY,
				project_id INT,
                user_id INT,
                status_name TEXT,
                status_description TEXT,
                FOREIGN KEY (project_id) REFERENCES projects(id),
                FOREIGN KEY (user_id) REFERENCES users(id)
            );`,
		"tasks": `
            CREATE TABLE IF NOT EXISTS tasks (
                id SERIAL PRIMARY KEY,
                project_id INT,
                task_name VARCHAR(100) NOT NULL,
                description TEXT,
                status_id INT,
                priority VARCHAR(20),
                assigned_user INT,
                deadline DATE,
				image_path TEXT[],
                FOREIGN KEY (project_id) REFERENCES projects(id),
                FOREIGN KEY (assigned_user) REFERENCES users(id),
                FOREIGN KEY (status_id) REFERENCES task_status(id)
            );`,
		"team_members": `
            CREATE TABLE IF NOT EXISTS team_members (
                id SERIAL PRIMARY KEY,
                project_id INT,
                user_id INT,
                role_id INT,
                FOREIGN KEY (project_id) REFERENCES projects(id),
                FOREIGN KEY (user_id) REFERENCES users(id)
               -- FOREIGN KEY (role_id) REFERENCES user_roles(id)
            );`,
		"comments": `
            CREATE TABLE IF NOT EXISTS comments (
                id SERIAL PRIMARY KEY,
                project_id INT,
                task_id INT,
                user_id INT,
                timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                comment_text TEXT,
                FOREIGN KEY (project_id) REFERENCES projects(id),
                FOREIGN KEY (task_id) REFERENCES tasks(id),
                FOREIGN KEY (user_id) REFERENCES users(id)
            );`,
		"sessions": `
            CREATE TABLE IF NOT EXISTS sessions (
                id SERIAL PRIMARY KEY,
                user_id INT,
				token VARCHAR(100) UNIQUE,
                last_activity_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    			expiration_minutes INT DEFAULT 30,
                FOREIGN KEY (user_id) REFERENCES users(id)
            );`,
		"user_contacts": `
			CREATE TABLE IF NOT EXISTS user_contacts (
				id SERIAL PRIMARY KEY,
				user_id INT,
				contact_name VARCHAR(100) NOT NULL,
				contact_email VARCHAR(100) NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id)
			);`,
	}
	tableOrder := []string{"users", "user_roles", "projects", "task_status", "tasks", "team_members", "comments", "sessions", "user_contacts"}
	// Iterate over the map and execute each SQL query
	done := make(chan bool)
	log.Println("waiting for table creation...")
	for _, tableName := range tableOrder {
		_TableExists(tableName, database)
		query, ok := createQueries[tableName]
		go func(tName string, tQuery string) {
			_, err := database.Exec(tQuery)
			if !ok {
				log.Printf("Table %s not found in createQueries", tName)
			}
			if err != nil {
				log.Fatalf("Error creating table %s: %s", tName, err)
				return
			}
			done <- true
		}(tableName, query)
		<-done
		// log.Printf("Table %s has been created successfully.", tableName)
	}
	log.Println("Tables have been created successfully.")
	// Insert data into the user_roles table
	insertUserRoles(database)
	return nil
}

func _TableExists(tableName string, db *sql.DB) {
	var exists string
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", tableName)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		panic(fmt.Sprintf("Error checking if table exists: %s", err))
	}
	if exists == "true" {
		log.Printf("Using %s table.", tableName)
	} else {
		log.Printf("Table %s has been created successfully.", tableName)
	}
}

func insertUserRoles(db *sql.DB) {
	log.Println("Inserting data into user_roles table...")
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_roles").Scan(&count)
	if err != nil {
		panic(fmt.Sprintf("Error counting rows in user_roles table: %s", err))
	}
	// If count is greater than zero, print a message indicating that the data already exists in the user_roles table
	if count > 0 {
		log.Println("Data already exists in user_roles table. Skipping insertion.")
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
			log.Printf("Error inserting role information into user_roles table: %s", err)
		}
		log.Println("Information has been inserted into the user_roles table.")
	}
}

func InsertTaskStatus(db *sql.DB, projectID int, userID int) {
	log.Println("Inserting data into task_status table...")
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM task_status").Scan(&count)
	if err != nil {
		panic(fmt.Sprintf("Error counting rows in task_status table: %s", err))
	}
	// If count is greater than zero, print a message indicating that the data already exists in the task_status table
	if count > 0 {
		log.Println("Data already exists in task_status table. Skipping insertion.")
		db.Close()
		return
	}
	// If the count is zero, insert the data into the task_status table
	_, er := db.Exec(`
		INSERT INTO task_status (status_name, status_description, project_id, user_id) VALUES
		('To Do', 'Task has not been started yet.'),
		('In Progress', 'Task is currently being worked on.'),
		('In Review', 'Task is being reviewed by the project manager or team lead.'),
		('Done', 'Task has been completed and delivered.', $1, $2)`, projectID, userID)
	if er != nil {
		log.Printf("Error inserting status information into task_status table: %s", err)
	}
	log.Println("Information has been inserted into the task_status table.")

}
