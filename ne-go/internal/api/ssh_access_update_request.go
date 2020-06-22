// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSHAccessUpdateRequest Ssh access update request
//
// swagger:model SshAccessUpdateRequest
type SSHAccessUpdateRequest struct {

	// ssh Acl
	SSHACL []interface{} `json:"sshAcl"`
}

// Validate validates this Ssh access update request
func (m *SSHAccessUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSHAccessUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHAccessUpdateRequest) UnmarshalBinary(b []byte) error {
	var res SSHAccessUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
