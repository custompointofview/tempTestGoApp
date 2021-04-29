package dbhandlers

import (
	"context"
	srv "portdomainservice/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemDBHandler(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		h := NewMemDBHandler(&DBConfig{})
		assert.NotNil(t, h)
		assert.IsType(t, &MemDBHandler{}, h)
	})
}

func TestMemDBHandler(t *testing.T) {
	ctx := context.Background()

	h := NewMemDBHandler(&DBConfig{})
	assert.NotNil(t, h)
	assert.IsType(t, &MemDBHandler{}, h)

	t.Run("Connect", func(t *testing.T) {
		err := h.Connect(ctx, "endpoint")
		assert.Nil(t, err)
	})
	t.Run("CreateOrUpdatePort_NoId", func(t *testing.T) {
		port, err := h.CreateOrUpdatePort(ctx, &srv.Port{})
		assert.Nil(t, port)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrorNotFound.Error())
	})
	t.Run("CreateOrUpdatePort_HappyFlow", func(t *testing.T) {
		expectedPort := &srv.Port{Id: "AEDXB", Code: "52000"}
		port, err := h.CreateOrUpdatePort(ctx, expectedPort)
		assert.NotNil(t, port)
		assert.Nil(t, err)
		assert.Equal(t, expectedPort, port)
	})
	t.Run("GetPort_NoId", func(t *testing.T) {
		port, err := h.GetPort(ctx, &srv.Port{})
		assert.Nil(t, port)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrorNotFound.Error())
	})
	t.Run("GetPort_NotFound", func(t *testing.T) {
		expectedPort := &srv.Port{Id: "AEFJR"}
		port, err := h.GetPort(ctx, expectedPort)
		assert.Nil(t, port)
		assert.NotNil(t, err)
		assert.EqualError(t, err, ErrorNotFound.Error())
	})
	t.Run("GetPort_HappyFlow", func(t *testing.T) {
		expectedPort := &srv.Port{Id: "AEDXB"}
		port, err := h.GetPort(ctx, expectedPort)
		assert.NotNil(t, port)
		assert.Nil(t, err)
		assert.Equal(t, expectedPort.Id, port.Id)
		assert.Equal(t, "52000", port.Code)
	})
}
