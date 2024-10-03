package routes

import (

	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not parse Event Id, error": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not read event, error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)

}

func GetEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not read events, error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func CreteEvents(context *gin.Context) {

	/*Reemplaze this code with the middleware
	//Get the Token from the header
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "No token provided"})
		return
	}

	//Check if the token is valid
	userId, err := utils.VerifyToken(token)

	if err != nil {	
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}
	//End validation of token
	*/



	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not create event, error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")

	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not create event, error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not parse Event Id, error": err.Error()})
		return
	}

	
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not read event, error": err.Error()})
		return
	}

	//Only the User that created the event can update it
	//Get the userId from the context
	userId := context.GetInt64("userId")

	if event.UserID != userId {  
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission to update this event."})
		return
	}






	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not update event, error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not update event, error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!", "event": event})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Coult not parse Event Id, error": err.Error()})
		return
	}

	
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not delete event, error": err.Error()})
		return
	}

	//Only the User that created the event can update it
	//Get the userId from the context
		userId := context.GetInt64("userId")

		if event.UserID != userId {  
			context.JSON(http.StatusUnauthorized, gin.H{"message": "You do not have permission to update this event."})
			return
		}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Coult not delete event, error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
