package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	apis "clientapi/apis"
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
	return nil
}

func (rh *RESTHandler) Close() error {
	return nil
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
}

func (rh *RESTHandler) DeletePort(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: DeletePort")
}

func (rh *RESTHandler) CreateOrUpdatePort(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: CreateOrUpdatePort")
}
