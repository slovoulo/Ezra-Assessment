package routes

//The routes package specifies how the elevator API will handle various api calls

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovoulo/Ezra-Assessment/task-a/api/handlers"
)

type Route struct {
	Method      string
	Name        string
	Path        string
	HandlerFunc http.HandlerFunc
}

var routes=[]Route{

    //Home
    {
		Name:        "welcomeScreen",
		Method:      "GET",
		Path:     "/v1/",
        HandlerFunc: handlers.HomeHandler,
		
	},
	


    {
        Name: "Call Elevator",
        Method: "POST",
        Path: "/v1/elevator",
        HandlerFunc: handlers.CallElevator,

    },

   
    
}

func AddRoutes (router *mux.Router) *mux.Router{
    for _, route:= range routes{
        router.
        Methods(route.Method).
        Name(route.Name).
        Path(route.Path).
        HandlerFunc(route.HandlerFunc)

    }
    return router
}