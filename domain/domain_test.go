package domain

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	ovh "github.com/toorop/go-ovh"
)

func TestGetDomains(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `["domain1.tld", "domain2.tld"]`)
	}))
	domains, err := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL)).GetDomains("")
	assert.NoError(t, err)
	assert.Equal(t, []string{"domain1.tld", "domain2.tld"}, domains)
}

func TestGetZones(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `["domain1.tld", "domain2.tld"]`)
	}))
	zones, err := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL)).GetZones()
	assert.NoError(t, err)
	assert.Equal(t, []string{"domain1.tld", "domain2.tld"}, zones)
}

func TestGetZone(t *testing.T) {
	// {"dnssecSupported":true,"hasDnsAnycast":true,"lastUpdate":"2018-02-19T14:11:26Z","nameServers":["string"]}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"lastUpdate":"2017-09-04T14:27:16+02:00","hasDnsAnycast":false,"nameServers":["ns12.ovh.net","dns12.ovh.net"],"dnssecSupported":true}`)
	}))
	zone, err := New(ovh.NewClient("ovh-eu").WithEndpoint(ts.URL)).GetZone("foo.com")
	assert.NoError(t, err)
	assert.Equal(t, "2017-09-04T14:27:16.000+02:00", zone.LastUpdate.String())
	assert.Equal(t, false, zone.HasDNSAnycast)
	assert.Equal(t, "ns12.ovh.net", zone.NameServers[0])
	log.Println(zone.LastUpdate.String())
}
