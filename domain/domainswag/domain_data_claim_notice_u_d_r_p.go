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

// DomainDataClaimNoticeUDRP Definition of a UDRP procedure
// swagger:model domain.Data.ClaimNotice.UDRP
type DomainDataClaimNoticeUDRP struct {

	// Case number
	// Required: true
	// Read Only: true
	CaseNumber string `json:"caseNumber"`

	// UDRP Provider
	// Required: true
	// Read Only: true
	UdrpProvider string `json:"udrpProvider"`
}

// Validate validates this domain data claim notice u d r p
func (m *DomainDataClaimNoticeUDRP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaseNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUdrpProvider(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDataClaimNoticeUDRP) validateCaseNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("caseNumber", "body", string(m.CaseNumber)); err != nil {
		return err
	}

	return nil
}

func (m *DomainDataClaimNoticeUDRP) validateUdrpProvider(formats strfmt.Registry) error {

	if err := validate.RequiredString("udrpProvider", "body", string(m.UdrpProvider)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDataClaimNoticeUDRP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDataClaimNoticeUDRP) UnmarshalBinary(b []byte) error {
	var res DomainDataClaimNoticeUDRP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}