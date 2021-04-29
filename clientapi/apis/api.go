package apis

import (
	pds "clientapi/portdomainservice"
	"context"
	"log"
)

const API_PORTDOMAINSERV = 1

type apiType int

type APIInterf interface {
	Connect(context.Context) error
	Close() error
	InitializeDatabase(string) error

	GetPort(string) (*pds.Port, error)
	GetAllPorts() (map[string]*pds.Port, error)
	CreateOrUpdatePort(*pds.Port) error
	DeletePort(string) error
}

func NewAPI(apiT apiType, endpoint string) APIInterf {
	switch apiT {
	case API_PORTDOMAINSERV:
		return NewPortDomainServAPI(endpoint)
	default:
		log.Printf("API type undefined")
		return nil
	}
}
