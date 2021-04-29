package apis

import (
	pds "clientapi/portdomainservice"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockPDSClient struct{}

func (m *MockPDSClient) GetPort(ctx context.Context, in *pds.PortRequest, opts ...grpc.CallOption) (*pds.PortResponse, error) {
	return &pds.PortResponse{}, nil
}
func (m *MockPDSClient) CreateOrUpdatePort(ctx context.Context, in *pds.PortRequest, opts ...grpc.CallOption) (*pds.PortResponse, error) {
	return &pds.PortResponse{}, nil
}
func (m *MockPDSClient) DeletePort(ctx context.Context, in *pds.PortRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}
