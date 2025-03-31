# TaskManagerAPI

A Scalable RESTful Task Manager API built with Golang

This is API allows users to manage tasks with CRUD operations, implement role-based access control (RBAC) and handle user authentication and authorization.

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

Role-Based Access Control (RBAC):

-   Regular Users: Can create and manage only their own tasks.

-   Admins: Can manage all tasks and users.

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

### In production

-   Role-Based Access Control (Admin, User)
-   Task prioritization and due dates for tasks

### In the near future

#### For Production readiness:

-   Dockerize the API
-   Implement rate limiting and logging
-   CI/CD integraton
-   Unit tests

#### Advance features

-   Integrate RabbitMQ/Kafka for asynchronous processing
-   GraphQL
-   OAuth2 or OpenID Connect
