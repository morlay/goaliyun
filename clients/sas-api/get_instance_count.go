package sas_api

import (
	"github.com/morlay/goaliyun"
)

type GetInstanceCountRequest struct {
	OwnerId  int64  `position:"Query" name:"OwnerId"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *GetInstanceCountRequest) Invoke(client goaliyun.Client) (*GetInstanceCountResponse, error) {
	resp := &GetInstanceCountResponse{}
	err := client.Request("sas-api", "GetInstanceCount", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetInstanceCountResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	Success   bool
	RequestId goaliyun.String
	Data      goaliyun.Integer
}
