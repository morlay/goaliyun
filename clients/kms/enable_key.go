package kms

import (
	"github.com/morlay/goaliyun"
)

type EnableKeyRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *EnableKeyRequest) Invoke(client goaliyun.Client) (*EnableKeyResponse, error) {
	resp := &EnableKeyResponse{}
	err := client.Request("kms", "EnableKey", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type EnableKeyResponse struct {
	RequestId goaliyun.String
}
