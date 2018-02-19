// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostDomainServiceNameChangeContactParamsBody post domain service name change contact params body
// swagger:model postDomainServiceNameChangeContactParamsBody
type PostDomainServiceNameChangeContactParamsBody struct {

	// contact admin
	ContactAdmin string `json:"contactAdmin,omitempty"`

	// contact billing
	ContactBilling string `json:"contactBilling,omitempty"`

	// contact tech
	ContactTech string `json:"contactTech,omitempty"`
}

// Validate validates this post domain service name change contact params body
func (m *PostDomainServiceNameChangeContactParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostDomainServiceNameChangeContactParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostDomainServiceNameChangeContactParamsBody) UnmarshalBinary(b []byte) error {
	var res PostDomainServiceNameChangeContactParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}