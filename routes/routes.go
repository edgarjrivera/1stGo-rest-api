package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes will register the routes for the server. GET, POST, PUT, PATCH, DELETE
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // This will get all the events
	server.GET("/events/:id", getEvent) // This will parse the id as a parameter. e.g. /events/1
	server.POST("/events", createEvent) // POST request to create an event
}
