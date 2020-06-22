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

// VirtualDeviceRequest virtual device request
//
// swagger:model VirtualDeviceRequest
type VirtualDeviceRequest struct {

	// Account number. Either an account number or accountReferenceId is required.
	AccountNumber string `json:"accountNumber,omitempty"`

	// AccountReferenceId. This is a temporary ID that can be used to create a device when the account status is still pending, not active. Either an account number or accountReferenceId is required.
	AccountReferenceID string `json:"accountReferenceId,omitempty"`

	// IP addresses, no more than 50, in CIDR format
	ACL []string `json:"acl"`

	// Secondary additional bandwidth to be configured (in Mbps for HA). Default bandwidth provided is 15 Mbps.
	AdditionalBandwidth int32 `json:"additionalBandwidth,omitempty"`

	// Virtual device type (device type code)
	// Required: true
	DeviceTypeCode *string `json:"deviceTypeCode"`

	// Host name prefix for identification. Only a-z, A-Z, 0-9 and hyphen(-) are allowed. It should start with a letter and end with a letter or a digit. Also, it should be minimum 2 and maximum 10 characters long.
	// Required: true
	HostNamePrefix *string `json:"hostNamePrefix"`

	// For Juniper devices you need to provide a licenseFileId if you want to BYOL (Bring Your Own License). You get a licenseFileId when you upload a license file by calling license upload API (Upload a license file before creating a virtual device). For Cisco devices, you do not need to provide a licenseFileId at the time of device creation. Once the device is provisioned, you can get the deviceSerialNo by calling Get virtual device by UUID API. With the deviceSerialNo you can generate a license file on Cisco site. Afterward, you can upload the license file by calling license upload API (Upload a license file after creating a virtual device).
	LicenseFileID string `json:"licenseFileId,omitempty"`

	// license key
	LicenseKey string `json:"licenseKey,omitempty"`

	// License type. One of SUB (Subscription) or BYOL (Bring Your Own License)
	// Required: true
	LicenseMode *string `json:"licenseMode"`

	// license secret
	LicenseSecret string `json:"licenseSecret,omitempty"`

	// In case you want to BYOL (Bring Your Own License) for a Palo Alto device, you must provide a license token. This field must have 8 alphanumeric characters.
	LicenseToken string `json:"licenseToken,omitempty"`

	// Metro code
	// Required: true
	MetroCode *string `json:"metroCode"`

	// Email addresses for notification. We need a minimum of 1 and no more than 5 email addresses.
	// Required: true
	Notifications []string `json:"notifications"`

	// Software package code
	PackageCode string `json:"packageCode,omitempty"`

	// secondary
	Secondary *VirtualDevicHARequest `json:"secondary,omitempty"`

	// site Id
	SiteID string `json:"siteId,omitempty"`

	// Array of sshUsernames and passwords
	SSHUsers []*SSHUsers `json:"sshUsers"`

	// system Ip address
	SystemIPAddress string `json:"systemIpAddress,omitempty"`

	// Device throughput. This is required for Cisco and Juniper devices.
	Throughput int32 `json:"throughput,omitempty"`

	// Throughput unit. This is required for Cisco and Juniper devices.
	ThroughputUnit string `json:"throughputUnit,omitempty"`

	// Virtual device name for identification. This should be minimum 3 and maximum 50 characters long.
	// Required: true
	VirtualDeviceName *string `json:"virtualDeviceName"`
}

// Validate validates this virtual device request
func (m *VirtualDeviceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceTypeCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostNamePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetroCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualDeviceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualDeviceRequest) validateDeviceTypeCode(formats strfmt.Registry) error {

	if err := validate.Required("deviceTypeCode", "body", m.DeviceTypeCode); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceRequest) validateHostNamePrefix(formats strfmt.Registry) error {

	if err := validate.Required("hostNamePrefix", "body", m.HostNamePrefix); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceRequest) validateLicenseMode(formats strfmt.Registry) error {

	if err := validate.Required("licenseMode", "body", m.LicenseMode); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceRequest) validateMetroCode(formats strfmt.Registry) error {

	if err := validate.Required("metroCode", "body", m.MetroCode); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceRequest) validateNotifications(formats strfmt.Registry) error {

	if err := validate.Required("notifications", "body", m.Notifications); err != nil {
		return err
	}

	return nil
}

func (m *VirtualDeviceRequest) validateSecondary(formats strfmt.Registry) error {

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

func (m *VirtualDeviceRequest) validateSSHUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.SSHUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.SSHUsers); i++ {
		if swag.IsZero(m.SSHUsers[i]) { // not required
			continue
		}

		if m.SSHUsers[i] != nil {
			if err := m.SSHUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VirtualDeviceRequest) validateVirtualDeviceName(formats strfmt.Registry) error {

	if err := validate.Required("virtualDeviceName", "body", m.VirtualDeviceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualDeviceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualDeviceRequest) UnmarshalBinary(b []byte) error {
	var res VirtualDeviceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
