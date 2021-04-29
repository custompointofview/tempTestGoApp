package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"clientapi/apis"
	"clientapi/handlers"

	"github.com/gorilla/mux"
)

// DEFAULT PORTDOMAINSERVICE ENDPOINT
const PORTDOMAINSERVICE_ENDPOINT = "portdomainservice:8001"

func CreateRouter(handler handlers.RESTHandlerInterf) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.Homepage)
	router.HandleFunc("/ports", handler.GetAllPorts).Methods("GET")
	router.HandleFunc("/port/{id}", handler.GetPort).Methods("GET")
	router.HandleFunc("/port/{id}", handler.CreateOrUpdatePort).Methods("POST")
	router.HandleFunc("/port/{id}", handler.DeletePort).Methods("DELETE")
	return router
}

func main() {
	log.Println("CLIENT REST API v0.1 - Port Router")
	portsFilePath := flag.String("ports", "", "Path to ports file")
	flag.Parse()

	ctx, _ := context.WithCancel(context.Background())

	// select API & create handler
	api := apis.NewAPI(apis.API_PORTDOMAINSERV, PORTDOMAINSERVICE_ENDPOINT)
	handler := handlers.NewRESTHandler(api)
	// run initial setup
	if err := handler.InitialSetup(ctx, *portsFilePath); err != nil {
		log.Fatalf("Initial Setup Failed: %v", err)
	}
	defer handler.Close()

	// start server
	log.Printf("Launching REST server on 8000...")
	router := CreateRouter(handler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
