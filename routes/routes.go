package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes will register the routes for the server. GET, POST, PUT, PATCH, DELETE
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)          // This will get all the events
	server.GET("/events/:id", getEvent)       // This will parse the id as a parameter. e.g. /events/1
	server.POST("/events", createEvent)       // POST request to create an event
	server.PUT("/events/:id", updateEvent)    // PUT request to update an event
	server.DELETE("/events/:id", deleteEvent) // DELETE request to delete an event
	server.POST("/signup", signup)            // POST request to create a user
	server.POST("/login", login)              // POST request to login a user
}
