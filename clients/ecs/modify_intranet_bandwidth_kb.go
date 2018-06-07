package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyIntranetBandwidthKbRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	IntranetMaxBandwidthOut int64  `position:"Query" name:"IntranetMaxBandwidthOut"`
	InstanceId              string `position:"Query" name:"InstanceId"`
	IntranetMaxBandwidthIn  int64  `position:"Query" name:"IntranetMaxBandwidthIn"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ModifyIntranetBandwidthKbRequest) Invoke(client goaliyun.Client) (*ModifyIntranetBandwidthKbResponse, error) {
	resp := &ModifyIntranetBandwidthKbResponse{}
	err := client.Request("ecs", "ModifyIntranetBandwidthKb", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyIntranetBandwidthKbResponse struct {
	RequestId goaliyun.String
}
