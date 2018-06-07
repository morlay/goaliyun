package sts

import (
	"github.com/morlay/goaliyun"
)

type GenerateSessionAccessKeyRequest struct {
	DurationSeconds int64  `position:"Query" name:"DurationSeconds"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *GenerateSessionAccessKeyRequest) Invoke(client goaliyun.Client) (*GenerateSessionAccessKeyResponse, error) {
	resp := &GenerateSessionAccessKeyResponse{}
	err := client.Request("sts", "GenerateSessionAccessKey", "2015-04-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GenerateSessionAccessKeyResponse struct {
	RequestId        goaliyun.String
	SessionAccessKey GenerateSessionAccessKeySessionAccessKey
}

type GenerateSessionAccessKeySessionAccessKey struct {
	SessionAccessKeyId     goaliyun.String
	SessionAccessKeySecret goaliyun.String
	Expiration             goaliyun.String
}
