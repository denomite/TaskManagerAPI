# TaskManagerAPI

A Scalable RESTful Task Manager API built with Golang

This is a simple and scalable Task Manager API that allows users to create, read, update, delete (CRUD) tasks, and link them to user accounts. It implements secure authentication using JSON Web Tokens (JWT) and stores data in a PostgreSQL database.

## Tech Stack:

Golang: Backend logic, routing, and API handling.

Gin: Web framework for routing and handling requests.

GORM: ORM (Object-Relational Mapping) to interact with the database.

PostgreSQL: Database for persistent storage of tasks and user data.

JWT (JSON Web Tokens): Secure authentication system using JWT for user login and API endpoint protection.

## Features:

-   CRUD Operations:

Create, Read, Update, and Delete tasks.

Tasks are stored in a PostgreSQL database.

-   User Accounts:

Tasks are linked to user accounts, enabling personalized task management.

-   Authentication and Authorization:

JWT: Implemented for secure login and token-based authentication.

Middleware: Protects API endpoints by verifying JWT tokens passed in HTTP headers.

-   Error Handling:

Returns appropriate HTTP status codes for common errors such as 404 (Not Found), 400 (Bad Request), and 500 (Internal Server Error).

JSON response format for consistent API communication.

-   Task Management API:

API allows users to interact with their tasks through routes:

POST /tasks – Create a new task.

GET /tasks – Get a list of all tasks.

GET /tasks/:id – Get a specific task by ID.

PUT /tasks/:id – Update an existing task.

DELETE /tasks/:id – Delete a task.

## Project structure

/TaskManagerAPI  
├── /controllers # API logic and handler functions  
├── /models # Task and User models  
├── /repository # Database interaction (CRUD operations)  
├── /routes # Route definitions and middleware  
├── /utils # Helper functions (e.g., for JWT, password hashing)  
├── /config # Configuration files (e.g., DB, environment variables)  
├── main.go # Entry point to run the application  
└── go.mod # Go module dependencies

### In the near future

-   Docker and dockerized API
-   Implement rate limiting and logging for production-level realibility and protect against
    exessive api requests
-   CI/CD integraton
-   Unit test

-   RabbitMQ/Kafka
-   GraphQL
-   OAuth2 or OpenID Connect
