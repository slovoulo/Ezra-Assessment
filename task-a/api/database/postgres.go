package database

import (
	"fmt"
	"log"
	"os"



	"github.com/slovoulo/Ezra-Assessment/task-a/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

func ConnectDB() {
	// Loading environment variables
	dbPort := os.Getenv("ELEVATORDBPORT")
	dbUser := os.Getenv("ELEVATORDBUSER")
	dbName := os.Getenv("ELEVATORDBNAME")
	password := os.Getenv("ELEVATORDBPASSWORD")

	


	log.Println("attempting to connect to postgres")

	dbURI := fmt.Sprintf("host=elevatordatabase user=%s dbname=%s sslmode=disable password=%s port=%s", dbUser, dbName, password,dbPort)  //Uncomment when using Docker 
	
	//Opening connection to database
	Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	

	if err != nil {
		log.Printf("An error occured connecting to db %e", err)
		return
	} else {
		fmt.Println("Successfully connected to db")
	}



	Db.AutoMigrate(&models.Elevator{})

}