package main

import (
	"go-project-one/go-api/db"
	"go-project-one/go-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connstr := "postgresql://kutay:kutay@localhost/local_dev?sslmode=disable"
	db.InitDB(connstr)
	db.CreateTables()

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8081")
}

func getEvents(c *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	//? create event
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request"})
		return
	}

	event.UserID = 1

	_, err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
