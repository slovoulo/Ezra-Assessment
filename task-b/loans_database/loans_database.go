package loans_database

import (
	"fmt"
	"log"
	"os"



	"github.com/slovoulo/Ezra-Assessment/task-b/loans_models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

func ConnectDB() {
	// Loading environment variables
	dbPort := os.Getenv("LOANDBPORT")
	dbUser := os.Getenv("LOANDBUSER")
	dbName := os.Getenv("LOANDBNAME")
	password := os.Getenv("LOANDBPASSWORD")

	


	log.Println("attempting to connect to loans database")

	dbURI := fmt.Sprintf("host=loansdatabase user=%s dbname=%s sslmode=disable password=%s port=%s", dbUser, dbName, password,dbPort)  //Uncomment when using Docker 
	
	//Opening connection to database
	Db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	

	if err != nil {
		log.Printf("An error occured connecting to db %e", err)
		return
	} else {
		fmt.Println("Successfully connected to db")
	}



	Db.AutoMigrate(&loans_models.Account{})
	Db.AutoMigrate(&loans_models.TransactionEntries{})

}