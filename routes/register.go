package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/operations"
	"github.com/gin-gonic/gin"
)

// Here we define the handlers for the routes
func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, the Id should be an integer."})
		return
	}

	event, err := operations.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful."})
}

// This function will cancel the registration for an event
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event operations.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled successfully."})
}
