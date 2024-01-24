package controllers

import (
	"example/Backend/initializers"
	"example/Backend/models"

	"github.com/google/uuid"

	"time"

	"github.com/gin-gonic/gin"
)

func ParentsCreate(c *gin.Context) {
	addressId := uuid.New()
	//Get data off req body
	var body struct {
		IDNumber string
		Name     string
		Surname  string
		Number   string

		Street    string
		City      string
		ParentID  string
		CreatedAt time.Time
	}

	c.Bind(&body)

	//Create a parent with address

	// parent := models.Parent{ID: "3", Name: "Beloved", Surname: "Nethengwe", Number: "0813792428", CreatedAt: time.Now()}
	parent := models.Parent{IDNumber: body.IDNumber, Name: body.Name, Surname: body.Surname, Number: body.Number, CreatedAt: body.CreatedAt}
	parentAddress := models.Address{ID: addressId.String(), Street: body.Street, City: body.City, ParentID: body.IDNumber, CreatedAt: body.CreatedAt}
	result := initializers.DB.Create(&parent)
	addrResult := initializers.DB.Create(&parentAddress)

	if result.Error != nil {
		c.Status(400)
		return
	}
	if addrResult.Error != nil {
		c.Status(400)
		return
	}
	//Return it

	c.JSON(200, gin.H{
		"message": parent,
	})
}

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
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")
	//Get trhe data off req body
	var body struct {
		IDNumber string
		Name     string
		Surname  string
		Number   string

		Street    string
		City      string
		ParentID  string
		CreatedAt time.Time
	}

	c.Bind(&body)

	addressId := uuid.New()

	//Find the post where updating
	var post models.Parent
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.Parent{IDNumber: body.IDNumber, Name: body.Name, Surname: body.Surname, Number: body.Number, CreatedAt: body.CreatedAt})
	initializers.DB.Model(&post).Updates(models.Address{ID: addressId.String(), Street: body.Street, City: body.City, ParentID: body.IDNumber, CreatedAt: body.CreatedAt})

	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {

	//Get the id off the url
	id := c.Param("id")

	//Delete the parent
	initializers.DB.Where("parent_id = ?", id).Delete(&models.Address{})
	initializers.DB.Where("id_number = ?", id).Delete(&models.Parent{})

	//Respond
	c.Status(200)
}
