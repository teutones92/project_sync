# Project Management System - Backend

This directory contains the backend of the Project Management System, developed in Go. The backend handles user authentication, project and task management, and interactions with the PostgreSQL database.

## Technologies Used

- **Backend:** Go (Golang)
- **Database:** PostgreSQL

## Features

- User and role management.
- Project and task creation and management.
- Task assignment to users.
- Task comments.
- Authentication and session management.
- Task priority and status settings.

## Prerequisites

- Go (v1.16 or higher)
- PostgreSQL

## Configuration

1. Clone the repository:

    ```bash
    git clone https://github.com/your_username/your_repository.git
    cd your_repository/backend
    ```

2. Configure the database connection variables in the `/.env` file:

    ```plaintext
    POSTGRESQL_HOST = "localhost"
    POSTGRESQL_PORT = 5432
    POSTGRESQL_USER = "postgres"
    POSTGRESQL_PASSWORD = "postgres_pass"
    POSTGRESQL_USER_ADMIN = "psadmin"
    POSTGRESQL_PASSWORD_ADMIN = "....."
    POSTGRESQL_DB_NAME = "psdb"
    ```

3. Start the backend server:

    ```bash
    go run main.go
    ```

## Database

Make sure you have PostgreSQL installed and running. You can create the necessary database and users using the backend functions:

- `_CreateDataBaseIfNotExists()`
- `_CreateUserAndPasswordIfNotExists()`

## Initialization Scripts

The backend includes scripts for automatic table creation and initial data insertion. These scripts run automatically when the backend server starts.

## Contributing

If you want to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/new_feature`).
3. Make your changes and commit (`git commit -am 'Add new feature'`).
4. Push your changes (`git push origin feature/new_feature`).
5. Open a Pull Request.

## License

This project is under the MIT License. See the `LICENSE` file for more details.

## Contact

If you have any questions or suggestions, feel free to contact me at [teutones92@gmail.com].

---

Thank you for using our project management system!
