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

// DomainDataClaimNoticeAddress Address for a claim notice holder
// swagger:model domain.Data.ClaimNotice.Address
type DomainDataClaimNoticeAddress struct {

	// City
	// Read Only: true
	City string `json:"city,omitempty"`

	// Country code
	// Read Only: true
	CountryCode string `json:"countryCode,omitempty"`

	// Fax number
	// Read Only: true
	Fax string `json:"fax,omitempty"`

	// Fax number extension
	// Read Only: true
	FaxExtension string `json:"faxExtension,omitempty"`

	// Postal zip code
	// Read Only: true
	PostalCode string `json:"postalCode,omitempty"`

	// State of province
	// Read Only: true
	StateOrProvince string `json:"stateOrProvince,omitempty"`

	// Array of street name
	// Required: true
	// Read Only: true
	Streets []string `json:"streets"`

	// Phone number
	// Read Only: true
	Voice string `json:"voice,omitempty"`

	// Phone number extension
	// Read Only: true
	VoiceExtension string `json:"voiceExtension,omitempty"`
}

// Validate validates this domain data claim notice address
func (m *DomainDataClaimNoticeAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountryCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStreets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainDataClaimNoticeAddressTypeCountryCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AC","AD","AE","AF","AG","AI","AL","AM","AO","AQ","AR","AS","AT","AU","AW","AX","AZ","BA","BB","BD","BE","BF","BG","BH","BI","BJ","BL","BM","BN","BO","BQ","BR","BS","BT","BW","BY","BZ","CA","CC","CD","CF","CG","CH","CI","CK","CL","CM","CN","CO","CR","CU","CV","CW","CX","CY","CZ","DE","DG","DJ","DK","DM","DO","DZ","EA","EC","EE","EG","EH","ER","ES","ET","FI","FJ","FK","FM","FO","FR","GA","GB","GD","GE","GF","GG","GH","GI","GL","GM","GN","GP","GQ","GR","GS","GT","GU","GW","GY","HK","HN","HR","HT","HU","IC","ID","IE","IL","IM","IN","IO","IQ","IR","IS","IT","JE","JM","JO","JP","KE","KG","KH","KI","KM","KN","KP","KR","KW","KY","KZ","LA","LB","LC","LI","LK","LR","LS","LT","LU","LV","LY","MA","MC","MD","ME","MF","MG","MH","MK","ML","MM","MN","MO","MP","MQ","MR","MS","MT","MU","MV","MW","MX","MY","MZ","NA","NC","NE","NF","NG","NI","NL","NO","NP","NR","NU","NZ","OM","PA","PE","PF","PG","PH","PK","PL","PM","PN","PR","PS","PT","PW","PY","QA","RE","RO","RS","RU","RW","SA","SB","SC","SD","SE","SG","SH","SI","SJ","SK","SL","SM","SN","SO","SR","SS","ST","SV","SX","SY","SZ","TA","TC","TD","TF","TG","TH","TJ","TK","TL","TM","TN","TO","TR","TT","TV","TW","TZ","UA","UG","UM","UNKNOWN","US","UY","UZ","VA","VC","VE","VG","VI","VN","VU","WF","WS","XK","YE","YT","ZA","ZM","ZW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainDataClaimNoticeAddressTypeCountryCodePropEnum = append(domainDataClaimNoticeAddressTypeCountryCodePropEnum, v)
	}
}

