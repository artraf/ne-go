// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VirtualDeviceDeleteRequest virtual device delete request
//
// swagger:model VirtualDeviceDeleteRequest
type VirtualDeviceDeleteRequest struct {

	// Some devices, e.g. Palo Alto devices, require a mandatory deactivation key for Equinix to successfully process the request.
	DeactivationKey string `json:"deactivationKey,omitempty"`

	// secondary
	Secondary *SecondaryDeviceDeleteRequest `json:"secondary,omitempty"`
}

// Validate validates this virtual device delete request
func (m *VirtualDeviceDeleteRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecondary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualDeviceDeleteRequest) validateSecondary(formats strfmt.Registry) error {

	if swag.IsZero(m.Secondary) { // not required
		return nil
	}

	if m.Secondary != nil {
		if err := m.Secondary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualDeviceDeleteRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualDeviceDeleteRequest) UnmarshalBinary(b []byte) error {
	var res VirtualDeviceDeleteRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
