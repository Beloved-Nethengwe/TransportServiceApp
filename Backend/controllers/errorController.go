package controllers

import (
	"example/Backend/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func userExists(username string) (bool, error) {
	var count int

	err := initializers.DB.Raw("SELECT COUNT(*) FROM public.parents WHERE email = ?", username).Scan(&count).Error
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %w", err)
	}

	return count > 0, nil
}

func CheckIfUserAlreadyExist(c *gin.Context) {
	username := c.Query("email")
	if username == "" {
		c.JSON(400, gin.H{"error": "Missing username parameter"})
		return
	}

	exists, err := userExists(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"exists": exists})
}
