package service

import (
	"context"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type PDSServer struct {
	dbHandler DBHandlerInterf
}

func NewServer(dbHandler DBHandlerInterf) *PDSServer {
	return &PDSServer{
		dbHandler: dbHandler,
	}
}

func (pds *PDSServer) EstablishDBConnection(ctx context.Context, endpoint string) error {
	return nil
}

func (pds *PDSServer) GetPort(ctx context.Context, in *PortRequest) (*PortResponse, error) {
	log.Printf("Endpoint Hit: GetPort")
	log.Printf("Client Request: %v", in)
	return nil, nil
}

func (pds *PDSServer) CreateOrUpdatePort(ctx context.Context, in *PortRequest) (*PortResponse, error) {
	log.Printf("Endpoint Hit: CreateOrUpdatePort")
	log.Printf("Client Request: %v", in)

	return nil, nil
}

func (pds *PDSServer) DeletePort(ctx context.Context, in *PortRequest) (*emptypb.Empty, error) {
	log.Printf("Endpoint Hit: DeletePort")
	log.Printf("Client Request: %v", in)
	return nil, nil
}

func (pds *PDSServer) mustEmbedUnimplementedPortDomainServiceServer() {}
