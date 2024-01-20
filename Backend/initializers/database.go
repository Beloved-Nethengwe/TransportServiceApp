package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	/***********************************************/
	/* Single Connection to TimescaleDB/ PostgreSQL */
	/***********************************************/
	var err error
	dsn := os.Getenv("DB_CONN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect")
	}

}
