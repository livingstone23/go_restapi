package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, from Rest-Api!")

	server := gin.Default()

	server.GET("/events", getEvents) // GET request

	server.Run(":8089") // localhost:8089

}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, from Rest-Api!"})
}
