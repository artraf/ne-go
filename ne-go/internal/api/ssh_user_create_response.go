// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSHUserCreateResponse Ssh user create response
//
// swagger:model SshUserCreateResponse
type SSHUserCreateResponse struct {

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this Ssh user create response
func (m *SSHUserCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSHUserCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHUserCreateResponse) UnmarshalBinary(b []byte) error {
	var res SSHUserCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
