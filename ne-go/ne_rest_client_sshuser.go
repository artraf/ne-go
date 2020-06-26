package ne

import (
	"fmt"
	"ne-go/v1/internal/api"
	"net/url"

	"github.com/go-resty/resty/v2"
)

const (
	associateDevice   = "ADD"
	unassociateDevice = "DELETE"
)

type restSSHUserUpdateRequest struct {
	uuid           string
	newPassword    string
	newDevices     []string
	removedDevices []string
	c              RestClient
}

//CreateSSHUser creates new Network Edge SSH user with a given parameters and returns its UUID upon successful creation
func (c RestClient) CreateSSHUser(username string, password string, device string) (string, error) {
	u := c.baseURL + "/ne/v1/services/ssh-user"
	reqBody := api.SSHUserCreateRequest{
		Username:   &username,
		Password:   &password,
		DeviceUUID: &device,
	}
	respBody := api.SSHUserCreateResponse{}
	req := c.R().SetBody(&reqBody).SetResult(&respBody)
	if err := c.execute(req, resty.MethodPost, u); err != nil {
		return "", err
	}
	return respBody.UUID, nil
}

//GetSSHUser fetches details of a SSH user with a given UUID
func (c RestClient) GetSSHUser(uuid string) (*SSHUser, error) {
	url := c.baseURL + "/ne/v1/services/ssh-user/" + url.PathEscape(uuid)
	respBody := api.SSHUserInfoVerbose{}
	req := c.R().SetResult(&respBody)
	if err := c.execute(req, resty.MethodGet, url); err != nil {
		return nil, err
	}
	return mapSSHUserAPIToDomain(respBody), nil
}

//NewSSHUserUpdateRequest creates new composite update request for a user with a given UUID
func (c RestClient) NewSSHUserUpdateRequest(uuid string) SSHUserUpdateRequest {
	return &restSSHUserUpdateRequest{
		uuid: uuid,
		c:    c}
}

//DeleteSSHUser deletes ssh user with a given UUID
func (c RestClient) DeleteSSHUser(uuid string) error {
	user, err := c.GetSSHUser(uuid)
	if err != nil {
		return err
	}
	updateErr := UpdateError{}
	for _, dev := range user.DeviceUUIDs {
		if err := c.changeDeviceAssociation(unassociateDevice, uuid, dev); err != nil {
			updateErr.AddChangeError(changeTypeDelete, "devices", dev, err)
		}
	}
	if updateErr.ChangeErrorsCount() > 0 {
		return updateErr
	}
	return nil
}

func (req *restSSHUserUpdateRequest) WithNewPassword(password string) SSHUserUpdateRequest {
	req.newPassword = password
	return req
}

func (req *restSSHUserUpdateRequest) WithNewDevices(uuids []string) SSHUserUpdateRequest {
	req.newDevices = uuids
	return req
}

func (req *restSSHUserUpdateRequest) WithRemovedDevices(uuids []string) SSHUserUpdateRequest {
	req.removedDevices = uuids
	return req
}

func (req *restSSHUserUpdateRequest) Execute() error {
	updateErr := UpdateError{}
	if req.newPassword != "" {
		if err := req.c.changeUserPassword(req.uuid, req.newPassword); err != nil {
			updateErr.AddChangeError(changeTypeUpdate, "password", req.newPassword, err)
		}
	}
	for _, dev := range req.newDevices {
		if err := req.c.changeDeviceAssociation(associateDevice, req.uuid, dev); err != nil {
			updateErr.AddChangeError(changeTypeCreate, "devices", dev, err)
		}
	}
	for _, dev := range req.removedDevices {
		if err := req.c.changeDeviceAssociation(unassociateDevice, req.uuid, dev); err != nil {
			updateErr.AddChangeError(changeTypeDelete, "devices", dev, err)
		}
	}
	if updateErr.ChangeErrorsCount() > 0 {
		return updateErr
	}
	return nil
}

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// Unexported package methods
//_______________________________________________________________________

func (c RestClient) changeUserPassword(userID string, newPassword string) error {
	url := fmt.Sprintf("%s/ne/v1/services/ssh-user/%s",
		c.baseURL, url.PathEscape(userID))
	reqBody := api.SSHUserUpdateRequest{Password: &newPassword}
	req := c.R().SetBody(&reqBody)
	if err := c.execute(req, resty.MethodPut, url); err != nil {
		return err
	}
	return nil
}

func (c RestClient) changeDeviceAssociation(changeType string, userID string, deviceID string) error {
	url := fmt.Sprintf("%s/ne/v1/services/ssh-user/%s/association?deviceUuid=%s",
		c.baseURL, url.PathEscape(userID), url.PathEscape(deviceID))
	var method string
	switch changeType {
	case associateDevice:
		method = resty.MethodPatch
	case unassociateDevice:
		method = resty.MethodDelete
	default:
		return fmt.Errorf("unsupported association change type")
	}
	req := c.R().
		//due to bug in NE API that requires content type and content len = 0 altough there is no content needed in any case
		SetHeader("Content-Type", "application/json").
		SetBody("{}")
	if err := c.execute(req, method, url); err != nil {
		return err
	}
	return nil
}

func mapSSHUserAPIToDomain(apiUser api.SSHUserInfoVerbose) *SSHUser {
	return &SSHUser{
		UUID:        apiUser.UUID,
		Username:    apiUser.Username,
		DeviceUUIDs: apiUser.DeviceUuids,
		MetroCodes:  apiUser.Metros}
}
