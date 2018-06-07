package kms

import (
	"github.com/morlay/goaliyun"
)

type GenerateDataKeyRequest struct {
	EncryptionContext string `position:"Query" name:"EncryptionContext"`
	KeyId             string `position:"Query" name:"KeyId"`
	KeySpec           string `position:"Query" name:"KeySpec"`
	STSToken          string `position:"Query" name:"STSToken"`
	NumberOfBytes     int64  `position:"Query" name:"NumberOfBytes"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *GenerateDataKeyRequest) Invoke(client goaliyun.Client) (*GenerateDataKeyResponse, error) {
	resp := &GenerateDataKeyResponse{}
	err := client.Request("kms", "GenerateDataKey", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GenerateDataKeyResponse struct {
	CiphertextBlob goaliyun.String
	KeyId          goaliyun.String
	Plaintext      goaliyun.String
	RequestId      goaliyun.String
}
