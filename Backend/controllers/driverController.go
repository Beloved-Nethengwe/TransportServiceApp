package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"
	"net/http"
	"strconv"
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
		Email                 string
		Password              string
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
		Email:                 body.Email,
		Password:              body.Password,
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

func ViewTransportRequests(c *gin.Context) {
	driver_id := c.Param("id")

	var transportRequest []struct {
		Name            string
		Allergy         string
		Destination     string
		PickUp          string
		P_Name          string
		CellphoneNumber string
		Email           string
		IDNumber        string
	}

	if err := initializers.DB.Raw(`
    SELECT c.name, allergy, destination, pick_up,
    p_name, cellphone_number, email
    FROM children c
    INNER JOIN request_bridges rb ON c.id = rb.child_id
    INNER JOIN parents p ON rb.parent_id = p.id
    WHERE rb.driver_id = ?`, driver_id).Scan(&transportRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"ok":       true,
		"requests": transportRequest,
	})
}

func UpdateRequestStatus(c *gin.Context) {
	driver_id := c.Param("driver_id")
	child_id := c.Param("child_id")

	childId, err := strconv.Atoi(child_id)
	driverId, err := strconv.Atoi(driver_id)
	requestStatus := "Assigned"

	var post models.RequestBridge
	initializers.DB.First(&post, &driverId, &childId)

	initializers.DB.Model(&post).
		Where("driver_id", &driverId).
		Where("child_id", &childId).
		Updates(models.RequestBridge{
			Status: requestStatus,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}
