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

// PostDomainDataAfnicCorporationTrademarkInformationParamsBody post domain data afnic corporation trademark information params body
// swagger:model postDomainDataAfnicCorporationTrademarkInformationParamsBody
type PostDomainDataAfnicCorporationTrademarkInformationParamsBody struct {

	// contact Id
	// Required: true
	ContactID *int64 `json:"contactId"`

	// inpi number
	// Required: true
	InpiNumber *string `json:"inpiNumber"`

	// inpi trademark owner
	// Required: true
	InpiTrademarkOwner *string `json:"inpiTrademarkOwner"`
}

// Validate validates this post domain data afnic corporation trademark information params body
func (m *PostDomainDataAfnicCorporationTrademarkInformationParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInpiNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInpiTrademarkOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostDomainDataAfnicCorporationTrademarkInformationParamsBody) validateContactID(formats strfmt.Registry) error {

	if err := validate.Required("contactId", "body", m.ContactID); err != nil {
		return err
	}

	return nil
}

func (m *PostDomainDataAfnicCorporationTrademarkInformationParamsBody) validateInpiNumber(formats strfmt.Registry) error {

	if err := validate.Required("inpiNumber", "body", m.InpiNumber); err != nil {
		return err
	}

	return nil
}

func (m *PostDomainDataAfnicCorporationTrademarkInformationParamsBody) validateInpiTrademarkOwner(formats strfmt.Registry) error {

	if err := validate.Required("inpiTrademarkOwner", "body", m.InpiTrademarkOwner); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostDomainDataAfnicCorporationTrademarkInformationParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostDomainDataAfnicCorporationTrademarkInformationParamsBody) UnmarshalBinary(b []byte) error {
	var res PostDomainDataAfnicCorporationTrademarkInformationParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
