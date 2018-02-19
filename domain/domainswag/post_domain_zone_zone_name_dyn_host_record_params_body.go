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

// PostDomainZoneZoneNameDynHostRecordParamsBody post domain zone zone name dyn host record params body
// swagger:model postDomainZoneZoneNameDynHostRecordParamsBody
type PostDomainZoneZoneNameDynHostRecordParamsBody struct {

	// ip
	// Required: true
	IP *string `json:"ip"`

	// sub domain
	SubDomain string `json:"subDomain,omitempty"`
}

// Validate validates this post domain zone zone name dyn host record params body
func (m *PostDomainZoneZoneNameDynHostRecordParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostDomainZoneZoneNameDynHostRecordParamsBody) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostDomainZoneZoneNameDynHostRecordParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostDomainZoneZoneNameDynHostRecordParamsBody) UnmarshalBinary(b []byte) error {
	var res PostDomainZoneZoneNameDynHostRecordParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
