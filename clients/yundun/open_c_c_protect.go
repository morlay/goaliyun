package yundun

import (
	"github.com/morlay/goaliyun"
)

type OpenCCProtectRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *OpenCCProtectRequest) Invoke(client goaliyun.Client) (*OpenCCProtectResponse, error) {
	resp := &OpenCCProtectResponse{}
	err := client.Request("yundun", "OpenCCProtect", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type OpenCCProtectResponse struct {
	RequestId goaliyun.String
}
