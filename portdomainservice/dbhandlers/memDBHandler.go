package dbhandlers

import (
	"context"
	"fmt"
	"log"
	srv "portdomainservice/service"
	"sync"
)

var ErrorNotFound = fmt.Errorf("item not found")

type MemDBHandler struct {
	lock           *sync.Mutex
	client         map[string]*srv.Port
	databaseName   string
	collectionName string
}

func NewMemDBHandler(dbConfig *DBConfig) srv.DBHandlerInterf {
	return &MemDBHandler{
		lock:           &sync.Mutex{},
		client:         make(map[string]*srv.Port),
		databaseName:   dbConfig.database,
		collectionName: dbConfig.collection,
	}
}

func (mem *MemDBHandler) Connect(ctx context.Context, endpoint string) error {
	log.Printf("Connected to MEMORY DB endpoint @: %s", endpoint)
	return nil
}

func (mem *MemDBHandler) GetPort(ctx context.Context, port *srv.Port) (*srv.Port, error) {
	mem.lock.Lock()
	defer mem.lock.Unlock()
	if port.Id == "" {
		return nil, ErrorNotFound
	}
	retPort, found := mem.client[port.Id]
	if !found {
		log.Printf("ERROR finding requested port: %v", port.Id)
		return nil, ErrorNotFound
	}
	log.Printf("Returning : %v", retPort)
	return retPort, nil
}
func (mem *MemDBHandler) GetAllPorts(ctx context.Context) (map[string]*srv.Port, error) {
	return mem.client, nil
}
func (mem *MemDBHandler) CreateOrUpdatePort(ctx context.Context, port *srv.Port) (*srv.Port, error) {
	mem.lock.Lock()
	defer mem.lock.Unlock()
	if port.Id == "" {
		return nil, ErrorNotFound
	}
	mem.client[port.Id] = port
	return mem.client[port.Id], nil
}
func (mem *MemDBHandler) DeletePort(ctx context.Context, port *srv.Port) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()
	delete(mem.client, port.Id)
	return nil
}
