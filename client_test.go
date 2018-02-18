package ovh

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient("ovh-eu")
	assert.Equal(t, c.Endpoint, "https://eu.api.ovh.com/1.0")
	assert.Nil(t, c.Keyring)
}

func TestNewClientWithKeyring(t *testing.T) {
	c := NewClient("ovh-eu").WithKeyring("AK", "AS", "CK")
	assert.NotNil(t, c.Keyring)
	assert.Equal(t, "AK", c.Keyring.AK)
	assert.Equal(t, "AS", c.Keyring.AS)
	assert.Equal(t, "CK", c.Keyring.CK)
}

func TestNewClientWithEndpoint(t *testing.T) {
	c := NewClient("ovh-eu").WithEndpoint("endpoint")
	assert.Equal(t, "endpoint", c.Endpoint)
}

func TestGET(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HELLO")
	}))
	body, err := NewClient("ovh-eu").WithEndpoint(ts.URL).GET("/foo")
	assert.NoError(t, err)
	assert.Equal(t, []byte("HELLO\n"), body)
}

func TestPOST(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		defer r.Body.Close()
		fmt.Fprintln(w, string(body))
	}))
	body, err := NewClient("ovh-eu").WithEndpoint(ts.URL).POST("/foo", []byte("HELLO"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("HELLO\n"), body)
}
