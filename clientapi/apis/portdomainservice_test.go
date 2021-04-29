package apis

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var portsJson = `
{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  }
}
`

func TestPortDomainServAPI(t *testing.T) {
	ctx := context.Background()
	papi := NewPortDomainServAPI("")

	tmpDestDir, err := ioutil.TempDir("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDestDir)

	tmpSrcFile, err := ioutil.TempFile(tmpDestDir, "ports.json")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpSrcFile.Name())

	_, err = tmpSrcFile.WriteString(portsJson)
	assert.NoError(t, err)

	t.Run("Connect", func(t *testing.T) {
		err := papi.Connect(ctx)
		assert.Nil(t, err)
	})
	t.Run("InitializeDatabase", func(t *testing.T) {
		papi.pdsClient = &MockPDSClient{}
		err := papi.InitializeDatabase(tmpSrcFile.Name())
		assert.NoError(t, err)
	})
}
