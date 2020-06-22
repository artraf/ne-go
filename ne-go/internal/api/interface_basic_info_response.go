// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InterfaceBasicInfoResponse interface basic info response
//
// swagger:model InterfaceBasicInfoResponse
type InterfaceBasicInfoResponse struct {

	// asn
	Asn int64 `json:"asn,omitempty"`

	// assigned type
	AssignedType string `json:"assignedType,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// ipv4 mask
	IPV4Mask string `json:"ipv4Mask,omitempty"`

	// ipv4 subnet
	IPV4Subnet string `json:"ipv4Subnet,omitempty"`

	// mac address
	MacAddress string `json:"macAddress,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// operational status
	OperationalStatus string `json:"operationalStatus,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this interface basic info response
func (m *InterfaceBasicInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InterfaceBasicInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InterfaceBasicInfoResponse) UnmarshalBinary(b []byte) error {
	var res InterfaceBasicInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
