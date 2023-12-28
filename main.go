package main
import (
	"github.com/gin-gonic/gin"
	"example.com/restAPI/db"
	"net/http"
	"example.com/restAPI/models"
	"strconv"
)

func main() {
	db.InitDB()
	// default configures a http service
	server := gin.Default()

	// setting handler for incoming GET requests
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	// run starts listening for incoming req on port 8080 on some domain
	server.Run(":8080") //localhost:8080
}


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
	event.ID =1
	event.UserID = 1
	err = event.Save()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"save events failed"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})

}