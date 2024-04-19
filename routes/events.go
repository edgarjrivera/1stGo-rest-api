package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api/operations"
	"github.com/gin-gonic/gin"
)

// Here we define the handlers for the routes
func getEvents(context *gin.Context) {
	events, err := operations.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

// This function will get a single event by its ID
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, the Id should be an integer."})
		return
	}

	// Get the event by its ID
	event, err := operations.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, event)
}

// This function will create an event
func createEvent(context *gin.Context) {
	var event operations.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err.Error()) // Logging the error
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	userId := context.GetInt64("userId")

	// Set the user ID
	event.UserID = userId
	event.DateTime = time.Now().Local()

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// This function will update an event
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, the Id should be an integer."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := operations.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event."})
		return
	}

	var updatedEvent operations.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data" + err.Error()})
		return
	}

	updatedEvent.ID = eventId
	updatedEvent.DateTime = updatedEvent.DateTime.Local()
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

// This function will delete an event
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID, the Id should be an integer."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := operations.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again later."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
