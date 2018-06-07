package ots

import (
	"github.com/morlay/goaliyun"
)

type GetInstanceRequest struct {
	Accept      string `position:"Header" name:"Accept"`
	VersionName string `position:"Header" name:"VersionName"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *GetInstanceRequest) Invoke(client goaliyun.Client) (*GetInstanceResponse, error) {
	resp := &GetInstanceResponse{}
	err := client.Request("ots", "GetInstance", "2013-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetInstanceResponse struct {
	RequestId goaliyun.String
}
