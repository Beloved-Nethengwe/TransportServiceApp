package main

//This is where you establish your database tables with relationships...
//after all has been done. I must run the migrate.go using this command
//go run .\migrate.go
import (
	"example/Backend/initializers"
	"example/Backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Parent{}, &models.Driver{}, &models.Child{}, &models.RequestBridge{}, &models.Destination{})

}
