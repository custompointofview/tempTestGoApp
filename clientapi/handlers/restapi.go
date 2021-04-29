package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	apis "clientapi/apis"
	pds "clientapi/portdomainservice"
)

type RESTHandlerInterf interface {
	Homepage(w http.ResponseWriter, r *http.Request)
	GetAllPorts(w http.ResponseWriter, r *http.Request)
	GetPort(w http.ResponseWriter, r *http.Request)
	DeletePort(w http.ResponseWriter, r *http.Request)
	CreateOrUpdatePort(w http.ResponseWriter, r *http.Request)
	InitialSetup(context.Context, string) error
	Close() error
}

type RESTHandler struct {
	api apis.APIInterf
}

func NewRESTHandler(api apis.APIInterf) *RESTHandler {
	return &RESTHandler{
		api: api,
	}
}

func (rh *RESTHandler) InitialSetup(ctx context.Context, path string) error {
	if err := rh.api.Connect(ctx); err != nil {
		return err
	}
	if path == "" {
		log.Printf("! No initial DB was provided. Skipping initialization.")
		return nil
	}
	if err := rh.api.InitializeDatabase(path); err != nil {
		return err
	}
	return nil
}

func (rh *RESTHandler) Close() error {
	return rh.api.Close()
}

func (rh *RESTHandler) Homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: Homepage")
	fmt.Fprintf(w, "Welcome to the Homepage!")
}

func (rh *RESTHandler) GetAllPorts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: GetAllPorts")
}

func (rh *RESTHandler) GetPort(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: GetPort")
	vars := mux.Vars(r)
	portID := vars["id"]

	log.Printf("Requested Port: %s\n", portID)
	port, err := rh.api.GetPort(portID)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	log.Printf("Response: %v", port)
	json.NewEncoder(w).Encode(port)
}

func (rh *RESTHandler) DeletePort(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: DeletePort")

	vars := mux.Vars(r)
	portID := vars["id"]

	log.Printf("Requested Port: %s\n", portID)
	if err := rh.api.DeletePort(portID); err != nil {
		http.Error(w, err.Error(), 404)
	}
	log.Printf("Deleted Port: %s\n", portID)
}

func (rh *RESTHandler) CreateOrUpdatePort(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: CreateOrUpdatePort")

	vars := mux.Vars(r)
	portID := vars["id"]

	port := &pds.Port{Id: portID}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(port); err != nil {
		http.Error(w, err.Error(), 400)
	}

	log.Printf("Create or Update: %v", port)
	if err := rh.api.CreateOrUpdatePort(port); err != nil {
		http.Error(w, err.Error(), 404)
	}
}
