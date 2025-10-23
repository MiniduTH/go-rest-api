package main

import (
	"REST_API/db"
	"REST_API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8090")
	if err != nil {
		panic("An error occurred in server")
	}

}

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
