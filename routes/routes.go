package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes will register the routes for the server. GET, POST, PUT, PATCH, DELETE
// Using the gin framework, we can define the routes for the server and the handlers for each route.
// The handlers are defined in separate files, e.g. routes/events.go, routes/users.go
// In Gin you can register multiple handlers for a single route.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // This will get all the events
	server.GET("/events/:id", getEvent) // This will parse the id as a parameter. e.g. /events/1

	// Group of routes that are authenticated using the Authenticate middleware function.
	authenticated := server.Group("/") // This will create a group of routes that are authenticated
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)       // POST request to create an event
	authenticated.PUT("/events/:id", updateEvent)    // PUT request to update an event
	authenticated.DELETE("/events/:id", deleteEvent) // DELETE request to delete an event

	server.POST("/signup", signup) // POST request to create a user
	server.POST("/login", login)   // POST request to login a user
}
