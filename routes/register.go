package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"
	"github.com/gin-gonic/gin"
)


func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not parse Event Id, error": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not fetch event": err.Error()})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not register for event, error": err.Error()})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "Registered for event!"})

}


func cancelRegister(context *gin.Context) {
	//Get the userId from the context
	userId := context.MustGet("userId").(int64)
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not parse Event Id, error": err.Error()})
		return
	}

	
	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not cancel register for event, error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Canceled register for event!"})

}
