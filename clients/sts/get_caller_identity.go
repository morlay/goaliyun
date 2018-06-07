package sts

import (
	"github.com/morlay/goaliyun"
)

type GetCallerIdentityRequest struct {
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetCallerIdentityRequest) Invoke(client goaliyun.Client) (*GetCallerIdentityResponse, error) {
	resp := &GetCallerIdentityResponse{}
	err := client.Request("sts", "GetCallerIdentity", "2015-04-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetCallerIdentityResponse struct {
	AccountId goaliyun.String
	UserId    goaliyun.String
	Arn       goaliyun.String
	RequestId goaliyun.String
}
