// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChainingMetroThroughputPatchRequest chaining metro throughput patch request
//
// swagger:model ChainingMetroThroughputPatchRequest
type ChainingMetroThroughputPatchRequest struct {

	// An array of metroCodes and their throughput values.
	MetrosThroughput []*ChainingGroupMetroThroughputInfo `json:"metrosThroughput"`
}

// Validate validates this chaining metro throughput patch request
func (m *ChainingMetroThroughputPatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetrosThroughput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChainingMetroThroughputPatchRequest) validateMetrosThroughput(formats strfmt.Registry) error {

	if swag.IsZero(m.MetrosThroughput) { // not required
		return nil
	}

	for i := 0; i < len(m.MetrosThroughput); i++ {
		if swag.IsZero(m.MetrosThroughput[i]) { // not required
			continue
		}

		if m.MetrosThroughput[i] != nil {
			if err := m.MetrosThroughput[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrosThroughput" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChainingMetroThroughputPatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChainingMetroThroughputPatchRequest) UnmarshalBinary(b []byte) error {
	var res ChainingMetroThroughputPatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
