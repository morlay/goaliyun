package kms

import (
	"github.com/morlay/goaliyun"
)

type DisableKeyRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DisableKeyRequest) Invoke(client goaliyun.Client) (*DisableKeyResponse, error) {
	resp := &DisableKeyResponse{}
	err := client.Request("kms", "DisableKey", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DisableKeyResponse struct {
	RequestId goaliyun.String
}
