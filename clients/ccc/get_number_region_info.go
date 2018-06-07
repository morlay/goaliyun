package ccc

import (
	"github.com/morlay/goaliyun"
)

type GetNumberRegionInfoRequest struct {
	Number     string `position:"Query" name:"Number"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *GetNumberRegionInfoRequest) Invoke(client goaliyun.Client) (*GetNumberRegionInfoResponse, error) {
	resp := &GetNumberRegionInfoResponse{}
	err := client.Request("ccc", "GetNumberRegionInfo", "2017-07-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetNumberRegionInfoResponse struct {
	RequestId   goaliyun.String
	Success     bool
	Code        goaliyun.String
	Message     goaliyun.String
	PhoneNumber GetNumberRegionInfoPhoneNumber
}

type GetNumberRegionInfoPhoneNumber struct {
	Number   goaliyun.String
	Province goaliyun.String
	City     goaliyun.String
}