const (

	// DomainDataClaimNoticeAddressCountryCodeAC captures enum value "AC"
	DomainDataClaimNoticeAddressCountryCodeAC string = "AC"

	// DomainDataClaimNoticeAddressCountryCodeAD captures enum value "AD"
	DomainDataClaimNoticeAddressCountryCodeAD string = "AD"

	// DomainDataClaimNoticeAddressCountryCodeAE captures enum value "AE"
	DomainDataClaimNoticeAddressCountryCodeAE string = "AE"

	// DomainDataClaimNoticeAddressCountryCodeAF captures enum value "AF"
	DomainDataClaimNoticeAddressCountryCodeAF string = "AF"

	// DomainDataClaimNoticeAddressCountryCodeAG captures enum value "AG"
	DomainDataClaimNoticeAddressCountryCodeAG string = "AG"

	// DomainDataClaimNoticeAddressCountryCodeAI captures enum value "AI"
	DomainDataClaimNoticeAddressCountryCodeAI string = "AI"

	// DomainDataClaimNoticeAddressCountryCodeAL captures enum value "AL"
	DomainDataClaimNoticeAddressCountryCodeAL string = "AL"

	// DomainDataClaimNoticeAddressCountryCodeAM captures enum value "AM"
	DomainDataClaimNoticeAddressCountryCodeAM string = "AM"

	// DomainDataClaimNoticeAddressCountryCodeAO captures enum value "AO"
	DomainDataClaimNoticeAddressCountryCodeAO string = "AO"

	// DomainDataClaimNoticeAddressCountryCodeAQ captures enum value "AQ"
	DomainDataClaimNoticeAddressCountryCodeAQ string = "AQ"

	// DomainDataClaimNoticeAddressCountryCodeAR captures enum value "AR"
	DomainDataClaimNoticeAddressCountryCodeAR string = "AR"

	// DomainDataClaimNoticeAddressCountryCodeAS captures enum value "AS"
	DomainDataClaimNoticeAddressCountryCodeAS string = "AS"

	// DomainDataClaimNoticeAddressCountryCodeAT captures enum value "AT"
	DomainDataClaimNoticeAddressCountryCodeAT string = "AT"

	// DomainDataClaimNoticeAddressCountryCodeAU captures enum value "AU"
	DomainDataClaimNoticeAddressCountryCodeAU string = "AU"

	// DomainDataClaimNoticeAddressCountryCodeAW captures enum value "AW"
	DomainDataClaimNoticeAddressCountryCodeAW string = "AW"

	// DomainDataClaimNoticeAddressCountryCodeAX captures enum value "AX"
	DomainDataClaimNoticeAddressCountryCodeAX string = "AX"

	// DomainDataClaimNoticeAddressCountryCodeAZ captures enum value "AZ"
	DomainDataClaimNoticeAddressCountryCodeAZ string = "AZ"

	// DomainDataClaimNoticeAddressCountryCodeBA captures enum value "BA"
	DomainDataClaimNoticeAddressCountryCodeBA string = "BA"

	// DomainDataClaimNoticeAddressCountryCodeBB captures enum value "BB"
	DomainDataClaimNoticeAddressCountryCodeBB string = "BB"

	// DomainDataClaimNoticeAddressCountryCodeBD captures enum value "BD"
	DomainDataClaimNoticeAddressCountryCodeBD string = "BD"

	// DomainDataClaimNoticeAddressCountryCodeBE captures enum value "BE"
	DomainDataClaimNoticeAddressCountryCodeBE string = "BE"

	// DomainDataClaimNoticeAddressCountryCodeBF captures enum value "BF"
	DomainDataClaimNoticeAddressCountryCodeBF string = "BF"

	// DomainDataClaimNoticeAddressCountryCodeBG captures enum value "BG"
	DomainDataClaimNoticeAddressCountryCodeBG string = "BG"

	// DomainDataClaimNoticeAddressCountryCodeBH captures enum value "BH"
	DomainDataClaimNoticeAddressCountryCodeBH string = "BH"

	// DomainDataClaimNoticeAddressCountryCodeBI captures enum value "BI"
	DomainDataClaimNoticeAddressCountryCodeBI string = "BI"

	// DomainDataClaimNoticeAddressCountryCodeBJ captures enum value "BJ"
	DomainDataClaimNoticeAddressCountryCodeBJ string = "BJ"

	// DomainDataClaimNoticeAddressCountryCodeBL captures enum value "BL"
	DomainDataClaimNoticeAddressCountryCodeBL string = "BL"

	// DomainDataClaimNoticeAddressCountryCodeBM captures enum value "BM"
	DomainDataClaimNoticeAddressCountryCodeBM string = "BM"

	// DomainDataClaimNoticeAddressCountryCodeBN captures enum value "BN"
	DomainDataClaimNoticeAddressCountryCodeBN string = "BN"

	// DomainDataClaimNoticeAddressCountryCodeBO captures enum value "BO"
	DomainDataClaimNoticeAddressCountryCodeBO string = "BO"

	// DomainDataClaimNoticeAddressCountryCodeBQ captures enum value "BQ"
	DomainDataClaimNoticeAddressCountryCodeBQ string = "BQ"

	// DomainDataClaimNoticeAddressCountryCodeBR captures enum value "BR"
	DomainDataClaimNoticeAddressCountryCodeBR string = "BR"

	// DomainDataClaimNoticeAddressCountryCodeBS captures enum value "BS"
	DomainDataClaimNoticeAddressCountryCodeBS string = "BS"

	// DomainDataClaimNoticeAddressCountryCodeBT captures enum value "BT"
	DomainDataClaimNoticeAddressCountryCodeBT string = "BT"

	// DomainDataClaimNoticeAddressCountryCodeBW captures enum value "BW"
	DomainDataClaimNoticeAddressCountryCodeBW string = "BW"

	// DomainDataClaimNoticeAddressCountryCodeBY captures enum value "BY"
	DomainDataClaimNoticeAddressCountryCodeBY string = "BY"

	// DomainDataClaimNoticeAddressCountryCodeBZ captures enum value "BZ"
	DomainDataClaimNoticeAddressCountryCodeBZ string = "BZ"

	// DomainDataClaimNoticeAddressCountryCodeCA captures enum value "CA"
	DomainDataClaimNoticeAddressCountryCodeCA string = "CA"

	// DomainDataClaimNoticeAddressCountryCodeCC captures enum value "CC"
	DomainDataClaimNoticeAddressCountryCodeCC string = "CC"

	// DomainDataClaimNoticeAddressCountryCodeCD captures enum value "CD"
	DomainDataClaimNoticeAddressCountryCodeCD string = "CD"

	// DomainDataClaimNoticeAddressCountryCodeCF captures enum value "CF"
	DomainDataClaimNoticeAddressCountryCodeCF string = "CF"

	// DomainDataClaimNoticeAddressCountryCodeCG captures enum value "CG"
	DomainDataClaimNoticeAddressCountryCodeCG string = "CG"

	// DomainDataClaimNoticeAddressCountryCodeCH captures enum value "CH"
	DomainDataClaimNoticeAddressCountryCodeCH string = "CH"

	// DomainDataClaimNoticeAddressCountryCodeCI captures enum value "CI"
	DomainDataClaimNoticeAddressCountryCodeCI string = "CI"

	// DomainDataClaimNoticeAddressCountryCodeCK captures enum value "CK"
	DomainDataClaimNoticeAddressCountryCodeCK string = "CK"

	// DomainDataClaimNoticeAddressCountryCodeCL captures enum value "CL"
	DomainDataClaimNoticeAddressCountryCodeCL string = "CL"

	// DomainDataClaimNoticeAddressCountryCodeCM captures enum value "CM"
	DomainDataClaimNoticeAddressCountryCodeCM string = "CM"

	// DomainDataClaimNoticeAddressCountryCodeCN captures enum value "CN"
	DomainDataClaimNoticeAddressCountryCodeCN string = "CN"

	// DomainDataClaimNoticeAddressCountryCodeCO captures enum value "CO"
	DomainDataClaimNoticeAddressCountryCodeCO string = "CO"

	// DomainDataClaimNoticeAddressCountryCodeCR captures enum value "CR"
	DomainDataClaimNoticeAddressCountryCodeCR string = "CR"

	// DomainDataClaimNoticeAddressCountryCodeCU captures enum value "CU"
	DomainDataClaimNoticeAddressCountryCodeCU string = "CU"

	// DomainDataClaimNoticeAddressCountryCodeCV captures enum value "CV"
	DomainDataClaimNoticeAddressCountryCodeCV string = "CV"

	// DomainDataClaimNoticeAddressCountryCodeCW captures enum value "CW"
	DomainDataClaimNoticeAddressCountryCodeCW string = "CW"

	// DomainDataClaimNoticeAddressCountryCodeCX captures enum value "CX"
	DomainDataClaimNoticeAddressCountryCodeCX string = "CX"

	// DomainDataClaimNoticeAddressCountryCodeCY captures enum value "CY"
	DomainDataClaimNoticeAddressCountryCodeCY string = "CY"

	// DomainDataClaimNoticeAddressCountryCodeCZ captures enum value "CZ"
	DomainDataClaimNoticeAddressCountryCodeCZ string = "CZ"

	// DomainDataClaimNoticeAddressCountryCodeDE captures enum value "DE"
	DomainDataClaimNoticeAddressCountryCodeDE string = "DE"

	// DomainDataClaimNoticeAddressCountryCodeDG captures enum value "DG"
	DomainDataClaimNoticeAddressCountryCodeDG string = "DG"

	// DomainDataClaimNoticeAddressCountryCodeDJ captures enum value "DJ"
	DomainDataClaimNoticeAddressCountryCodeDJ string = "DJ"

	// DomainDataClaimNoticeAddressCountryCodeDK captures enum value "DK"
	DomainDataClaimNoticeAddressCountryCodeDK string = "DK"

	// DomainDataClaimNoticeAddressCountryCodeDM captures enum value "DM"
	DomainDataClaimNoticeAddressCountryCodeDM string = "DM"

	// DomainDataClaimNoticeAddressCountryCodeDO captures enum value "DO"
	DomainDataClaimNoticeAddressCountryCodeDO string = "DO"

	// DomainDataClaimNoticeAddressCountryCodeDZ captures enum value "DZ"
	DomainDataClaimNoticeAddressCountryCodeDZ string = "DZ"

	// DomainDataClaimNoticeAddressCountryCodeEA captures enum value "EA"
	DomainDataClaimNoticeAddressCountryCodeEA string = "EA"

	// DomainDataClaimNoticeAddressCountryCodeEC captures enum value "EC"
	DomainDataClaimNoticeAddressCountryCodeEC string = "EC"

	// DomainDataClaimNoticeAddressCountryCodeEE captures enum value "EE"
	DomainDataClaimNoticeAddressCountryCodeEE string = "EE"

	// DomainDataClaimNoticeAddressCountryCodeEG captures enum value "EG"
	DomainDataClaimNoticeAddressCountryCodeEG string = "EG"

	// DomainDataClaimNoticeAddressCountryCodeEH captures enum value "EH"
	DomainDataClaimNoticeAddressCountryCodeEH string = "EH"

	// DomainDataClaimNoticeAddressCountryCodeER captures enum value "ER"
	DomainDataClaimNoticeAddressCountryCodeER string = "ER"

	// DomainDataClaimNoticeAddressCountryCodeES captures enum value "ES"
	DomainDataClaimNoticeAddressCountryCodeES string = "ES"

	// DomainDataClaimNoticeAddressCountryCodeET captures enum value "ET"
	DomainDataClaimNoticeAddressCountryCodeET string = "ET"

	// DomainDataClaimNoticeAddressCountryCodeFI captures enum value "FI"
	DomainDataClaimNoticeAddressCountryCodeFI string = "FI"

	// DomainDataClaimNoticeAddressCountryCodeFJ captures enum value "FJ"
	DomainDataClaimNoticeAddressCountryCodeFJ string = "FJ"

	// DomainDataClaimNoticeAddressCountryCodeFK captures enum value "FK"
	DomainDataClaimNoticeAddressCountryCodeFK string = "FK"

	// DomainDataClaimNoticeAddressCountryCodeFM captures enum value "FM"
	DomainDataClaimNoticeAddressCountryCodeFM string = "FM"

	// DomainDataClaimNoticeAddressCountryCodeFO captures enum value "FO"
	DomainDataClaimNoticeAddressCountryCodeFO string = "FO"

	// DomainDataClaimNoticeAddressCountryCodeFR captures enum value "FR"
	DomainDataClaimNoticeAddressCountryCodeFR string = "FR"

	// DomainDataClaimNoticeAddressCountryCodeGA captures enum value "GA"
	DomainDataClaimNoticeAddressCountryCodeGA string = "GA"

	// DomainDataClaimNoticeAddressCountryCodeGB captures enum value "GB"
	DomainDataClaimNoticeAddressCountryCodeGB string = "GB"

	// DomainDataClaimNoticeAddressCountryCodeGD captures enum value "GD"
	DomainDataClaimNoticeAddressCountryCodeGD string = "GD"

	// DomainDataClaimNoticeAddressCountryCodeGE captures enum value "GE"
	DomainDataClaimNoticeAddressCountryCodeGE string = "GE"

	// DomainDataClaimNoticeAddressCountryCodeGF captures enum value "GF"
	DomainDataClaimNoticeAddressCountryCodeGF string = "GF"

	// DomainDataClaimNoticeAddressCountryCodeGG captures enum value "GG"
	DomainDataClaimNoticeAddressCountryCodeGG string = "GG"

	// DomainDataClaimNoticeAddressCountryCodeGH captures enum value "GH"
	DomainDataClaimNoticeAddressCountryCodeGH string = "GH"

	// DomainDataClaimNoticeAddressCountryCodeGI captures enum value "GI"
	DomainDataClaimNoticeAddressCountryCodeGI string = "GI"

	// DomainDataClaimNoticeAddressCountryCodeGL captures enum value "GL"
	DomainDataClaimNoticeAddressCountryCodeGL string = "GL"

	// DomainDataClaimNoticeAddressCountryCodeGM captures enum value "GM"
	DomainDataClaimNoticeAddressCountryCodeGM string = "GM"

	// DomainDataClaimNoticeAddressCountryCodeGN captures enum value "GN"
	DomainDataClaimNoticeAddressCountryCodeGN string = "GN"

	// DomainDataClaimNoticeAddressCountryCodeGP captures enum value "GP"
	DomainDataClaimNoticeAddressCountryCodeGP string = "GP"

	// DomainDataClaimNoticeAddressCountryCodeGQ captures enum value "GQ"
	DomainDataClaimNoticeAddressCountryCodeGQ string = "GQ"

	// DomainDataClaimNoticeAddressCountryCodeGR captures enum value "GR"
	DomainDataClaimNoticeAddressCountryCodeGR string = "GR"

	// DomainDataClaimNoticeAddressCountryCodeGS captures enum value "GS"
	DomainDataClaimNoticeAddressCountryCodeGS string = "GS"

	// DomainDataClaimNoticeAddressCountryCodeGT captures enum value "GT"
	DomainDataClaimNoticeAddressCountryCodeGT string = "GT"

	// DomainDataClaimNoticeAddressCountryCodeGU captures enum value "GU"
	DomainDataClaimNoticeAddressCountryCodeGU string = "GU"

	// DomainDataClaimNoticeAddressCountryCodeGW captures enum value "GW"
	DomainDataClaimNoticeAddressCountryCodeGW string = "GW"

	// DomainDataClaimNoticeAddressCountryCodeGY captures enum value "GY"
	DomainDataClaimNoticeAddressCountryCodeGY string = "GY"

	// DomainDataClaimNoticeAddressCountryCodeHK captures enum value "HK"
	DomainDataClaimNoticeAddressCountryCodeHK string = "HK"

	// DomainDataClaimNoticeAddressCountryCodeHN captures enum value "HN"
	DomainDataClaimNoticeAddressCountryCodeHN string = "HN"

	// DomainDataClaimNoticeAddressCountryCodeHR captures enum value "HR"
	DomainDataClaimNoticeAddressCountryCodeHR string = "HR"

	// DomainDataClaimNoticeAddressCountryCodeHT captures enum value "HT"
	DomainDataClaimNoticeAddressCountryCodeHT string = "HT"

	// DomainDataClaimNoticeAddressCountryCodeHU captures enum value "HU"
	DomainDataClaimNoticeAddressCountryCodeHU string = "HU"

	// DomainDataClaimNoticeAddressCountryCodeIC captures enum value "IC"
	DomainDataClaimNoticeAddressCountryCodeIC string = "IC"

	// DomainDataClaimNoticeAddressCountryCodeID captures enum value "ID"
	DomainDataClaimNoticeAddressCountryCodeID string = "ID"

	// DomainDataClaimNoticeAddressCountryCodeIE captures enum value "IE"
	DomainDataClaimNoticeAddressCountryCodeIE string = "IE"

	// DomainDataClaimNoticeAddressCountryCodeIL captures enum value "IL"
	DomainDataClaimNoticeAddressCountryCodeIL string = "IL"

	// DomainDataClaimNoticeAddressCountryCodeIM captures enum value "IM"
	DomainDataClaimNoticeAddressCountryCodeIM string = "IM"

	// DomainDataClaimNoticeAddressCountryCodeIN captures enum value "IN"
	DomainDataClaimNoticeAddressCountryCodeIN string = "IN"

	// DomainDataClaimNoticeAddressCountryCodeIO captures enum value "IO"
	DomainDataClaimNoticeAddressCountryCodeIO string = "IO"

	// DomainDataClaimNoticeAddressCountryCodeIQ captures enum value "IQ"
	DomainDataClaimNoticeAddressCountryCodeIQ string = "IQ"

	// DomainDataClaimNoticeAddressCountryCodeIR captures enum value "IR"
	DomainDataClaimNoticeAddressCountryCodeIR string = "IR"

	// DomainDataClaimNoticeAddressCountryCodeIS captures enum value "IS"
	DomainDataClaimNoticeAddressCountryCodeIS string = "IS"

	// DomainDataClaimNoticeAddressCountryCodeIT captures enum value "IT"
	DomainDataClaimNoticeAddressCountryCodeIT string = "IT"

	// DomainDataClaimNoticeAddressCountryCodeJE captures enum value "JE"
	DomainDataClaimNoticeAddressCountryCodeJE string = "JE"

	// DomainDataClaimNoticeAddressCountryCodeJM captures enum value "JM"
	DomainDataClaimNoticeAddressCountryCodeJM string = "JM"

	// DomainDataClaimNoticeAddressCountryCodeJO captures enum value "JO"
	DomainDataClaimNoticeAddressCountryCodeJO string = "JO"

	// DomainDataClaimNoticeAddressCountryCodeJP captures enum value "JP"
	DomainDataClaimNoticeAddressCountryCodeJP string = "JP"

	// DomainDataClaimNoticeAddressCountryCodeKE captures enum value "KE"
	DomainDataClaimNoticeAddressCountryCodeKE string = "KE"

	// DomainDataClaimNoticeAddressCountryCodeKG captures enum value "KG"
	DomainDataClaimNoticeAddressCountryCodeKG string = "KG"

	// DomainDataClaimNoticeAddressCountryCodeKH captures enum value "KH"
	DomainDataClaimNoticeAddressCountryCodeKH string = "KH"

	// DomainDataClaimNoticeAddressCountryCodeKI captures enum value "KI"
	DomainDataClaimNoticeAddressCountryCodeKI string = "KI"

	// DomainDataClaimNoticeAddressCountryCodeKM captures enum value "KM"
	DomainDataClaimNoticeAddressCountryCodeKM string = "KM"

	// DomainDataClaimNoticeAddressCountryCodeKN captures enum value "KN"
	DomainDataClaimNoticeAddressCountryCodeKN string = "KN"

	// DomainDataClaimNoticeAddressCountryCodeKP captures enum value "KP"
	DomainDataClaimNoticeAddressCountryCodeKP string = "KP"

	// DomainDataClaimNoticeAddressCountryCodeKR captures enum value "KR"
	DomainDataClaimNoticeAddressCountryCodeKR string = "KR"

	// DomainDataClaimNoticeAddressCountryCodeKW captures enum value "KW"
	DomainDataClaimNoticeAddressCountryCodeKW string = "KW"

	// DomainDataClaimNoticeAddressCountryCodeKY captures enum value "KY"
	DomainDataClaimNoticeAddressCountryCodeKY string = "KY"

	// DomainDataClaimNoticeAddressCountryCodeKZ captures enum value "KZ"
	DomainDataClaimNoticeAddressCountryCodeKZ string = "KZ"

	// DomainDataClaimNoticeAddressCountryCodeLA captures enum value "LA"
	DomainDataClaimNoticeAddressCountryCodeLA string = "LA"

	// DomainDataClaimNoticeAddressCountryCodeLB captures enum value "LB"
	DomainDataClaimNoticeAddressCountryCodeLB string = "LB"

	// DomainDataClaimNoticeAddressCountryCodeLC captures enum value "LC"
	DomainDataClaimNoticeAddressCountryCodeLC string = "LC"

	// DomainDataClaimNoticeAddressCountryCodeLI captures enum value "LI"
	DomainDataClaimNoticeAddressCountryCodeLI string = "LI"

	// DomainDataClaimNoticeAddressCountryCodeLK captures enum value "LK"
	DomainDataClaimNoticeAddressCountryCodeLK string = "LK"

	// DomainDataClaimNoticeAddressCountryCodeLR captures enum value "LR"
	DomainDataClaimNoticeAddressCountryCodeLR string = "LR"

	// DomainDataClaimNoticeAddressCountryCodeLS captures enum value "LS"
	DomainDataClaimNoticeAddressCountryCodeLS string = "LS"

	// DomainDataClaimNoticeAddressCountryCodeLT captures enum value "LT"
	DomainDataClaimNoticeAddressCountryCodeLT string = "LT"

	// DomainDataClaimNoticeAddressCountryCodeLU captures enum value "LU"
	DomainDataClaimNoticeAddressCountryCodeLU string = "LU"

	// DomainDataClaimNoticeAddressCountryCodeLV captures enum value "LV"
	DomainDataClaimNoticeAddressCountryCodeLV string = "LV"

	// DomainDataClaimNoticeAddressCountryCodeLY captures enum value "LY"
	DomainDataClaimNoticeAddressCountryCodeLY string = "LY"

	// DomainDataClaimNoticeAddressCountryCodeMA captures enum value "MA"
	DomainDataClaimNoticeAddressCountryCodeMA string = "MA"

	// DomainDataClaimNoticeAddressCountryCodeMC captures enum value "MC"
	DomainDataClaimNoticeAddressCountryCodeMC string = "MC"

	// DomainDataClaimNoticeAddressCountryCodeMD captures enum value "MD"
	DomainDataClaimNoticeAddressCountryCodeMD string = "MD"

	// DomainDataClaimNoticeAddressCountryCodeME captures enum value "ME"
	DomainDataClaimNoticeAddressCountryCodeME string = "ME"

	// DomainDataClaimNoticeAddressCountryCodeMF captures enum value "MF"
	DomainDataClaimNoticeAddressCountryCodeMF string = "MF"

	// DomainDataClaimNoticeAddressCountryCodeMG captures enum value "MG"
	DomainDataClaimNoticeAddressCountryCodeMG string = "MG"

	// DomainDataClaimNoticeAddressCountryCodeMH captures enum value "MH"
	DomainDataClaimNoticeAddressCountryCodeMH string = "MH"

	// DomainDataClaimNoticeAddressCountryCodeMK captures enum value "MK"
	DomainDataClaimNoticeAddressCountryCodeMK string = "MK"

	// DomainDataClaimNoticeAddressCountryCodeML captures enum value "ML"
	DomainDataClaimNoticeAddressCountryCodeML string = "ML"

	// DomainDataClaimNoticeAddressCountryCodeMM captures enum value "MM"
	DomainDataClaimNoticeAddressCountryCodeMM string = "MM"

	// DomainDataClaimNoticeAddressCountryCodeMN captures enum value "MN"
	DomainDataClaimNoticeAddressCountryCodeMN string = "MN"

	// DomainDataClaimNoticeAddressCountryCodeMO captures enum value "MO"
	DomainDataClaimNoticeAddressCountryCodeMO string = "MO"

	// DomainDataClaimNoticeAddressCountryCodeMP captures enum value "MP"
	DomainDataClaimNoticeAddressCountryCodeMP string = "MP"

	// DomainDataClaimNoticeAddressCountryCodeMQ captures enum value "MQ"
	DomainDataClaimNoticeAddressCountryCodeMQ string = "MQ"

	// DomainDataClaimNoticeAddressCountryCodeMR captures enum value "MR"
	DomainDataClaimNoticeAddressCountryCodeMR string = "MR"

	// DomainDataClaimNoticeAddressCountryCodeMS captures enum value "MS"
	DomainDataClaimNoticeAddressCountryCodeMS string = "MS"

	// DomainDataClaimNoticeAddressCountryCodeMT captures enum value "MT"
	DomainDataClaimNoticeAddressCountryCodeMT string = "MT"

	// DomainDataClaimNoticeAddressCountryCodeMU captures enum value "MU"
	DomainDataClaimNoticeAddressCountryCodeMU string = "MU"

	// DomainDataClaimNoticeAddressCountryCodeMV captures enum value "MV"
	DomainDataClaimNoticeAddressCountryCodeMV string = "MV"

	// DomainDataClaimNoticeAddressCountryCodeMW captures enum value "MW"
	DomainDataClaimNoticeAddressCountryCodeMW string = "MW"

	// DomainDataClaimNoticeAddressCountryCodeMX captures enum value "MX"
	DomainDataClaimNoticeAddressCountryCodeMX string = "MX"

	// DomainDataClaimNoticeAddressCountryCodeMY captures enum value "MY"
	DomainDataClaimNoticeAddressCountryCodeMY string = "MY"

	// DomainDataClaimNoticeAddressCountryCodeMZ captures enum value "MZ"
	DomainDataClaimNoticeAddressCountryCodeMZ string = "MZ"

	// DomainDataClaimNoticeAddressCountryCodeNA captures enum value "NA"
	DomainDataClaimNoticeAddressCountryCodeNA string = "NA"

	// DomainDataClaimNoticeAddressCountryCodeNC captures enum value "NC"
	DomainDataClaimNoticeAddressCountryCodeNC string = "NC"

	// DomainDataClaimNoticeAddressCountryCodeNE captures enum value "NE"
	DomainDataClaimNoticeAddressCountryCodeNE string = "NE"

	// DomainDataClaimNoticeAddressCountryCodeNF captures enum value "NF"
	DomainDataClaimNoticeAddressCountryCodeNF string = "NF"

	// DomainDataClaimNoticeAddressCountryCodeNG captures enum value "NG"
	DomainDataClaimNoticeAddressCountryCodeNG string = "NG"

	// DomainDataClaimNoticeAddressCountryCodeNI captures enum value "NI"
	DomainDataClaimNoticeAddressCountryCodeNI string = "NI"

	// DomainDataClaimNoticeAddressCountryCodeNL captures enum value "NL"
	DomainDataClaimNoticeAddressCountryCodeNL string = "NL"

	// DomainDataClaimNoticeAddressCountryCodeNO captures enum value "NO"
	DomainDataClaimNoticeAddressCountryCodeNO string = "NO"

	// DomainDataClaimNoticeAddressCountryCodeNP captures enum value "NP"
	DomainDataClaimNoticeAddressCountryCodeNP string = "NP"

	// DomainDataClaimNoticeAddressCountryCodeNR captures enum value "NR"
	DomainDataClaimNoticeAddressCountryCodeNR string = "NR"

	// DomainDataClaimNoticeAddressCountryCodeNU captures enum value "NU"
	DomainDataClaimNoticeAddressCountryCodeNU string = "NU"

	// DomainDataClaimNoticeAddressCountryCodeNZ captures enum value "NZ"
	DomainDataClaimNoticeAddressCountryCodeNZ string = "NZ"

	// DomainDataClaimNoticeAddressCountryCodeOM captures enum value "OM"
	DomainDataClaimNoticeAddressCountryCodeOM string = "OM"

	// DomainDataClaimNoticeAddressCountryCodePA captures enum value "PA"
	DomainDataClaimNoticeAddressCountryCodePA string = "PA"

	// DomainDataClaimNoticeAddressCountryCodePE captures enum value "PE"
	DomainDataClaimNoticeAddressCountryCodePE string = "PE"

	// DomainDataClaimNoticeAddressCountryCodePF captures enum value "PF"
	DomainDataClaimNoticeAddressCountryCodePF string = "PF"

	// DomainDataClaimNoticeAddressCountryCodePG captures enum value "PG"
	DomainDataClaimNoticeAddressCountryCodePG string = "PG"

	// DomainDataClaimNoticeAddressCountryCodePH captures enum value "PH"
	DomainDataClaimNoticeAddressCountryCodePH string = "PH"

	// DomainDataClaimNoticeAddressCountryCodePK captures enum value "PK"
	DomainDataClaimNoticeAddressCountryCodePK string = "PK"

	// DomainDataClaimNoticeAddressCountryCodePL captures enum value "PL"
	DomainDataClaimNoticeAddressCountryCodePL string = "PL"

	// DomainDataClaimNoticeAddressCountryCodePM captures enum value "PM"
	DomainDataClaimNoticeAddressCountryCodePM string = "PM"

	// DomainDataClaimNoticeAddressCountryCodePN captures enum value "PN"
	DomainDataClaimNoticeAddressCountryCodePN string = "PN"

	// DomainDataClaimNoticeAddressCountryCodePR captures enum value "PR"
	DomainDataClaimNoticeAddressCountryCodePR string = "PR"

	// DomainDataClaimNoticeAddressCountryCodePS captures enum value "PS"
	DomainDataClaimNoticeAddressCountryCodePS string = "PS"

	// DomainDataClaimNoticeAddressCountryCodePT captures enum value "PT"
	DomainDataClaimNoticeAddressCountryCodePT string = "PT"

	// DomainDataClaimNoticeAddressCountryCodePW captures enum value "PW"
	DomainDataClaimNoticeAddressCountryCodePW string = "PW"

	// DomainDataClaimNoticeAddressCountryCodePY captures enum value "PY"
	DomainDataClaimNoticeAddressCountryCodePY string = "PY"

	// DomainDataClaimNoticeAddressCountryCodeQA captures enum value "QA"
	DomainDataClaimNoticeAddressCountryCodeQA string = "QA"

	// DomainDataClaimNoticeAddressCountryCodeRE captures enum value "RE"
	DomainDataClaimNoticeAddressCountryCodeRE string = "RE"

	// DomainDataClaimNoticeAddressCountryCodeRO captures enum value "RO"
	DomainDataClaimNoticeAddressCountryCodeRO string = "RO"

	// DomainDataClaimNoticeAddressCountryCodeRS captures enum value "RS"
	DomainDataClaimNoticeAddressCountryCodeRS string = "RS"

	// DomainDataClaimNoticeAddressCountryCodeRU captures enum value "RU"
	DomainDataClaimNoticeAddressCountryCodeRU string = "RU"

	// DomainDataClaimNoticeAddressCountryCodeRW captures enum value "RW"
	DomainDataClaimNoticeAddressCountryCodeRW string = "RW"

	// DomainDataClaimNoticeAddressCountryCodeSA captures enum value "SA"
	DomainDataClaimNoticeAddressCountryCodeSA string = "SA"

	// DomainDataClaimNoticeAddressCountryCodeSB captures enum value "SB"
	DomainDataClaimNoticeAddressCountryCodeSB string = "SB"

	// DomainDataClaimNoticeAddressCountryCodeSC captures enum value "SC"
	DomainDataClaimNoticeAddressCountryCodeSC string = "SC"

	// DomainDataClaimNoticeAddressCountryCodeSD captures enum value "SD"
	DomainDataClaimNoticeAddressCountryCodeSD string = "SD"

	// DomainDataClaimNoticeAddressCountryCodeSE captures enum value "SE"
	DomainDataClaimNoticeAddressCountryCodeSE string = "SE"

	// DomainDataClaimNoticeAddressCountryCodeSG captures enum value "SG"
	DomainDataClaimNoticeAddressCountryCodeSG string = "SG"

	// DomainDataClaimNoticeAddressCountryCodeSH captures enum value "SH"
	DomainDataClaimNoticeAddressCountryCodeSH string = "SH"

	// DomainDataClaimNoticeAddressCountryCodeSI captures enum value "SI"
	DomainDataClaimNoticeAddressCountryCodeSI string = "SI"

	// DomainDataClaimNoticeAddressCountryCodeSJ captures enum value "SJ"
	DomainDataClaimNoticeAddressCountryCodeSJ string = "SJ"

	// DomainDataClaimNoticeAddressCountryCodeSK captures enum value "SK"
	DomainDataClaimNoticeAddressCountryCodeSK string = "SK"

	// DomainDataClaimNoticeAddressCountryCodeSL captures enum value "SL"
	DomainDataClaimNoticeAddressCountryCodeSL string = "SL"

	// DomainDataClaimNoticeAddressCountryCodeSM captures enum value "SM"
	DomainDataClaimNoticeAddressCountryCodeSM string = "SM"

	// DomainDataClaimNoticeAddressCountryCodeSN captures enum value "SN"
	DomainDataClaimNoticeAddressCountryCodeSN string = "SN"

	// DomainDataClaimNoticeAddressCountryCodeSO captures enum value "SO"
	DomainDataClaimNoticeAddressCountryCodeSO string = "SO"

	// DomainDataClaimNoticeAddressCountryCodeSR captures enum value "SR"
	DomainDataClaimNoticeAddressCountryCodeSR string = "SR"

	// DomainDataClaimNoticeAddressCountryCodeSS captures enum value "SS"
	DomainDataClaimNoticeAddressCountryCodeSS string = "SS"

	// DomainDataClaimNoticeAddressCountryCodeST captures enum value "ST"
	DomainDataClaimNoticeAddressCountryCodeST string = "ST"

	// DomainDataClaimNoticeAddressCountryCodeSV captures enum value "SV"
	DomainDataClaimNoticeAddressCountryCodeSV string = "SV"

	// DomainDataClaimNoticeAddressCountryCodeSX captures enum value "SX"
	DomainDataClaimNoticeAddressCountryCodeSX string = "SX"

	// DomainDataClaimNoticeAddressCountryCodeSY captures enum value "SY"
	DomainDataClaimNoticeAddressCountryCodeSY string = "SY"

	// DomainDataClaimNoticeAddressCountryCodeSZ captures enum value "SZ"
	DomainDataClaimNoticeAddressCountryCodeSZ string = "SZ"

	// DomainDataClaimNoticeAddressCountryCodeTA captures enum value "TA"
	DomainDataClaimNoticeAddressCountryCodeTA string = "TA"

	// DomainDataClaimNoticeAddressCountryCodeTC captures enum value "TC"
	DomainDataClaimNoticeAddressCountryCodeTC string = "TC"

	// DomainDataClaimNoticeAddressCountryCodeTD captures enum value "TD"
	DomainDataClaimNoticeAddressCountryCodeTD string = "TD"

	// DomainDataClaimNoticeAddressCountryCodeTF captures enum value "TF"
	DomainDataClaimNoticeAddressCountryCodeTF string = "TF"

	// DomainDataClaimNoticeAddressCountryCodeTG captures enum value "TG"
	DomainDataClaimNoticeAddressCountryCodeTG string = "TG"

	// DomainDataClaimNoticeAddressCountryCodeTH captures enum value "TH"
	DomainDataClaimNoticeAddressCountryCodeTH string = "TH"

	// DomainDataClaimNoticeAddressCountryCodeTJ captures enum value "TJ"
	DomainDataClaimNoticeAddressCountryCodeTJ string = "TJ"

	// DomainDataClaimNoticeAddressCountryCodeTK captures enum value "TK"
	DomainDataClaimNoticeAddressCountryCodeTK string = "TK"

	// DomainDataClaimNoticeAddressCountryCodeTL captures enum value "TL"
	DomainDataClaimNoticeAddressCountryCodeTL string = "TL"

	// DomainDataClaimNoticeAddressCountryCodeTM captures enum value "TM"
	DomainDataClaimNoticeAddressCountryCodeTM string = "TM"

	// DomainDataClaimNoticeAddressCountryCodeTN captures enum value "TN"
	DomainDataClaimNoticeAddressCountryCodeTN string = "TN"

	// DomainDataClaimNoticeAddressCountryCodeTO captures enum value "TO"
	DomainDataClaimNoticeAddressCountryCodeTO string = "TO"

	// DomainDataClaimNoticeAddressCountryCodeTR captures enum value "TR"
	DomainDataClaimNoticeAddressCountryCodeTR string = "TR"

	// DomainDataClaimNoticeAddressCountryCodeTT captures enum value "TT"
	DomainDataClaimNoticeAddressCountryCodeTT string = "TT"

	// DomainDataClaimNoticeAddressCountryCodeTV captures enum value "TV"
	DomainDataClaimNoticeAddressCountryCodeTV string = "TV"

	// DomainDataClaimNoticeAddressCountryCodeTW captures enum value "TW"
	DomainDataClaimNoticeAddressCountryCodeTW string = "TW"

	// DomainDataClaimNoticeAddressCountryCodeTZ captures enum value "TZ"
	DomainDataClaimNoticeAddressCountryCodeTZ string = "TZ"

	// DomainDataClaimNoticeAddressCountryCodeUA captures enum value "UA"
	DomainDataClaimNoticeAddressCountryCodeUA string = "UA"

	// DomainDataClaimNoticeAddressCountryCodeUG captures enum value "UG"
	DomainDataClaimNoticeAddressCountryCodeUG string = "UG"

	// DomainDataClaimNoticeAddressCountryCodeUM captures enum value "UM"
	DomainDataClaimNoticeAddressCountryCodeUM string = "UM"

	// DomainDataClaimNoticeAddressCountryCodeUNKNOWN captures enum value "UNKNOWN"
	DomainDataClaimNoticeAddressCountryCodeUNKNOWN string = "UNKNOWN"

	// DomainDataClaimNoticeAddressCountryCodeUS captures enum value "US"
	DomainDataClaimNoticeAddressCountryCodeUS string = "US"

	// DomainDataClaimNoticeAddressCountryCodeUY captures enum value "UY"
	DomainDataClaimNoticeAddressCountryCodeUY string = "UY"

	// DomainDataClaimNoticeAddressCountryCodeUZ captures enum value "UZ"
	DomainDataClaimNoticeAddressCountryCodeUZ string = "UZ"

	// DomainDataClaimNoticeAddressCountryCodeVA captures enum value "VA"
	DomainDataClaimNoticeAddressCountryCodeVA string = "VA"

	// DomainDataClaimNoticeAddressCountryCodeVC captures enum value "VC"
	DomainDataClaimNoticeAddressCountryCodeVC string = "VC"

	// DomainDataClaimNoticeAddressCountryCodeVE captures enum value "VE"
	DomainDataClaimNoticeAddressCountryCodeVE string = "VE"

	// DomainDataClaimNoticeAddressCountryCodeVG captures enum value "VG"
	DomainDataClaimNoticeAddressCountryCodeVG string = "VG"

	// DomainDataClaimNoticeAddressCountryCodeVI captures enum value "VI"
	DomainDataClaimNoticeAddressCountryCodeVI string = "VI"

	// DomainDataClaimNoticeAddressCountryCodeVN captures enum value "VN"
	DomainDataClaimNoticeAddressCountryCodeVN string = "VN"

	// DomainDataClaimNoticeAddressCountryCodeVU captures enum value "VU"
	DomainDataClaimNoticeAddressCountryCodeVU string = "VU"

	// DomainDataClaimNoticeAddressCountryCodeWF captures enum value "WF"
	DomainDataClaimNoticeAddressCountryCodeWF string = "WF"

	// DomainDataClaimNoticeAddressCountryCodeWS captures enum value "WS"
	DomainDataClaimNoticeAddressCountryCodeWS string = "WS"

	// DomainDataClaimNoticeAddressCountryCodeXK captures enum value "XK"
	DomainDataClaimNoticeAddressCountryCodeXK string = "XK"

	// DomainDataClaimNoticeAddressCountryCodeYE captures enum value "YE"
	DomainDataClaimNoticeAddressCountryCodeYE string = "YE"

	// DomainDataClaimNoticeAddressCountryCodeYT captures enum value "YT"
	DomainDataClaimNoticeAddressCountryCodeYT string = "YT"

	// DomainDataClaimNoticeAddressCountryCodeZA captures enum value "ZA"
	DomainDataClaimNoticeAddressCountryCodeZA string = "ZA"

	// DomainDataClaimNoticeAddressCountryCodeZM captures enum value "ZM"
	DomainDataClaimNoticeAddressCountryCodeZM string = "ZM"

	// DomainDataClaimNoticeAddressCountryCodeZW captures enum value "ZW"
	DomainDataClaimNoticeAddressCountryCodeZW string = "ZW"
)

// prop value enum
func (m *DomainDataClaimNoticeAddress) validateCountryCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, domainDataClaimNoticeAddressTypeCountryCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DomainDataClaimNoticeAddress) validateCountryCode(formats strfmt.Registry) error {

	if swag.IsZero(m.CountryCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateCountryCodeEnum("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	return nil
}

func (m *DomainDataClaimNoticeAddress) validateStreets(formats strfmt.Registry) error {

	if err := validate.Required("streets", "body", m.Streets); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDataClaimNoticeAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDataClaimNoticeAddress) UnmarshalBinary(b []byte) error {
	var res DomainDataClaimNoticeAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
