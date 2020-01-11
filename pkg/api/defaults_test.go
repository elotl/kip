package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServicePortUnmarshal(t *testing.T) {
	type Contains struct {
		Port ServicePort
	}
	j := `{"Port":{"name":"foo","protocol":"TCP","port":55,"portRangeSize":5}}`
	var c Contains
	err := json.Unmarshal([]byte(j), &c)
	assert.NoError(t, err)
	assert.Equal(t, 5, c.Port.PortRangeSize)
	j = `{"Port":{"name":"foo","protocol":"TCP","port":55}}`
	err = json.Unmarshal([]byte(j), &c)
	assert.NoError(t, err)
	assert.Equal(t, 1, c.Port.PortRangeSize)
}
