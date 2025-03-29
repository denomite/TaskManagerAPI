/*
Entry point of application.
It initializes the database connection, set up the routes and start the server on port 8080
*/
package main

import (
	"TaskManagerAPI/repository"
	"TaskManagerAPI/routes"
)

func main() {

	db := repository.SetupDatabase()
	r := routes.SetupRouter(db)
	r.Run(":8080")
}
