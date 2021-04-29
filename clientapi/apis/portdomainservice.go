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
