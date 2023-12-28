package main

import (
	"example.com/restAPI/db"
	"example.com/restAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// default configures a http service
	server := gin.Default()

	routes.RegisterRoutes(server)
	// run starts listening for incoming req on port 8080 on some domain
	server.Run(":8080") //localhost:8080
}


