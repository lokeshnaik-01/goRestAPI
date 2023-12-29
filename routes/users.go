package routes
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"example.com/restAPI/models"
	
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err!=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req data"})
		return
	}
	err = user.Save()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"users creation failed"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "user Created", "user": user})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err!=nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse req data"})
		return
	}
	err = user.ValidateCredentials()
	if err!= nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message":err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "logged in"})

}