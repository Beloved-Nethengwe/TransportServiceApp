package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDestination(c *gin.Context) {

	var body struct {
		SchoolName string
		DriverID   string
	}

	c.Bind(&body)

	destination := models.Destination{
		SchoolName: body.SchoolName,
		DriverID:   body.DriverID,
	}
	result := initializers.DB.Create(&destination)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"message": destination,
	})
}

func DeleteDesstination(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Where("id = ?", id).Delete(models.Destination{})
	if result.Error != nil {
		// Handle the error, e.g., log it or return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Check if any records were deleted
	if result.RowsAffected == 0 {
		// No records were found with the specified ID
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}

	//Respond
	c.Status(200)
}
