package main

import (
	"project-2/app"
	"project-2/database"
)

func main() {
	database.StartDB()
	app.StartServer()
}
