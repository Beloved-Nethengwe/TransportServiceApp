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
	r.PUT("/parents/:id", controllers.PostUpdate)

	r.Run()
}

// {

// 	"IDNumber": "23",
// 	"Name": "John",
// 	"Surname": "Doe",
// 	"Number": "123456789",
// 	"Street": "43 Bosbok",
// 	"City": "Pretoria",
// 	"CreatedAt": "time.Now()"

// }
