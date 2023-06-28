package loans_routes

//The routes package specifies how the loans API will handle various api calls

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovoulo/Ezra-Assessment/task-b/loans_handlers"
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
        HandlerFunc: loans_handlers.HomeHandler,
		
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