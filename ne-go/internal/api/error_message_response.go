// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorMessageResponse error message response
//
// swagger:model ErrorMessageResponse
type ErrorMessageResponse struct {

	// error code
	ErrorCode string `json:"errorCode,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// more info
	MoreInfo string `json:"moreInfo,omitempty"`

	// property
	Property string `json:"property,omitempty"`
}

// Validate validates this error message response
func (m *ErrorMessageResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorMessageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorMessageResponse) UnmarshalBinary(b []byte) error {
	var res ErrorMessageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
