package api

//SSHPublicKey describes SSH public key
type SSHPublicKey struct {
	UUID      *string `json:"uuid,omitempty"`
	KeyName   *string `json:"keyName,omitempty"`
	KeyValue  *string `json:"keyValue,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	KeyType  *string `json:"keyType,omitempty"`
}
