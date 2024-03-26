package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ParentsCreate(c *gin.Context) {

	var body struct {
		ID              string
		IDNumber        string
		PName           string
		Surname         string
		CellphoneNumber string
		Address         string
		Email           string
		CreatedAt       time.Time
		RoleID          int
	}

	c.Bind(&body)

	tx := initializers.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Transaction failed"})
			return
		}
	}()

	// parent := models.Parent{ID: "3", Name: "Beloved", Surname: "Nethengwe", Number: "0813792428", CreatedAt: time.Now()}
	parent := models.Parent{
		ID:              body.ID,
		IDNumber:        body.IDNumber,
		PName:           body.PName,
		Surname:         body.Surname,
		CellphoneNumber: body.CellphoneNumber,
		Address:         body.Address,
		Email:           body.Email,
		CreatedAt:       body.CreatedAt,
		RoleID:          body.RoleID,
	}

	if err := tx.Create(&parent).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error creating parent: %v", err)})
		return
	}

	tx.Commit()
	// result := initializers.DB.Create(&parent)
	// roleResult := initializers.DB.Create(&role)

	// if result.Error != nil {
	// 	c.Status(400)
	// 	return
	// }

	// if roleResult.Error != nil {
	// 	c.Status(400)
	// 	return
	// }
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
	var parent models.Parent
	result := initializers.DB.First(&parent, &id)

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
		"parent": parent,
	})
}

func ParentUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")
	//Get trhe data off req body
	var body struct {
		IDNumber        string
		PName           string
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
	initializers.DB.Model(&post).Updates(models.Parent{IDNumber: body.IDNumber, PName: body.PName, Surname: body.Surname, CellphoneNumber: body.CellphoneNumber, Address: body.Address, CreatedAt: body.CreatedAt})

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

func RequestChildTransport(c *gin.Context) {
	child_id := c.Param("child_id")
	driver_id := c.Param("driver_id")
	parent_id := c.Param("parent_id")
	status := "Pending"

	childId, err := strconv.Atoi(child_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	transportRequest := models.RequestBridge{
		Status:   status,
		ParentID: parent_id,
		DriverID: driver_id,

		ChildID: childId,
	}

	result := initializers.DB.Create(&transportRequest)

	if result.Error != nil {
		c.Status(400)
		return
	}

}

func SearchChildTransportByDestination(c *gin.Context) {
	child_destination := c.Param("child_destination")

	var transportByDestination []struct {
		ID              string
		IDNumber        string
		Name            string
		Surname         string
		CellphoneNumber string
		SchoolName      string
	}

	if err := initializers.DB.Raw(
		`SELECT d.id,d.id_number,d."name",d.surname,d.cellphone_number,d2.school_name 
		FROM public.drivers d 
		INNER JOIN destinations d2 ON d2.driver_id  = d.id
		where d2.school_name = ? `, child_destination).Scan(&transportByDestination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"ok":                   true,
		"transportDestination": transportByDestination,
	})

}

func SendMailToDriver(c *gin.Context) {
	driver_mail := c.Param("driver_mail")
	child_name := c.Param("child_name")

	auth := smtp.PlainAuth(
		"",
		"belovednethengwe28@gmail.com",
		"icjy gjmh impc jssk",
		"smtp.gmail.com",
	)

	msg := "Subject: Transport Request\nYou have a new request for child " + child_name

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"belovednethengwe28@gmail.com",
		[]string{driver_mail},
		[]byte(msg),
	)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(200, gin.H{"message": "Email sent successfully"})
}
