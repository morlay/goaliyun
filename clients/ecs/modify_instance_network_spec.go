package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceNetworkSpecRequest struct {
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	AutoPay                 string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken             string `position:"Query" name:"ClientToken"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	InternetMaxBandwidthOut int64  `position:"Query" name:"InternetMaxBandwidthOut"`
	EndTime                 string `position:"Query" name:"EndTime"`
	StartTime               string `position:"Query" name:"StartTime"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	InstanceId              string `position:"Query" name:"InstanceId"`
	NetworkChargeType       string `position:"Query" name:"NetworkChargeType"`
	InternetMaxBandwidthIn  int64  `position:"Query" name:"InternetMaxBandwidthIn"`
	AllocatePublicIp        string `position:"Query" name:"AllocatePublicIp"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceNetworkSpecRequest) Invoke(client goaliyun.Client) (*ModifyInstanceNetworkSpecResponse, error) {
	resp := &ModifyInstanceNetworkSpecResponse{}
	err := client.Request("ecs", "ModifyInstanceNetworkSpec", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceNetworkSpecResponse struct {
	RequestId goaliyun.String
	OrderId   goaliyun.String
}
