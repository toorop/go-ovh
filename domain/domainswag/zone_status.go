// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ZoneStatus Zone status
// swagger:model zone.Status
type ZoneStatus struct {

	// Error list
	Errors []string `json:"errors"`

	// True if the zone has successfully been deployed
	IsDeployed bool `json:"isDeployed,omitempty"`

	// Warning list
	Warnings []string `json:"warnings"`
}

// Validate validates this zone status
func (m *ZoneStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneStatus) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	return nil
}

func (m *ZoneStatus) validateWarnings(formats strfmt.Registry) error {

	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneStatus) UnmarshalBinary(b []byte) error {
	var res ZoneStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
