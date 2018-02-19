package domain

import (
	"encoding/json"
	"net/url"

	ovh "github.com/toorop/go-ovh"
	"github.com/toorop/go-ovh/domain/domainswag"
)

// Domain wrap /domain
type Domain struct {
	*ovh.Client
}

// New return new Auth client
func New(client *ovh.Client) *Domain {
	return &Domain{
		client,
	}
}

// GetDomains return domains with ovh as registar
func (d *Domain) GetDomains(whoisOwner string) ([]string, error) {
	uri := "/domain"
	if whoisOwner != "" {
		uri += "?whoisOwner=" + url.QueryEscape(whoisOwner)
	}
	body, err := d.GET("/domain")
	if err != nil {
		return nil, err
	}
	body = body[:len(body)]
	var domains []string
	err = json.Unmarshal(body, &domains)
	return domains, nil
}

// GetZones retuns zones handled by OVH NS
func (d *Domain) GetZones() ([]string, error) {
	body, err := d.GET("/domain/zone")
	if err != nil {
		return nil, err
	}
	body = body[:len(body)]
	var zones []string
	err = json.Unmarshal(body, &zones)
	return zones, err
}

// GetZone return zone info
func (d *Domain) GetZone(zone string) (*domainswag.DomainZoneZone, error) {
	body, err := d.GET("/domain/zone/" + url.QueryEscape(zone))
	if err != nil {
		return nil, err
	}
	var z domainswag.DomainZoneZone
	err = z.UnmarshalBinary(body)
	return &z, err
}
