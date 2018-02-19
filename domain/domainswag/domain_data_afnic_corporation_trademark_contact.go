// Code generated by go-swagger; DO NOT EDIT.

package domainswag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DomainDataAfnicCorporationTrademarkContact Representation of an Inpi additional information for a corporation
// swagger:model domain.Data.AfnicCorporationTrademarkContact
type DomainDataAfnicCorporationTrademarkContact struct {

	// Contact ID related to the Inpi additional information
	ContactID int64 `json:"contactId,omitempty"`

	// Corporation Inpi additional information ID
	ID int64 `json:"id,omitempty"`

	// Number of the Inpi declaration
	InpiNumber string `json:"inpiNumber,omitempty"`

	// Owner of the trademark
	InpiTrademarkOwner string `json:"inpiTrademarkOwner,omitempty"`
}

// Validate validates this domain data afnic corporation trademark contact
func (m *DomainDataAfnicCorporationTrademarkContact) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DomainDataAfnicCorporationTrademarkContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDataAfnicCorporationTrademarkContact) UnmarshalBinary(b []byte) error {
	var res DomainDataAfnicCorporationTrademarkContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}