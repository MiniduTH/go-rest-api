package routes

import (
	"REST_API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not get events from database"})
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create event", "error": err.Error()})
		return
	}

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event to database", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
