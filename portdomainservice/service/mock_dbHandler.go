package service

import (
	"context"
	"fmt"
	"sync"
)

var ErrorNotFound = fmt.Errorf("item not found")

type MockDBHandler struct {
	lock   *sync.Mutex
	client map[string]*Port
	fail   error
}

func NewMockDBHandler() *MockDBHandler {
	return &MockDBHandler{
		lock:   &sync.Mutex{},
		client: make(map[string]*Port),
	}
}

func (mock *MockDBHandler) SetFail(err error) {
	mock.fail = err
}
func (mock *MockDBHandler) ResetFail() {
	mock.fail = nil
}

func (mock *MockDBHandler) Connect(ctx context.Context, endpoint string) error {
	if mock.fail != nil {
		return mock.fail
	}
	return nil
}

func (mock *MockDBHandler) GetPort(ctx context.Context, port *Port) (*Port, error) {
	mock.lock.Lock()
	defer mock.lock.Unlock()
	if mock.fail != nil {
		return nil, mock.fail
	}
	retPort, found := mock.client[port.Id]
	if !found {
		return nil, ErrorNotFound
	}
	return retPort, nil
}
func (mock *MockDBHandler) GetAllPorts(ctx context.Context) (map[string]*Port, error) {
	if mock.fail != nil {
		return nil, mock.fail
	}
	return mock.client, nil
}
func (mock *MockDBHandler) CreateOrUpdatePort(ctx context.Context, port *Port) (*Port, error) {
	mock.lock.Lock()
	defer mock.lock.Unlock()
	if mock.fail != nil {
		return nil, mock.fail
	}
	if port.Id == "" {
		return nil, ErrorNotFound
	}
	mock.client[port.Id] = port
	return mock.client[port.Id], nil
}
func (mock *MockDBHandler) DeletePort(ctx context.Context, port *Port) error {
	mock.lock.Lock()
	defer mock.lock.Unlock()
	if mock.fail != nil {
		return mock.fail
	}
	delete(mock.client, port.Id)
	return nil
}
