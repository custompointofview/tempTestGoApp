package service

import (
	"context"
)

type DBHandlerInterf interface {
	Connect(context.Context, string) error
	GetPort(context.Context, *Port) (*Port, error)
	GetAllPorts(context.Context) (map[string]*Port, error)
	CreateOrUpdatePort(context.Context, *Port) (*Port, error)
	DeletePort(context.Context, *Port) error
}
