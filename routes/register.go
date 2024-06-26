package routes

import (
	"net/http"
	"strconv"

	"example.com/restAPI/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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
	err = event.Register(userId)
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"registration failed"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"get events failed"})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.DeleteRegistration(userId)
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"cancellation failed"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Cancelled"})

}