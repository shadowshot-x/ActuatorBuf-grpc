package main

import (
	"fmt"
	"net/http"
	"shadowshot-x/actuatorbuf/rest"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	/// REST ROUTER FOR MAINTAINER INTERACTION
	mainRouter := mux.NewRouter()

	maintainerController := rest.MaintainerController{}

	mainRouter.HandleFunc("/ping", rest.PingHandler)
	mainRouter.HandleFunc("/variable", maintainerController.StatePostHandler).Methods("POST")
	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	server := &http.Server{
		Addr:    ":9090",
		Handler: cors(mainRouter),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Booting the Server", err)
	}
}
