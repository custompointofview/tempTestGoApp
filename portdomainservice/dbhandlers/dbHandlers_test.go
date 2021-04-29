package dbhandlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDBHandler(t *testing.T) {
	t.Run("MEM", func(t *testing.T) {
		h := NewDBHandler(DBHandler_MEM, &DBConfig{})
		assert.NotNil(t, h)
		assert.IsType(t, &MemDBHandler{}, h)
	})
	t.Run("BadType", func(t *testing.T) {
		h := NewDBHandler(3, &DBConfig{})
		assert.Nil(t, h)
	})
}
