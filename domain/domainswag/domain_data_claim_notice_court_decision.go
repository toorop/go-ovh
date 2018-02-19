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

// DomainDataClaimNoticeCourtDecision Definition of a court decision
// swagger:model domain.Data.ClaimNotice.CourtDecision
type DomainDataClaimNoticeCourtDecision struct {

	// Country code
	// Required: true
	// Read Only: true
	CountryCode string `json:"countryCode"`

	// Court name
	// Required: true
	// Read Only: true
	CourtName string `json:"courtName"`

	// Reference number of court decision
	// Required: true
	// Read Only: true
	ReferenceNumber string `json:"referenceNumber"`

	// Regions where court decision apply
	// Required: true
	// Read Only: true
	Regions []string `json:"regions"`
}

// Validate validates this domain data claim notice court decision
func (m *DomainDataClaimNoticeCourtDecision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCourtName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReferenceNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDataClaimNoticeCourtDecision) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.RequiredString("countryCode", "body", string(m.CountryCode)); err != nil {
		return err
	}

	return nil
}

func (m *DomainDataClaimNoticeCourtDecision) validateCourtName(formats strfmt.Registry) error {

	if err := validate.RequiredString("courtName", "body", string(m.CourtName)); err != nil {
		return err
	}

	return nil
}

func (m *DomainDataClaimNoticeCourtDecision) validateReferenceNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("referenceNumber", "body", string(m.ReferenceNumber)); err != nil {
		return err
	}

	return nil
}

func (m *DomainDataClaimNoticeCourtDecision) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDataClaimNoticeCourtDecision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDataClaimNoticeCourtDecision) UnmarshalBinary(b []byte) error {
	var res DomainDataClaimNoticeCourtDecision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
