// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OAuthErrorResponse o auth error response
//
// swagger:model OAuthErrorResponse
type OAuthErrorResponse struct {

	// developer message
	DeveloperMessage string `json:"developerMessage,omitempty"`

	// error code
	ErrorCode string `json:"errorCode,omitempty"`

	// error domain
	ErrorDomain string `json:"errorDomain,omitempty"`

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// error title
	ErrorTitle string `json:"errorTitle,omitempty"`
}

// Validate validates this o auth error response
func (m *OAuthErrorResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuthErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthErrorResponse) UnmarshalBinary(b []byte) error {
	var res OAuthErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
