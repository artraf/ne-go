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

// PriceResponse price response
//
// swagger:model PriceResponse
type PriceResponse struct {

	// billing commencement date
	BillingCommencementDate string `json:"billingCommencementDate,omitempty"`

	// billing enabled
	BillingEnabled bool `json:"billingEnabled,omitempty"`

	// Monthly recurring charges
	Charges []*Charges `json:"charges"`

	// currency
	Currency string `json:"currency,omitempty"`
}

// Validate validates this price response
func (m *PriceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceResponse) validateCharges(formats strfmt.Registry) error {

	if swag.IsZero(m.Charges) { // not required
		return nil
	}

	for i := 0; i < len(m.Charges); i++ {
		if swag.IsZero(m.Charges[i]) { // not required
			continue
		}

		if m.Charges[i] != nil {
			if err := m.Charges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("charges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceResponse) UnmarshalBinary(b []byte) error {
	var res PriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
