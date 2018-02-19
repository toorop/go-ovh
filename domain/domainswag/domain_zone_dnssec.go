// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainZoneDnssec Manage Dnssec for this zone
// swagger:model domain.Zone.Dnssec
type DomainZoneDnssec struct {

	// status
	// Required: true
	// Read Only: true
	Status string `json:"status"`
}

// Validate validates this domain zone dnssec
func (m *DomainZoneDnssec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainZoneDnssecTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disableInProgress","disabled","enableInProgress","enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainZoneDnssecTypeStatusPropEnum = append(domainZoneDnssecTypeStatusPropEnum, v)
	}
}

const (

	// DomainZoneDnssecStatusDisableInProgress captures enum value "disableInProgress"
	DomainZoneDnssecStatusDisableInProgress string = "disableInProgress"

	// DomainZoneDnssecStatusDisabled captures enum value "disabled"
	DomainZoneDnssecStatusDisabled string = "disabled"

	// DomainZoneDnssecStatusEnableInProgress captures enum value "enableInProgress"
	DomainZoneDnssecStatusEnableInProgress string = "enableInProgress"

	// DomainZoneDnssecStatusEnabled captures enum value "enabled"
	DomainZoneDnssecStatusEnabled string = "enabled"
)

// prop value enum
func (m *DomainZoneDnssec) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainZoneDnssecTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainZoneDnssec) validateStatus(formats strfmt.Registry) error {

	if err := validate.RequiredString("status", "body", string(m.Status)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainZoneDnssec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainZoneDnssec) UnmarshalBinary(b []byte) error {
	var res DomainZoneDnssec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
