package kms

import (
	"github.com/morlay/goaliyun"
)

type EncryptRequest struct {
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	KeyId             string `position:"Query" name:"KeyId"`
	STSToken          string `position:"Query" name:"STSToken"`
	Plaintext         string `position:"Query" name:"Plaintext"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *EncryptRequest) Invoke(client goaliyun.Client) (*EncryptResponse, error) {
	resp := &EncryptResponse{}
	err := client.Request("kms", "Encrypt", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EncryptResponse struct {
	CiphertextBlob goaliyun.String
	KeyId          goaliyun.String
	RequestId      goaliyun.String
}
