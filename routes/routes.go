package routes

import (
	"rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/events", GetEvents) // GET request
	server.GET("/events/:id", GetEvent)

	// Agroup the Endpoiint to Add the middleware to the routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreteEvents)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)

	authenticated.POST("/evets/:id/register", registerForEvent)
	authenticated.DELETE("/evets/:id/register", cancelRegister)

	//With the agroup the code is more organized
	//server.POST("/events", middlewares.Authenticate, CreteEvents)       // POST request
	//server.PUT("/events/:id", UpdateEvent)    // PUT request
	//server.DELETE("/events/:id", DeleteEvent) // DELETE request

}
