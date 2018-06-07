package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyInstanceSpecRequest struct {
	ResourceOwnerId                  int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount             string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                      string `position:"Query" name:"ClientToken"`
	AllowMigrateAcrossZone           string `position:"Query" name:"AllowMigrateAcrossZone"`
	OwnerAccount                     string `position:"Query" name:"OwnerAccount"`
	InternetMaxBandwidthOut          int64  `position:"Query" name:"InternetMaxBandwidthOut"`
	OwnerId                          int64  `position:"Query" name:"OwnerId"`
	TemporaryInternetMaxBandwidthOut int64  `position:"Query" name:"TemporaryInternetMaxBandwidthOut"`
	SystemDiskCategory               string `position:"Query" name:"SystemDiskCategory"`
	TemporaryStartTime               string `position:"Query" name:"TemporaryStartTime"`
	Async                            string `position:"Query" name:"Async"`
	InstanceId                       string `position:"Query" name:"InstanceId"`
	InstanceType                     string `position:"Query" name:"InstanceType"`
	TemporaryEndTime                 string `position:"Query" name:"TemporaryEndTime"`
	InternetMaxBandwidthIn           int64  `position:"Query" name:"InternetMaxBandwidthIn"`
	RegionId                         string `position:"Query" name:"RegionId"`
}

func (req *ModifyInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyInstanceSpecResponse, error) {
	resp := &ModifyInstanceSpecResponse{}
	err := client.Request("ecs", "ModifyInstanceSpec", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyInstanceSpecResponse struct {
	RequestId goaliyun.String
}
