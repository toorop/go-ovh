package ovh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient("ovh-eu")
	assert.Equal(t, c.Endpoint, "https://eu.api.ovh.com/1.0")
	assert.Nil(t, c.Keyring)
}

// TODO test withKeyring
