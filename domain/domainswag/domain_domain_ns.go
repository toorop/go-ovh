// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DomainDomainNs Name server
// swagger:model domain.DomainNs
type DomainDomainNs struct {

	// Host
	Host string `json:"host,omitempty"`

	// Ip
	IP string `json:"ip,omitempty"`
}

// Validate validates this domain domain ns
func (m *DomainDomainNs) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DomainDomainNs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDomainNs) UnmarshalBinary(b []byte) error {
	var res DomainDomainNs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
