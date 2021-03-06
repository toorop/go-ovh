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

// DomainZoneDynHostRecord DynHost record
// swagger:model domain.Zone.DynHostRecord
type DomainZoneDynHostRecord struct {

	// Id of the DynHost record
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Ip address of the DynHost record
	IP string `json:"ip,omitempty"`

	// Subdomain of the DynHost record
	SubDomain string `json:"subDomain,omitempty"`

	// Zone of the DynHost record
	// Required: true
	// Read Only: true
	Zone string `json:"zone"`
}

// Validate validates this domain zone dyn host record
func (m *DomainZoneDynHostRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainZoneDynHostRecord) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DomainZoneDynHostRecord) validateZone(formats strfmt.Registry) error {

	if err := validate.RequiredString("zone", "body", string(m.Zone)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainZoneDynHostRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainZoneDynHostRecord) UnmarshalBinary(b []byte) error {
	var res DomainZoneDynHostRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
