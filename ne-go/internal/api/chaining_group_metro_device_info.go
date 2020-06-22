// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChainingGroupMetroDeviceInfo chaining group metro device info
//
// swagger:model ChainingGroupMetroDeviceInfo
type ChainingGroupMetroDeviceInfo struct {

	// An array of devices.
	Devices []*ChainingDeviceInfo `json:"devices"`

	// Metro code.
	// Required: true
	MetroCode *string `json:"metroCode"`

	// Metro Throughput.
	Throughput string `json:"throughput,omitempty"`

	// Throughput unit.
	ThroughputUnit string `json:"throughputUnit,omitempty"`
}

// Validate validates this chaining group metro device info
func (m *ChainingGroupMetroDeviceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetroCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChainingGroupMetroDeviceInfo) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChainingGroupMetroDeviceInfo) validateMetroCode(formats strfmt.Registry) error {

	if err := validate.Required("metroCode", "body", m.MetroCode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChainingGroupMetroDeviceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChainingGroupMetroDeviceInfo) UnmarshalBinary(b []byte) error {
	var res ChainingGroupMetroDeviceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
