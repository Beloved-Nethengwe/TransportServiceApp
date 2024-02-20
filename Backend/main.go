package main

import (
	"example/Backend/controllers"
	"example/Backend/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"} // Replace with your allowed origins
	corsConfig.AllowCredentials = true                          // To allow sending tokens to the server
	corsConfig.AddAllowMethods("OPTIONS")                       // Enable OPTIONS method for ReactJS
	router.Use(cors.New(corsConfig))

	router.POST("/parents", controllers.ParentsCreate)
	router.GET("/parents", controllers.ParentsIndex)
	router.GET("/parents/:id", controllers.ParentsShow)
	router.PUT("/parents/:id", controllers.ParentUpdate)
	router.DELETE("/parents/:id", controllers.ParentDelete)

	router.POST("/children", controllers.ChildCreate)
	router.GET("/children", controllers.ViewChildren)
	router.GET("/children/:id", controllers.ChildById)
	router.PUT("/children/:id", controllers.UpdateChild)
	router.DELETE("/children/:id", controllers.DeleteChild)

	router.POST("/driver", controllers.CreateDriver)
	router.GET("/driver", controllers.ViewDrivers)
	router.GET("/driver/:id", controllers.DriverByID)
	router.PUT("/driver/:id", controllers.UpdateDriver)
	router.DELETE("/driver/:id", controllers.DeleteDriver)

	router.POST("/destination", controllers.CreateDestination)
	router.DELETE("/destination/:id", controllers.CreateDestination)

	router.Run()
}
