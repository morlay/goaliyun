package yundun

import (
	"github.com/morlay/goaliyun"
)

type SetDdosAutoRequest struct {
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *SetDdosAutoRequest) Invoke(client goaliyun.Client) (*SetDdosAutoResponse, error) {
	resp := &SetDdosAutoResponse{}
	err := client.Request("yundun", "SetDdosAuto", "2015-04-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type SetDdosAutoResponse struct {
	RequestId goaliyun.String
}
