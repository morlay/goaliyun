package yundun

import (
	"github.com/morlay/goaliyun"
)

type CloseCCProtectRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CloseCCProtectRequest) Invoke(client goaliyun.Client) (*CloseCCProtectResponse, error) {
	resp := &CloseCCProtectResponse{}
	err := client.Request("yundun", "CloseCCProtect", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CloseCCProtectResponse struct {
	RequestId goaliyun.String
}
