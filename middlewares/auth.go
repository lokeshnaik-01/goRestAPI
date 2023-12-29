package middlewares

import (
	"net/http"

	"example.com/restAPI/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized"})
		return
	}
	
	err,userId := utils.VerifyToken((token))
	if err!=nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Unauthorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}