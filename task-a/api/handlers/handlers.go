// @title Elevator app API
// @version 1.0.0
// @description Elevator API documentation.
// @host localhost:7070
// @Accept json
// @Produce json
// @BasePath /
package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	//"strconv"
	"time"


    "github.com/slovoulo/Ezra-Assessment/task-a/api/models"
    "github.com/slovoulo/Ezra-Assessment/task-a/api/database"

	//"github.com/gorilla/mux"
)

type RequestedElevator struct {
	FromFloor int `json:"from_floor"`
	ToFloor   int `json:"to_floor"`
}

type ElevatorRequest struct {
	
	ElevatorID   int `gorm:"not null"`
	
	CurrentFloor int    `gorm:"not null"`
	TargetFloor  int    `gorm:"not null"`
	State        string `gorm:"not null"`
	
	CallerName   string `gorm:"not null"`
	CallerID     string `gorm:"not null"`
}

// Home godoc
// @Summary Landing page
// @Description Landing page
// @Tags Elevator-App
// @Success 200 
// @Failure 400 
// @Router /v1/ [get]
func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the Elevator app!"))

}

//This function counts down the time in seconds as the elevator covers each floor
func DisplayCounter(floors int) {
    
	// time.Sleep(time.Duration(5) * time.Second*5)
	// log.Println("You have arrived")

    seconds:=floors*5
    for i := seconds; i > 0; i-- {
		log.Println(i)
		time.Sleep(1 * time.Second)
	}
    log.Println("You have arrived")

}
 //Determine whether elevator is going up or down
 func getElevatorDirection(currentf int, targetf int) string {
    //If current floor number is greater than target floor number it means the elevator is going down
    if(currentf>targetf){
        return "Going down"
    }else if(currentf<targetf){
        return "Going up"
    }
    return "Going up"

 }

///Annotation for calling elevator

// @Summary Call elevator
// @Description Call elevator: For floor numbers (Current floor and target floor) use numbers whose difference is small because this request waits 5 seconds per floor
// @Tags Elevator-App
// @Accept  json
// @Produce  json
// @Success 200 
// @Failure 400 
// @Param Elevator body ElevatorRequest true "Elevator struct"
// @Router /v1/elevator [post]
func CallElevator(w http.ResponseWriter, r *http.Request) {
    
    //Since information  of who/where the elevator is called from are required
    //This workflow assumes that the elevator can only be accessed by programmable key cards
    //Each keycard contains details of the user

    //Get the body of the request from the keycard
    reqBody, _ := io.ReadAll(r.Body)
    var elevator models.Elevator
    //Unmarshall the request into an Elevator struct
    json.Unmarshal(reqBody,&elevator)

    //Get elevator direction
    dir:=getElevatorDirection(elevator.CurrentFloor, elevator.TargetFloor)
    elevator.Direction=dir
    elevator.CalledAt=time.Now()
    log.Println("Calling elevator")
    

    //Log the details of the call to db
    if result := database.Db.Create(&elevator); result.Error !=nil{
        log.Println("An error occured logging elevator information")
        return
    }

    //Logic to wait 5 seconds per elevator floor called
    floorsToTravel:=0
    if(elevator.CurrentFloor>elevator.TargetFloor){
        floorsToTravel=elevator.CurrentFloor-elevator.TargetFloor
    }else if(elevator.CurrentFloor<elevator.TargetFloor){
        floorsToTravel=elevator.TargetFloor-elevator.CurrentFloor
    }
    DisplayCounter(floorsToTravel)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Calling elevator"))




    

 
}


