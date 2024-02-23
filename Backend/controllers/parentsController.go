package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ParentsCreate(c *gin.Context) {
	// addressId := uuid.New()
	//Get data off req body
	var body struct {
		ID              string
		IDNumber        string
		Name            string
		Surname         string
		CellphoneNumber string
		Address         string
		Email           string
		Password        string
		CreatedAt       time.Time
	}

	c.Bind(&body)

	//Create a parent with address

	// parent := models.Parent{ID: "3", Name: "Beloved", Surname: "Nethengwe", Number: "0813792428", CreatedAt: time.Now()}
	parent := models.Parent{
		ID:              body.ID,
		IDNumber:        body.IDNumber,
		Name:            body.Name,
		Surname:         body.Surname,
		CellphoneNumber: body.CellphoneNumber,
		Address:         body.Address,
		Email:           body.Email,
		Password:        body.Password,
		CreatedAt:       body.CreatedAt,
	}

	result := initializers.DB.Create(&parent)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it

	c.JSON(200, gin.H{
		"message": parent,
	})
}

// https://stackoverflow.com/questions/66221270/how-to-join-two-tables-in-gorm
func ParentsIndex(c *gin.Context) {
	var posts []models.Parent
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func ParentsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	//Get the posts
	var post models.Parent
	result := initializers.DB.First(&post, id)

	//Error Check
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

func ParentUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")
	//Get trhe data off req body
	var body struct {
		IDNumber        string
		Name            string
		Surname         string
		CellphoneNumber string
		Address         string
		CreatedAt       time.Time
	}

	c.Bind(&body)

	//Find the post where updating
	var post models.Parent
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Parent{IDNumber: body.IDNumber, Name: body.Name, Surname: body.Surname, CellphoneNumber: body.CellphoneNumber, Address: body.Address, CreatedAt: body.CreatedAt})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func ParentDelete(c *gin.Context) {

	//Get the id off the url
	id := c.Param("id")

	//Delete the parent
	// Attempt to delete the parent

	result := initializers.DB.Where("id = ?", id).Delete(&models.Parent{})

	// Check for errors
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
