package main

import (
	"github.com/screamingarrow/gin-rest-api/database"
	"github.com/screamingarrow/gin-rest-api/routes"
)

func main() {
	database.ConectaDb()

	routes.HandleRequests()
}
