package waf

import (
	"github.com/morlay/goaliyun"
)

type ReleaseInstanceRequest struct {
	OwnerId    int64  `position:"Query" name:"OwnerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	RegionId   string `position:"Query" name:"RegionId"`
}

func (req *ReleaseInstanceRequest) Invoke(client goaliyun.Client) (*ReleaseInstanceResponse, error) {
	resp := &ReleaseInstanceResponse{}
	err := client.Request("waf-openapi", "ReleaseInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseInstanceResponse struct {
	RequestId goaliyun.String
}
