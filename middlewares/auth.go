package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/utils"
)

func Authenticate(context *gin.Context) {

	//Get the Token from the header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
		return
	}

	//Check if the token is valid
	userId, err := utils.VerifyToken(token)

	if err != nil {	
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	//End validation of token

	//Set the userId in the context
	context.Set("userId", userId)

	context.Next()


}