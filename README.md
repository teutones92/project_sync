# Project Management System

This repository contains a comprehensive project management system with a frontend developed in Flutter and a backend in Go. The system allows the creation and management of projects, tasks, users, and roles, as well as task assignments and comments.

## Project Structure

### Branches
- `frontend/` - Contains the source code of the frontend developed in Flutter.
- `backend/` - Contains the source code of the backend developed in Go.

## Technologies Used

- **Frontend:** Flutter
- **Backend:** Go (Golang)
- **Database:** PostgreSQL

## Features

- User and role management.
- Project and task creation and management.
- Task assignment to users.
- Task comments.
- Authentication and session system.
- Task priority and status settings.

## Prerequisites

- Go (v1.16 or higher)
- Flutter (v2.0 or higher)
- PostgreSQL

## Backend Configuration

1. Clone the repository:

    ```bash
    git clone (https://github.com/teutones92/project_sync.git) brand frontend and backend
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

## Frontend Configuration

1. Go to the frontend directory:

    ```bash
    cd ../frontend
    ```

2. Install Flutter dependencies:

    ```bash
    flutter pub get
    ```

3. Configure the backend URL in your Flutter configuration file (e.g., `lib/config.dart`):

    ```dart
    const String backendUrl = 'http://localhost:8080';
    ```

4. Run the Flutter application:

    ```bash
    flutter run
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

Auth Frontend Image
![image](https://github.com/teutones92/project_sync/assets/72642474/f1620888-72cd-4468-ade4-368c7d1d572d)

![image](https://github.com/teutones92/project_sync/assets/72642474/d1f3886c-3956-424e-be9a-6c9a86377493)

![image](https://github.com/teutones92/project_sync/assets/72642474/c4b0c3e7-9e75-464c-853e-e9a0c8c0f479)

![image](https://github.com/teutones92/project_sync/assets/72642474/f9fe6502-24cf-499b-af7b-8fddd1ed6477)

![image](https://github.com/teutones92/project_sync/assets/72642474/6410783e-74c0-4951-bdfe-44a7f58d3794)
