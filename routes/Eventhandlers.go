package routes

import (
	"net/http"
	"strconv"

	models "crud.Restapi/crud/Models"
	"github.com/gin-gonic/gin"
)

func Events(context *gin.Context) {
	models, err := models.GetallEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Some Err occur in getting events",
		})
		return
	}
	context.JSON(http.StatusOK, models)
}

func Getevent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in param",
		})
		return
	}
	event, err := models.GetEventbyId(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in  geting data from function",
		})
		return
	}
	context.JSON(http.StatusAccepted, event)

}
func CreateEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Enter all fields"})
		return
	}
	userid := context.GetInt64("userid")
	event.Id = 1
	event.UserId = (userid)
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Some error occurred while creating event",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created", "event": event,
	})
}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID parameter",
		})
		return
	}
	userid := context.GetInt64("userid")
	event, err := models.GetEventbyId(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event ID not found",
		})
		return
	}

	if event.UserId != userid {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Not authorized to update event",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Please enter all fields correctly",
		})
		return
	}

	updatedEvent.Id = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
	})
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Error in parse int",
		})
		return
	}
	err = models.DeleteEventbyId(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event might not be found",
		})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{
		"message": "Event Deleted Succesfully",
	})
}
