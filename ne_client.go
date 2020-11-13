//Package ne implements Network Edge client
package ne

import (
	"fmt"
)

const (
	//DeviceStateInitializing Equinix is allocating resources and creating new device
	DeviceStateInitializing = "INITIALIZING"
	//DeviceStateProvisioning Network Edge device is booting
	DeviceStateProvisioning = "PROVISIONING"
	//DeviceStateWaitingPrimary secondary Network Edge device is waiting for provisioning of its redundant (primary) device
	DeviceStateWaitingPrimary = "WAITING_FOR_PRIMARY"
	//DeviceStateWaitingSecondary primary Network Edge device is waiting for provisioning of its redundant (secondary) device
	DeviceStateWaitingSecondary = "WAITING_FOR_SECONDARY"
	//DeviceStateFailed Network Edge device creation and provisioning have failed
	DeviceStateFailed = "FAILED"
	//DeviceStateProvisioned Network Edge device was successfully provisioned and is fully operational
	DeviceStateProvisioned = "PROVISIONED"
	//DeviceStateDeprovisioning Network Edge device is in process of deprovisioning
	DeviceStateDeprovisioning = "DEPROVISIONING"
	//DeviceStateDeprovisioned Network Edge device was successfully deprovisioned
	DeviceStateDeprovisioned = "DEPROVISIONED"

	//DeviceLicenseStateApplying license is in registration process
	DeviceLicenseStateApplying = "APPLYING_LICENSE"
	//DeviceLicenseStateRegistered license was successfully registered
	DeviceLicenseStateRegistered = "REGISTERED"
	//DeviceLicenseStateFailed license registration has failed
	DeviceLicenseStateFailed = "REGISTRATION_FAILED"

	//BGPStateIdle BGP peer state is idle
	BGPStateIdle = "Idle"
	//BGPStateConnect BGP peer state is connect
	BGPStateConnect = "Connect"
	//BGPStateActive BGP peer state is active
	BGPStateActive = "Active"
	//BGPStateOpenSent BGP peer state is OpenSent
	BGPStateOpenSent = "OpenSent"
	//BGPStateOpenConnect BGP peer state is OpenConfirm
	BGPStateOpenConnect = "OpenConfirm"
	//BGPStateEstablished BGP peer state is Established
	BGPStateEstablished = "Established"
	//BGPProvisioningStatusProvisioning BGP peering is provisioning
	BGPProvisioningStatusProvisioning = "PROVISIONING"
	//BGPProvisioningStatusPendingUpdate BGP peering is being updated
	BGPProvisioningStatusPendingUpdate = "PENDING_UPDATE"
	//BGPProvisioningStatusProvisioned BGP peering is provisioned
	BGPProvisioningStatusProvisioned = "PROVISIONED"
	//BGPProvisioningStatusFailed BGP peering failed
	BGPProvisioningStatusFailed = "FAILED"

	//ErrorCodeDeviceRemoved is used on attempt to remove device that is deprovisioning or already deprovisioned
	ErrorCodeDeviceRemoved = "IC-NE-VD-030"
)

//Client interface describes operations provided by Network Edge client library
type Client interface {
	GetAccounts(metroCode string) ([]Account, error)
	GetDeviceTypes() ([]DeviceType, error)
	GetDevicePlatforms(deviceTypeCode string) ([]DevicePlatform, error)
	GetDeviceSoftwareVersions(deviceTypeCode string) ([]DeviceSoftwareVersion, error)

	CreateDevice(device Device) (string, error)
	CreateRedundantDevice(primary Device, secondary Device) (string, string, error)
	GetDevice(uuid string) (*Device, error)
	GetDevices(statuses []string) ([]Device, error)
	NewDeviceUpdateRequest(uuid string) DeviceUpdateRequest
	DeleteDevice(uuid string) error

	CreateSSHUser(username string, password string, device string) (string, error)
	GetSSHUsers() ([]SSHUser, error)
	GetSSHUser(uuid string) (*SSHUser, error)
	NewSSHUserUpdateRequest(uuid string) SSHUserUpdateRequest
	DeleteSSHUser(uuid string) error

	CreateBGPConfiguration(config BGPConfiguration) (string, error)
	GetBGPConfiguration(uuid string) (*BGPConfiguration, error)
	NewBGPConfigurationUpdateRequest(uuid string) BGPUpdateRequest
	GetBGPConfigurationForConnection(uuid string) (*BGPConfiguration, error)

	CreateACLTemplate(template ACLTemplate) (string, error)
	GetACLTemplates() ([]ACLTemplate, error)
	GetACLTemplate(uuid string) (*ACLTemplate, error)
	ReplaceACLTemplate(uuid string, template ACLTemplate) error
	DeleteACLTemplate(uuid string) error
}

//DeviceUpdateRequest describes composite request to update given Network Edge device
type DeviceUpdateRequest interface {
	WithDeviceName(deviceName string) DeviceUpdateRequest
	WithTermLength(termLength int) DeviceUpdateRequest
	WithNotifications(notifications []string) DeviceUpdateRequest
	WithAdditionalBandwidth(additionalBandwidth int) DeviceUpdateRequest
	WithACLTemplate(templateID string) DeviceUpdateRequest
	Execute() error
}

