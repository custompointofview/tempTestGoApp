package apis

import (
	"context"

	pds "clientapi/portdomainservice"

	"google.golang.org/grpc"
)

type PortDomainServAPI struct {
	ctx       context.Context
	endpoint  string
	conn      *grpc.ClientConn
	pdsClient pds.PortDomainServiceClient
}

func NewPortDomainServAPI(endpoint string) *PortDomainServAPI {
	return &PortDomainServAPI{
		endpoint: endpoint,
	}
}

func (p *PortDomainServAPI) Connect(ctx context.Context) error {
	return nil
}

func (p *PortDomainServAPI) Close() error {
	return nil
}

func (p *PortDomainServAPI) InitializeDatabase(jsonFilePath string) error {
	return nil
}

func (p *PortDomainServAPI) GetPort(id string) (*pds.Port, error) {
	return nil, nil
}
func (p *PortDomainServAPI) GetAllPorts() (map[string]*pds.Port, error) {
	return nil, nil
}
func (p *PortDomainServAPI) CreateOrUpdatePort(port *pds.Port) error {
	return nil
}
func (p *PortDomainServAPI) DeletePort(id string) error {
	return nil
}
