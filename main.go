/*
Entry point of application.
It initializes the database connection, set up the routes and start the server on port 8080
*/
package main

func main() {

	db := SetupDatabase()
	r := SetupRouter(db)
	r.Run(":8080")
}
