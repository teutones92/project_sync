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

func Init() {
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
		_CreateTables()
	}
	// Close the database connection
	defer database.Close()
}

func _CreateTables() {
	// Define a map containing SQL queries to create the tables
	createQueries := map[string]string{
		"users": `
            CREATE TABLE IF NOT EXISTS users (
                user_id SERIAL PRIMARY KEY,
                username VARCHAR(50) NOT NULL,
                email VARCHAR(100) NOT NULL UNIQUE,
                password_hash VARCHAR(100) NOT NULL,
				user_avatar_path TEXT DEFAULT ''
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
                project_lead_id INT,
				image_path TEXT
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
				image_path TEXT[],
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
                FOREIGN KEY (user_id) REFERENCES users(user_id)
               -- FOREIGN KEY (role_id) REFERENCES user_roles(role_id)
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
				token VARCHAR(100) UNIQUE,
                last_activity_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    			expiration_minutes INT DEFAULT 30,
                FOREIGN KEY (user_id) REFERENCES users(user_id)
            );`,
	}
	tableOrder := []string{"users", "user_roles", "projects", "tasks", "team_members", "comments", "sessions"}
	// Iterate over the map and execute each SQL query
	tableCreated := make(chan bool)
	go func() {
		for _, tableName := range tableOrder {
			query, ok := createQueries[tableName]
			if !ok {
				log.Printf("Table %s not found in createQueries", tableName)
				continue
			}
			_, err := database.Exec(query)
			if err != nil {
				// panic(fmt.Sprintf("Error creating table %s: %s", tableName, err))
				log.Printf("Error creating table %s: %s", tableName, err)
			}
			log.Printf("Table %s has been created successfully.", tableName)
		}
		tableCreated <- true
	}()
	log.Println("waiting for table creation...")
	<-tableCreated
	// log.Println("Tables have been created successfully.")
	// Insert data into the user_roles table
	insertUserRoles(database)
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
