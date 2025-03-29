# TaskManagerAPI

A Scalable RESTful service built with Golang

A basic Task Manager API where user can add, update, delete and list tasks.

-   Golang: Backend logic, routing and API handling
-   CRUD operations, stored in database
-   Gorm
-   Gin: Web framework for rouint and handling requests
-   PostgreSQL database for storage
-   Basic error handling - Return appropriate HTTP status code for common errors(404, 505, etc).
-   JSON Response: Tasks are stored and returned in JSON format.
-   User accounts - tasks are linked to users

Upgrading the project:

-   JWT (JSON web Tokens) for secure authentication:
    Impelemtn login to generate JWT tokens
    Use middleware to protect API endpoints
    Store and validate tokens in HTTP headers
-   Docker and docker compose( dockerized API, package it for deployment)
-   Unit Test - Write basic tests

After JWT/docker/Unit test/Porject reconstruction

-   Implement rate limiting and logging for production-level realibility
-   RabbitMQ/Kafka
-   GraphQL
-   OAuth2 or OpenID Connect
