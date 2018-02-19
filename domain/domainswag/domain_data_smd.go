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

// DomainDataSmd Representation of a SMD Resource file
// swagger:model domain.Data.Smd
type DomainDataSmd struct {

	// SMD file content
	Data string `json:"data,omitempty"`

	// SMD file ID
	ID int64 `json:"id,omitempty"`

	// Date when information about SMD file aren't valid anymore
	NotAfter strfmt.DateTime `json:"notAfter,omitempty"`

	// Date before when information about SMD file aren't valid yet
	NotBefore strfmt.DateTime `json:"notBefore,omitempty"`

	// protected labels
	ProtectedLabels DomainDataSmdProtectedLabels `json:"protectedLabels"`

	// TMCH Internal identifier
	SmdID string `json:"smdId,omitempty"`
}

// Validate validates this domain data smd
func (m *DomainDataSmd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotAfter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNotBefore(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDataSmd) validateNotAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.NotAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("notAfter", "body", "date-time", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainDataSmd) validateNotBefore(formats strfmt.Registry) error {

	if swag.IsZero(m.NotBefore) { // not required
		return nil
	}

	if err := validate.FormatOf("notBefore", "body", "date-time", m.NotBefore.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDataSmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDataSmd) UnmarshalBinary(b []byte) error {
	var res DomainDataSmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}