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

// DomainZoneRecord Zone resource records
// swagger:model domain.Zone.Record
type DomainZoneRecord struct {

	// Resource record Name
	// Required: true
	// Read Only: true
	FieldType string `json:"fieldType"`

	// Id of the zone resource record
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Resource record subdomain
	SubDomain string `json:"subDomain,omitempty"`

	// Resource record target
	Target string `json:"target,omitempty"`

	// Resource record ttl
	TTL int64 `json:"ttl,omitempty"`

	// Resource record zone
	// Required: true
	// Read Only: true
	Zone string `json:"zone"`
}

// Validate validates this domain zone record
func (m *DomainZoneRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
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

var domainZoneRecordTypeFieldTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","AAAA","CAA","CNAME","DKIM","LOC","MX","NAPTR","NS","PTR","SPF","SRV","SSHFP","TLSA","TXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainZoneRecordTypeFieldTypePropEnum = append(domainZoneRecordTypeFieldTypePropEnum, v)
	}
}

const (

	// DomainZoneRecordFieldTypeA captures enum value "A"
	DomainZoneRecordFieldTypeA string = "A"

	// DomainZoneRecordFieldTypeAAAA captures enum value "AAAA"
	DomainZoneRecordFieldTypeAAAA string = "AAAA"

	// DomainZoneRecordFieldTypeCAA captures enum value "CAA"
	DomainZoneRecordFieldTypeCAA string = "CAA"

	// DomainZoneRecordFieldTypeCNAME captures enum value "CNAME"
	DomainZoneRecordFieldTypeCNAME string = "CNAME"

	// DomainZoneRecordFieldTypeDKIM captures enum value "DKIM"
	DomainZoneRecordFieldTypeDKIM string = "DKIM"

	// DomainZoneRecordFieldTypeLOC captures enum value "LOC"
	DomainZoneRecordFieldTypeLOC string = "LOC"

	// DomainZoneRecordFieldTypeMX captures enum value "MX"
	DomainZoneRecordFieldTypeMX string = "MX"

	// DomainZoneRecordFieldTypeNAPTR captures enum value "NAPTR"
	DomainZoneRecordFieldTypeNAPTR string = "NAPTR"

	// DomainZoneRecordFieldTypeNS captures enum value "NS"
	DomainZoneRecordFieldTypeNS string = "NS"

	// DomainZoneRecordFieldTypePTR captures enum value "PTR"
	DomainZoneRecordFieldTypePTR string = "PTR"

	// DomainZoneRecordFieldTypeSPF captures enum value "SPF"
	DomainZoneRecordFieldTypeSPF string = "SPF"

	// DomainZoneRecordFieldTypeSRV captures enum value "SRV"
	DomainZoneRecordFieldTypeSRV string = "SRV"

	// DomainZoneRecordFieldTypeSSHFP captures enum value "SSHFP"
	DomainZoneRecordFieldTypeSSHFP string = "SSHFP"

	// DomainZoneRecordFieldTypeTLSA captures enum value "TLSA"
	DomainZoneRecordFieldTypeTLSA string = "TLSA"

	// DomainZoneRecordFieldTypeTXT captures enum value "TXT"
	DomainZoneRecordFieldTypeTXT string = "TXT"
)

// prop value enum
func (m *DomainZoneRecord) validateFieldTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainZoneRecordTypeFieldTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainZoneRecord) validateFieldType(formats strfmt.Registry) error {

	if err := validate.RequiredString("fieldType", "body", string(m.FieldType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateFieldTypeEnum("fieldType", "body", m.FieldType); err != nil {
		return err
	}

	return nil
}

func (m *DomainZoneRecord) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *DomainZoneRecord) validateZone(formats strfmt.Registry) error {

	if err := validate.RequiredString("zone", "body", string(m.Zone)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainZoneRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainZoneRecord) UnmarshalBinary(b []byte) error {
	var res DomainZoneRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
