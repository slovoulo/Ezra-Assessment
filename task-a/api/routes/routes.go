package routes
//The routes package specifies how the elevator API will handle various api calls

import (
	"net/http"

	"github.com/gorilla/mux"
	
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
        Name: "Call Elevator",
        Method: "POST",
        Path: "/v1/elevator",
        HandlerFunc: recipegrpcclient.CreateRecipe,

    },

    {
        Name: "Get elevator info",
        Method: "GET",
        Path: "/v1/elevator",
        HandlerFunc: recipegrpcclient.GetSingleRecipe,

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