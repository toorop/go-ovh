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

// DomainDomainNsStatus DNS server status
// swagger:model domain.DomainNsStatus
type DomainDomainNsStatus struct {

	// Whether or not the DNS server is working
	State string `json:"state,omitempty"`

	// Whether or not the DNS server is managed by OVH
	Type string `json:"type,omitempty"`

	// Date from which the DNS server is used by the domain
	UsedSince strfmt.DateTime `json:"usedSince,omitempty"`
}

// Validate validates this domain domain ns status
func (m *DomainDomainNsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsedSince(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainDomainNsStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ko","ok"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainDomainNsStatusTypeStatePropEnum = append(domainDomainNsStatusTypeStatePropEnum, v)
	}
}

const (

	// DomainDomainNsStatusStateKo captures enum value "ko"
	DomainDomainNsStatusStateKo string = "ko"

	// DomainDomainNsStatusStateOk captures enum value "ok"
	DomainDomainNsStatusStateOk string = "ok"
)

// prop value enum
func (m *DomainDomainNsStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainDomainNsStatusTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainDomainNsStatus) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var domainDomainNsStatusTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["external","hosted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainDomainNsStatusTypeTypePropEnum = append(domainDomainNsStatusTypeTypePropEnum, v)
	}
}

const (

	// DomainDomainNsStatusTypeExternal captures enum value "external"
	DomainDomainNsStatusTypeExternal string = "external"

	// DomainDomainNsStatusTypeHosted captures enum value "hosted"
	DomainDomainNsStatusTypeHosted string = "hosted"
)

// prop value enum
func (m *DomainDomainNsStatus) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainDomainNsStatusTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainDomainNsStatus) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DomainDomainNsStatus) validateUsedSince(formats strfmt.Registry) error {

	if swag.IsZero(m.UsedSince) { // not required
		return nil
	}

	if err := validate.FormatOf("usedSince", "body", "date-time", m.UsedSince.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDomainNsStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDomainNsStatus) UnmarshalBinary(b []byte) error {
	var res DomainDomainNsStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
