package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateDriver(c *gin.Context) {

	var body struct {
		IDNumber              string
		Name                  string
		Surname               string
		CellphoneNumber       string
		Image                 string
		CarRegistrationNumber string
		CreatedAt             time.Time
	}
	c.Bind(&body)

	driver := models.Driver{
		IDNumber:              body.IDNumber,
		Name:                  body.Name,
		Surname:               body.Surname,
		CellphoneNumber:       body.CellphoneNumber,
		Image:                 body.Image,
		CarRegistrationNumber: body.CarRegistrationNumber,
		CreatedAt:             body.CreatedAt,
	}

	result := initializers.DB.Create(&driver)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": driver,
	})
}

func ViewDrivers(c *gin.Context) {
	var driver []models.Driver
	initializers.DB.Find(&driver)

	c.JSON(200, gin.H{
		"posts": driver,
	})
}

func DriverByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Driver
	result := initializers.DB.First(&post, &id)

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
	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeleteDriver(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Where("id = ?", id).Delete(models.Driver{})

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

func UpdateDriver(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		IDNumber              string
		Name                  string
		Surname               string
		CellphoneNumber       string
		Image                 string
		CarRegistrationNumber string
		CreatedAt             time.Time
	}

	c.Bind(&body)

	var post models.Driver
	initializers.DB.First(&post, &id)

	initializers.DB.Model(&post).Updates(models.Driver{
		IDNumber:              body.IDNumber,
		Name:                  body.Name,
		Surname:               body.Surname,
		CellphoneNumber:       body.CellphoneNumber,
		Image:                 body.Image,
		CarRegistrationNumber: body.CarRegistrationNumber,
		CreatedAt:             body.CreatedAt,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}
