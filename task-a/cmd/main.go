package main

import (
	"log"
	"net/http"

	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/slovoulo/Ezra-Assessment/task-a/api/database"
	"github.com/slovoulo/Ezra-Assessment/task-a/api/routes"

	_ "github.com/slovoulo/Ezra-Assessment/task-a/api/handlers/docs"
	"github.com/swaggo/http-swagger"
)




func main(){
    // Create a custom route matcher
	//routerr := mux.NewRouter().MatcherFunc(caseInsensitiveMatcher).Subrouter().StrictSlash(true)
	
	log.Println("Starting elevator  service")
    //Connect to postgres db
    database.ConnectDB()
   
    
    muxRouter := mux.NewRouter().StrictSlash(true)
  

	// Serve Swagger UI

	muxRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	

	//specify who's allowed to connect
	c:=cors.New(cors.Options{ 
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
})
	//router := airfryerdatabase.AddRoutes(muxRouter)
	router := routes.AddRoutes(muxRouter)
	handler := c.Handler(router)

    // Make URLs case insensitive
	handler = LowerCaseURI(handler)
	
	
	err := http.ListenAndServe(":7070", handler) //Uncomment this line when using docker
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

	
}
func LowerCaseURI(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
      r.URL.Path = strings.ToLower(r.URL.Path)
      h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
  }