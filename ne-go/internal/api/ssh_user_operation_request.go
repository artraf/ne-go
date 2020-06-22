// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SSHUserOperationRequest Ssh user operation request
//
// swagger:model SshUserOperationRequest
type SSHUserOperationRequest struct {

	// SSH operation to be performed
	// Required: true
	// Enum: [CREATE DELETE REUSE]
	Action *string `json:"action"`

	// SSH Password
	SSHPassword string `json:"sshPassword,omitempty"`

	// Required for DELETE operation.
	SSHUserUUID string `json:"sshUserUuid,omitempty"`

	// SSH User name
	SSHUsername string `json:"sshUsername,omitempty"`
}

// Validate validates this Ssh user operation request
func (m *SSHUserOperationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sshUserOperationRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","DELETE","REUSE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sshUserOperationRequestTypeActionPropEnum = append(sshUserOperationRequestTypeActionPropEnum, v)
	}
}

const (

	// SSHUserOperationRequestActionCREATE captures enum value "CREATE"
	SSHUserOperationRequestActionCREATE string = "CREATE"

	// SSHUserOperationRequestActionDELETE captures enum value "DELETE"
	SSHUserOperationRequestActionDELETE string = "DELETE"

	// SSHUserOperationRequestActionREUSE captures enum value "REUSE"
	SSHUserOperationRequestActionREUSE string = "REUSE"
)

// prop value enum
func (m *SSHUserOperationRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sshUserOperationRequestTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SSHUserOperationRequest) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SSHUserOperationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHUserOperationRequest) UnmarshalBinary(b []byte) error {
	var res SSHUserOperationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
