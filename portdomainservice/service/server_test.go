package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	dbh := NewMockDBHandler()
	server := NewServer(dbh)

	t.Run("Connect", func(t *testing.T) {
		err := server.EstablishDBConnection(ctx, "")
		assert.Nil(t, err)

		expectedErr := fmt.Errorf("failed connection")
		dbh.SetFail(expectedErr)
		defer dbh.ResetFail()

		err = server.EstablishDBConnection(ctx, "")
		assert.NotNil(t, err)
		assert.EqualError(t, expectedErr, err.Error())
	})
	t.Run("CreateOrUpdatePort_NoId", func(t *testing.T) {
		req := &PortRequest{Port: &Port{}}
		resp, err := server.CreateOrUpdatePort(ctx, req)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrorNotFound.Error())
	})
	t.Run("CreateOrUpdatePort_HappyFlow", func(t *testing.T) {
		req := &PortRequest{Port: &Port{Id: "AEDXB", Code: "52000"}}
		resp, err := server.CreateOrUpdatePort(ctx, req)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.Equal(t, req.Port.Id, resp.Port.Id)
	})
	t.Run("GetPort_NoId", func(t *testing.T) {
		req := &PortRequest{Port: &Port{}}
		resp, err := server.GetPort(ctx, req)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrorNotFound.Error())
	})
	t.Run("GetPort_NotFound", func(t *testing.T) {
		req := &PortRequest{Port: &Port{Id: "AEFJR"}}
		resp, err := server.GetPort(ctx, req)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrorNotFound.Error())
	})
	t.Run("GetPort_HappyFlow", func(t *testing.T) {
		expectedPort := &Port{Id: "AEDXB"}
		req := &PortRequest{Port: expectedPort}
		resp, err := server.GetPort(ctx, req)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.Equal(t, expectedPort.Id, resp.Port.Id)
		assert.Equal(t, "52000", resp.Port.Code)
	})
}
