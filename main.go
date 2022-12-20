package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/shadowshot-x/ActuatorBuf-grpc/grpcserver"
	"github.com/shadowshot-x/ActuatorBuf-grpc/protobufs"
	"github.com/shadowshot-x/ActuatorBuf-grpc/rest"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func restServer(p *atomic.Value) {
	fmt.Println("rest server started")
	// / REST ROUTER FOR MAINTAINER INTERACTION
	mainRouter := mux.NewRouter()

	maintainerController := rest.MaintainerController{P: p}

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

func grpcServer(p *atomic.Value) {
	fmt.Println("gRPC server started")
	// GRPC ROUTER FOR CLIENT INTERACTION
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	protobufs.RegisterPingRPCServer(s, &grpcserver.PingServer{})
	protobufs.RegisterActuatorServer(s, &grpcserver.ActuatorServer{P: p})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	// Atomic state map
	var simpAtomicMap atomic.Value
	var wg sync.WaitGroup
	wg.Add(2)
	go restServer(&simpAtomicMap)
	go grpcServer(&simpAtomicMap)
	wg.Wait()
}
