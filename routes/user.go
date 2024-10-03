package routes

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Could not parse request  user data, error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not save user, error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created sucessfully!", "user": user})
}

// Login function, validate the user credentials
func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Could not parse request  user data, error": err.Error()})
		return
	}

	err = user.ValidateCredential()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"Could not authenticate user1, error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Could not authentica User2, error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User authenticated sucessfully!", "token": token})
}
