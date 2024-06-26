package routes

import (
	"net/http"
	"strconv"

	"example.com/restAPI/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"get events failed"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"fetch event failed"})
		return
	}
	context.JSON(http.StatusOK, event)
}
// once we register the function in the handler
// context is directly passed by gin
func getEvents(context *gin.Context) {
	// JSON returns statusCode and data
	// gin.H creates a map with string key and any values map[string]any
    // context.JSON(http.StatusOK, gin.H{"message":"Hello"})
	events, err := models.GetAllEvents()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"get events failed"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	
	var event models.Event
	// gin extracts the input params and matches it with struct
	err := context.ShouldBindJSON(&event)
	if err!=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"save events failed"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})

}


func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"get events failed"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"couldnt not fetch"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Can't update"})
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err!=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"update event failed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "updated event"})
}


func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"get events failed"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"couldnt not fetch"})
		return
	}
	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Can't delete"})
	}
	err = event.Delete()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"dekete event failed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
