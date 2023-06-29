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
    {
		Name:        "Create loan account",
		Method:      "POST",
		Path:     "/v1/account",
        HandlerFunc: loans_handlers.CreateAccount,
		
	},
    {
		Name:        "Request loan",
		Method:      "POST",
		Path:     "/v1/loanrequest",
        HandlerFunc: loans_handlers.LoanRequest,
		
	},
    {
		Name:        "Repay loan",
		Method:      "POST",
		Path:     "/v1/loanrepayment",
        HandlerFunc: loans_handlers.LoanRepayment,
		
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