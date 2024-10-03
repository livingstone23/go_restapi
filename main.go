package main

import (
	"fmt"
	"rest-api/db"
	"rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, from Rest-Api!")

	// Initialize the database
	db.InitDB()

	server := gin.Default()

	// Register the routes
	routes.RegisterRoutes(server)


	server.Run(":8089") // localhost:8089

}


