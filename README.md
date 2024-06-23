# Project Management System

Este repositorio contiene un sistema de gestión de proyectos completo con un frontend desarrollado en Flutter y un backend en Go. El sistema permite la creación y gestión de proyectos, tareas, usuarios y roles, así como la asignación de tareas y comentarios.

## Estructura del Proyecto

- `frontend/` - Contiene el código fuente del frontend desarrollado en Flutter.
- `backend/` - Contiene el código fuente del backend desarrollado en Go.

## Tecnologías Utilizadas

- **Frontend:** Flutter
- **Backend:** Go (Golang)
- **Base de Datos:** PostgreSQL

## Características

- Gestión de usuarios y roles.
- Creación y gestión de proyectos y tareas.
- Asignación de tareas a usuarios.
- Comentarios en tareas.
- Sistema de autenticación y sesiones.
- Configuración de prioridades y estados de las tareas.

## Requisitos Previos

- Go (v1.16 o superior)
- Flutter (v2.0 o superior)
- PostgreSQL

## Configuración del Backend

1. Clona el repositorio:

    ```bash
    git clone https://github.com/tu_usuario/tu_repositorio.git
    cd tu_repositorio/backend
    ```

2. Configura las variables de conexión a la base de datos en el archivo `/.env`:

    
   
        POSTGRESQL_HOST = "localhost"
        POSTGRESQL_PORT = 5432
        POSTGRESQL_USER = "postgres"
        POSTGRESQL_PASSWORD = "postgres_pass"
        POSTGRESQL_USER_ADMIN = "psadmin"
        POSTGRESQL_PASSWORD_ADMIN = "....."
        POSTGRESQL_DB_NAME     = "psdb"
   

3. Inicia el servidor backend:

    ```bash
    go run main.go
    ```

## Configuración del Frontend

1. Ve al directorio del frontend:

    ```bash
    cd ../frontend
    ```

2. Instala las dependencias de Flutter:

    ```bash
    flutter pub get
    ```

3. Configura la URL del backend en tu archivo de configuración de Flutter (por ejemplo, `lib/config.dart`):

    ```dart
    const String backendUrl = 'http://localhost:8080';
    ```

4. Ejecuta la aplicación Flutter:

    ```bash
    flutter run
    ```

## Base de Datos

Asegúrate de tener PostgreSQL instalado y en ejecución. Puedes crear la base de datos y los usuarios necesarios utilizando las funciones del backend:

- `_CreateDataBaseIfNotExists()`
- `_CreateUserAndPasswordIfNotExists()`

## Scripts de Inicialización

El backend incluye scripts para la creación automática de tablas y la inserción de datos iniciales. Estos scripts se ejecutan automáticamente al iniciar el servidor backend.

## Contribuir

Si deseas contribuir a este proyecto, por favor sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una nueva rama (`git checkout -b feature/nueva_caracteristica`).
3. Realiza tus cambios y haz commit (`git commit -am 'Agrega nueva característica'`).
4. Sube tus cambios (`git push origin feature/nueva_caracteristica`).
5. Abre un Pull Request.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.

## Contacto

Si tienes alguna pregunta o sugerencia, no dudes en contactarme a través de [teutones92@gmail.com].

---

¡Gracias por usar nuestro sistema de gestión de proyectos!






Auth Frontend Image
![imagen](https://github.com/teutones92/project_sync/assets/72642474/f1620888-72cd-4468-ade4-368c7d1d572d)

![imagen](https://github.com/teutones92/project_sync/assets/72642474/d1f3886c-3956-424e-be9a-6c9a86377493)

![imagen](https://github.com/teutones92/project_sync/assets/72642474/c4b0c3e7-9e75-464c-853e-e9a0c8c0f479)

![imagen](https://github.com/teutones92/project_sync/assets/72642474/f9fe6502-24cf-499b-af7b-8fddd1ed6477)

![imagen](https://github.com/teutones92/project_sync/assets/72642474/6410783e-74c0-4951-bdfe-44a7f58d3794)








