package routes
import (
	"github.com/gin-gonic/gin"
	"example.com/restAPI/middlewares"
)
func RegisterRoutes(server *gin.Engine) {
	// setting handler for incoming GET requests
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	// by grouping this we make sure middleware is always executed
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)


	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}