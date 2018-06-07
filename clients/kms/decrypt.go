package kms

import (
	"github.com/morlay/goaliyun"
)

type DecryptRequest struct {
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	STSToken          string `position:"Query" name:"STSToken"`
	CiphertextBlob    string `position:"Query" name:"CiphertextBlob"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *DecryptRequest) Invoke(client goaliyun.Client) (*DecryptResponse, error) {
	resp := &DecryptResponse{}
	err := client.Request("kms", "Decrypt", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DecryptResponse struct {
	Plaintext goaliyun.String
	KeyId     goaliyun.String
	RequestId goaliyun.String
}
