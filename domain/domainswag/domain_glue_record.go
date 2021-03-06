// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainGlueRecord Glue record
// swagger:model domain.GlueRecord
type DomainGlueRecord struct {

	// Host of the glue record
	// Required: true
	// Read Only: true
	Host string `json:"host"`

	// Ips of the glue record
	// Required: true
	// Read Only: true
	Ips []string `json:"ips"`
}

// Validate validates this domain glue record
func (m *DomainGlueRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainGlueRecord) validateHost(formats strfmt.Registry) error {

	if err := validate.RequiredString("host", "body", string(m.Host)); err != nil {
		return err
	}

	return nil
}

func (m *DomainGlueRecord) validateIps(formats strfmt.Registry) error {

	if err := validate.Required("ips", "body", m.Ips); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainGlueRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainGlueRecord) UnmarshalBinary(b []byte) error {
	var res DomainGlueRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
