package kms

import (
	"github.com/morlay/goaliyun"
)

type CancelKeyDeletionRequest struct {
	KeyId    string `position:"Query" name:"KeyId"`
	STSToken string `position:"Query" name:"STSToken"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *CancelKeyDeletionRequest) Invoke(client goaliyun.Client) (*CancelKeyDeletionResponse, error) {
	resp := &CancelKeyDeletionResponse{}
	err := client.Request("kms", "CancelKeyDeletion", "2016-01-20", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CancelKeyDeletionResponse struct {
	RequestId goaliyun.String
}
