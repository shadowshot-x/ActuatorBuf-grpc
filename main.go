package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"shadowshot-x/actuatorbuf/grpcserver"
	"shadowshot-x/actuatorbuf/protobufs"
	"shadowshot-x/actuatorbuf/rest"
	"sync"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func restServer() {
	fmt.Println("gcp")
	// / REST ROUTER FOR MAINTAINER INTERACTION
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

func grpcServer() {
	fmt.Println("gcp")
	// GRPC ROUTER FOR CLIENT INTERACTION
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	protobufs.RegisterPingRPCServer(s, &grpcserver.PingServer{})
	protobufs.RegisterActuatorServer(s, &grpcserver.ActuatorServer{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go restServer()
	go grpcServer()
	wg.Wait()

}
