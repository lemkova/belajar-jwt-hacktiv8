package main

import (
	"jwt-h8/database"
	"jwt-h8/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
