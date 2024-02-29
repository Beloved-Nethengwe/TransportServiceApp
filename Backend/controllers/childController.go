package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChildCreate(c *gin.Context) {

	var body struct {
		Name         string
		Surname      string
		Allergy      string
		EmergContact string
		PickUp       string
		Destination  string
		ParentID     string
	}
	c.Bind(&body)

	child := models.Child{
		Name:         body.Name,
		Surname:      body.Surname,
		Allergy:      body.Allergy,
		EmergContact: body.EmergContact,
		PickUp:       body.PickUp,
		Destination:  body.Destination,
		ParentID:     body.ParentID,
	}
	result := initializers.DB.Create(&child)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": child,
	})
}

func ViewChildren(c *gin.Context) {

	var children []models.Child
	initializers.DB.Find(&children)

	c.JSON(200, gin.H{
		"posts": children,
	})
}

func ChildById(c *gin.Context) {
	id := c.Param("id")

	var post models.Child
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

func GetChildrenByParentID(c *gin.Context) {
	parentID := c.Param("parent_id")

	var children []models.Child
	if err := initializers.DB.Find(&children, "parent_id = ?", parentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"ok":       true,
		"children": children,
	})
}

func DeleteChild(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Where("id = ?", id).Delete(models.Child{})

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

func UpdateChild(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Name         string
		Surname      string
		Allergy      string
		EmergContact string
		PickUp       string
		Destination  string
		ParentID     string
	}

	c.Bind(&body)

	var post models.Child
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Child{
		Name:         body.Name,
		Surname:      body.Surname,
		Allergy:      body.Allergy,
		EmergContact: body.EmergContact,
		PickUp:       body.PickUp,
		Destination:  body.Destination,
		ParentID:     body.ParentID,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}
