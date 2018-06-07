package ecs

import (
	"github.com/morlay/goaliyun"
)

type AddIpRangeRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddIpRangeRequest) Invoke(client goaliyun.Client) (*AddIpRangeResponse, error) {
	resp := &AddIpRangeResponse{}
	err := client.Request("ecs", "AddIpRange", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddIpRangeResponse struct {
	RequestId goaliyun.String
}
