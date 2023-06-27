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



//This function counts down the time in seconds as the elevator covers each floor
func DisplayCounter(floors int) {
	time.Sleep(time.Duration(5) * time.Second*5)
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




    

    // vars := mux.Vars(r)
    // currentFloor, targetFloor := vars["currentFloor"], vars["targetFloor"]

    //Convert current and target floor to integers
    // intcurrentfloor,err := strconv.Atoi(currentFloor)
    // if err!=nil{
	// 	log.Printf("Current floor value must be a string")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Current floor value must be a string"))
	// 	return
	// }
    // inttargetFloor,err := strconv.Atoi(targetFloor)
    // if err!=nil{
	// 	log.Printf("Target floor value must be a string")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("Target floor value must be a string"))
	// 	return
	// }

    

    // elevatorDir:=getElevatorDirection(intcurrentfloor,inttargetFloor)
    // elevator := Elevator{
    //     CurrentFloor: currentFloor,
    //     TargetFloor: targetFloor,
    //     State: "Calling",
    //     Direction: elevatorDir,
    // }

    // log.Printf("Elevator called from floor %s to floor %s", currentFloor,targetFloor)

   
   

	// // Create a new log entry
	// log := Log{
	// 	Timestamp: time.Now(),
	// 	Event:     "elevator_called",
	// 	Details:   getLogDetails(r),
	// }
	// db.Create(&log)

	// // TODO: Implement elevator movement logic here

	// w.WriteHeader(http.StatusOK)
}

func GetElevatorStatus(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement elevator status retrieval logic here
	w.WriteHeader(http.StatusOK)
}

