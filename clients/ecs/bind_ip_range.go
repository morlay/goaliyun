package ecs

import (
	"github.com/morlay/goaliyun"
)

type BindIpRangeRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *BindIpRangeRequest) Invoke(client goaliyun.Client) (*BindIpRangeResponse, error) {
	resp := &BindIpRangeResponse{}
	err := client.Request("ecs", "BindIpRange", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type BindIpRangeResponse struct {
	RequestId goaliyun.String
}
