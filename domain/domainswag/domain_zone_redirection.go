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

// DomainZoneRedirection Redirection
// swagger:model domain.Zone.Redirection
type DomainZoneRedirection struct {

	// Desciption for invisible redirection
	Description string `json:"description,omitempty"`

	// Id of the redirection
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Keywords for invisible redirection
	Keywords string `json:"keywords,omitempty"`

	// subdomain to redirect
	// Read Only: true
	SubDomain string `json:"subDomain,omitempty"`

	// Target of the redirection
	Target string `json:"target,omitempty"`

	// Title for invisible redirection
	Title string `json:"title,omitempty"`

	// Redirection type
	Type string `json:"type,omitempty"`

	// Redirection zone
	// Required: true
	// Read Only: true
	Zone string `json:"zone"`
}

// Validate validates this domain zone redirection
func (m *DomainZoneRedirection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *DomainZoneRedirection) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

var domainZoneRedirectionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invisible","visible","visiblePermanent"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainZoneRedirectionTypeTypePropEnum = append(domainZoneRedirectionTypeTypePropEnum, v)
	}
}

const (

	// DomainZoneRedirectionTypeInvisible captures enum value "invisible"
	DomainZoneRedirectionTypeInvisible string = "invisible"

	// DomainZoneRedirectionTypeVisible captures enum value "visible"
	DomainZoneRedirectionTypeVisible string = "visible"

	// DomainZoneRedirectionTypeVisiblePermanent captures enum value "visiblePermanent"
	DomainZoneRedirectionTypeVisiblePermanent string = "visiblePermanent"
)

// prop value enum
func (m *DomainZoneRedirection) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainZoneRedirectionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainZoneRedirection) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DomainZoneRedirection) validateZone(formats strfmt.Registry) error {

	if err := validate.RequiredString("zone", "body", string(m.Zone)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainZoneRedirection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainZoneRedirection) UnmarshalBinary(b []byte) error {
	var res DomainZoneRedirection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
