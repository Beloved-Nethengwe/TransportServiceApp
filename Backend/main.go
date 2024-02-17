package main

import (
	"example/Backend/controllers"
	"example/Backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	r.POST("/parents", controllers.ParentsCreate)
	r.GET("/parents", controllers.ParentsIndex)
	r.GET("/parents/:id", controllers.ParentsShow)
	r.PUT("/parents/:id", controllers.ParentUpdate)
	r.DELETE("/parents/:id", controllers.ParentDelete)

	r.POST("/children", controllers.ChildCreate)
	r.GET("/children", controllers.ViewChildren)
	r.GET("/children/:id", controllers.ChildById)
	r.PUT("/children/:id", controllers.UpdateChild)
	r.DELETE("/children/:id", controllers.DeleteChild)

	r.POST("/driver", controllers.CreateDriver)
	r.GET("/driver", controllers.ViewDrivers)
	r.GET("/driver/:id", controllers.DriverByID)
	r.PUT("/driver/:id", controllers.UpdateDriver)
	r.DELETE("/driver/:id", controllers.DeleteDriver)

	r.POST("/destination", controllers.CreateDestination)
	r.DELETE("/destination/:id", controllers.CreateDestination)

	r.Run()
}
