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

// PostDomainDataProContactParamsBody post domain data pro contact params body
// swagger:model postDomainDataProContactParamsBody
type PostDomainDataProContactParamsBody struct {

	// authority
	// Required: true
	Authority *string `json:"authority"`

	// authority website
	// Required: true
	AuthorityWebsite *string `json:"authorityWebsite"`

	// contact Id
	ContactID int64 `json:"contactId,omitempty"`

	// job description
	// Required: true
	JobDescription *string `json:"jobDescription"`

	// license number
	// Required: true
	LicenseNumber *string `json:"licenseNumber"`
}

// Validate validates this post domain data pro contact params body
func (m *PostDomainDataProContactParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthority(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuthorityWebsite(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJobDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLicenseNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostDomainDataProContactParamsBody) validateAuthority(formats strfmt.Registry) error {

	if err := validate.Required("authority", "body", m.Authority); err != nil {
		return err
	}

	return nil
}

func (m *PostDomainDataProContactParamsBody) validateAuthorityWebsite(formats strfmt.Registry) error {

	if err := validate.Required("authorityWebsite", "body", m.AuthorityWebsite); err != nil {
		return err
	}

	return nil
}

func (m *PostDomainDataProContactParamsBody) validateJobDescription(formats strfmt.Registry) error {

	if err := validate.Required("jobDescription", "body", m.JobDescription); err != nil {
		return err
	}

	return nil
}

func (m *PostDomainDataProContactParamsBody) validateLicenseNumber(formats strfmt.Registry) error {

	if err := validate.Required("licenseNumber", "body", m.LicenseNumber); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostDomainDataProContactParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostDomainDataProContactParamsBody) UnmarshalBinary(b []byte) error {
	var res PostDomainDataProContactParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}