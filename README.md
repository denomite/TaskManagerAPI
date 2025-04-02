<br/>

# TaskManagerAPI

<br/>

A Scalable RESTful Task Manager API built with Golang for managing tasks, implementing Role-Based Access Control (RBAC), and handling user authentication and authorization.

## 🚀 Tech Stack:

<br/>

-   Golang: Backend logic, routing, and API handling.

-   Gin: Web framework for routing and handling requests.

-   GORM: ORM (Object-Relational Mapping) to interact with the PostgreSQL database.

-   PostgreSQL: Relational database for persistent storage of tasks and user data.

-   JWT (JSON Web Tokens): Secure authentication system using JWT for user login and API endpoint protection.

-   Docker: Containerized API for easy deployment.

<br/>

## ⚡ Features:

<br/>

1.  CRUD Operations:

    -   Create, Read, Update, and Delete tasks.

    -   Tasks are stored in a PostgreSQL database.

2.  User Accounts:

    -   Tasks are linked to user accounts, enabling personalized task management.

    -   Each user has their own task list, ensuring task segregation.

3.  Authentication and Authorization:

    -   JWT: Implemented for secure login and token-based authentication.

    -   Middleware: Protects API endpoints by verifying JWT tokens passed in HTTP headers.

4.  Role-Based Access Control (RBAC):

    -   Regular Users: Can create and manage only their own tasks.

    -   Admins: Can manage all tasks and users, offering full administrative capabilities.

5.  Error Handling:

    -   Returns appropriate HTTP status codes for common errors:

        -   404 (Not Found)
        -   400 (Bad Request)
        -   500 (Internal Server Error)

    -   JSON response format for consistent API communication.

6.  Task Management API:

    -   API allows users to interact with their tasks through routes:

    -   POST /tasks – Create a new task.

    -   GET /tasks – Get a list of all tasks.

    -   GET /tasks/:id – Get a specific task by ID.

    -   PUT /tasks/:id – Update an existing task.

    -   DELETE /tasks/:id – Delete a task.

<br/>

## 🔧 Project structure

<br/>

/TaskManagerAPI  
├── /controllers # API logic and handler functions  
├── /models # Task and User models  
├── /repository # Database interaction (CRUD operations)  
├── /routes # Route definitions and middleware  
├── /utils # Helper functions (e.g., for JWT, password hashing)  
├── /config # Configuration files (e.g., DB, environment variables)  
├── main.go # Entry point to run the application  
└── go.mod # Go module dependencies

<br/>

## 📦 Dockerized API

<br/>

The API is dockerized for easy deployment. You can run it using the following:

    1. Build and run the docker container

    docker-compose up --build

    2. Stop the docker container

    docker-compose down

    3. Docker compose configuration

    This repository includes a docker-compose.yml file that sets up the PostgreSQL database and the API container. You can modify the environment variables for the database   connection in the .env file.

<br/>

## 🧪 Unit Tests

<br/>

    The project includes simple unit tests for core functionality, located in the tests folder.

    To run the tests, use:

    go test ./tests
