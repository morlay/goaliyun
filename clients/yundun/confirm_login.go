package yundun

import (
	"github.com/morlay/goaliyun"
)

type ConfirmLoginRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Time       string `position:"Query" name:"Time"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ConfirmLoginRequest) Invoke(client goaliyun.Client) (*ConfirmLoginResponse, error) {
	resp := &ConfirmLoginResponse{}
	err := client.Request("yundun", "ConfirmLogin", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ConfirmLoginResponse struct {
	RequestId goaliyun.String
}
