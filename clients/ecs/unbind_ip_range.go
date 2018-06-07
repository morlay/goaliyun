package ecs

import (
	"github.com/morlay/goaliyun"
)

type UnbindIpRangeRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UnbindIpRangeRequest) Invoke(client goaliyun.Client) (*UnbindIpRangeResponse, error) {
	resp := &UnbindIpRangeResponse{}
	err := client.Request("ecs", "UnbindIpRange", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnbindIpRangeResponse struct {
	RequestId goaliyun.String
}