//SSHUserUpdateRequest describes composite request to update given Network Edge SSH user
type SSHUserUpdateRequest interface {
	WithNewPassword(password string) SSHUserUpdateRequest
	WithDeviceChange(old []string, new []string) SSHUserUpdateRequest
	Execute() error
}

//BGPUpdateRequest describes request to update given BGP configuration
type BGPUpdateRequest interface {
	WithLocalIPAddress(localIPAddress string) BGPUpdateRequest
	WithLocalASN(localASN int) BGPUpdateRequest
	WithRemoteASN(remoteASN int) BGPUpdateRequest
	WithRemoteIPAddress(remoteIPAddress string) BGPUpdateRequest
	WithAuthenticationKey(authenticationKey string) BGPUpdateRequest
	Execute() error
}

//Error describes Network Edge error that occurs during API call processing
type Error struct {
	//ErrorCode is short error identifier
	ErrorCode string
	//ErrorMessage is textual description of an error
	ErrorMessage string
}

//ChangeError describes single error that occurred during update of selected target property
type ChangeError struct {
	Type   string
	Target string
	Value  interface{}
	Cause  error
}

func (e ChangeError) Error() string {
	return fmt.Sprintf("change type '%s', target '%s', value '%s', cause: '%s'", e.Type, e.Target, e.Value, e.Cause)
}

//UpdateError describes error that occurred during composite update request and consists of multiple atomic change errors
type UpdateError struct {
	Failed []ChangeError
}

//AddChangeError functions add new atomic change error to update error structure
func (e *UpdateError) AddChangeError(changeType string, target string, value interface{}, cause error) {
	e.Failed = append(e.Failed, ChangeError{
		Type:   changeType,
		Target: target,
		Value:  value,
		Cause:  cause})
}

//ChangeErrorsCount returns number of atomic change errors in a given composite update error
func (e UpdateError) ChangeErrorsCount() int {
	return len(e.Failed)
}

func (e UpdateError) Error() string {
	str := fmt.Sprintf("update error: %d changes failed.", len(e.Failed))
	for _, err := range e.Failed {
		str = fmt.Sprintf("%s [%s]", str, err.Error())
	}
	return str
}

//Account describes Network Edge customer account details
type Account struct {
	Name   string
	Number string
	Status string
	UCMID  string
}

//Device describes Network Edge device
type Device struct {
	UUID                string
	Name                string
	TypeCode            string
	Status              string
	LicenseStatus       string
	MetroCode           string
	IBX                 string
	Region              string
	Throughput          int
	ThroughputUnit      string
	HostName            string
	PackageCode         string
	Version             string
	IsBYOL              bool
	LicenseToken        string
	ACLTemplateUUID     string
	SSHIPAddress        string
	SSHIPFqdn           string
	AccountNumber       string
	Notifications       []string
	PurchaseOrderNumber string
	RedundancyType      string
	RedundantUUID       string
	TermLength          int
	AdditionalBandwidth int
	OrderReference      string
	InterfaceCount      int
	CoreCount           int
	IsSelfManaged       bool
	Interfaces          []DeviceInterface
	VendorConfiguration map[string]string
}

//DeviceInterface describes Network Edge device interface
type DeviceInterface struct {
	ID                int
	Name              string
	Status            string
	OperationalStatus string
	MACAddress        string
	IPAddress         string
	AssignedType      string
	Type              string
}

//DeviceType describes Network Edge device type
type DeviceType struct {
	Name        string
	Code        string
	Description string
	Vendor      string
	Category    string
	MetroCodes  []string
}

//DevicePlatform describes Network Edge platform configurations
//available for a given device type
type DevicePlatform struct {
	Flavor          string
	CoreCount       int
	Memory          int
	MemoryUnit      string
	PackageCodes    []string
	ManagementTypes []string
	LicenseOptions  []string
}

//DeviceSoftwareVersion describes available software packages and versions for a Network Edge device
type DeviceSoftwareVersion struct {
	Version          string
	ImageName        string
	Date             string
	Status           string
	IsStable         bool
	ReleaseNotesLink string
	PackageCodes     []string
}

//SSHUser describes Network Edge SSH user
type SSHUser struct {
	UUID        string
	Username    string
	Password    string
	DeviceUUIDs []string
}

//BGPConfiguration describes Network Edge BGP configuration
type BGPConfiguration struct {
	UUID               string
	ConnectionUUID     string
	DeviceUUID         string
	LocalIPAddress     string
	LocalASN           int
	RemoteIPAddress    string
	RemoteASN          int
	AuthenticationKey  string
	State              string
	ProvisioningStatus string
}

//ACLTemplate describes Network Edge device ACL template
type ACLTemplate struct {
	UUID            string
	Name            string
	Description     string
	MetroCode       string
	DeviceUUID      string
	DeviceACLStatus string
	InboundRules    []ACLTemplateInboundRule
}

//ACLTemplateInboundRule describes inbound ACL rule that is part of
//Network Edge device ACL template
type ACLTemplateInboundRule struct {
	SeqNo    int
	SrcType  string
	FQDN     string
	Subnets  []string
	Protocol string
	SrcPort  string
	DstPort  string
}
