package waf

import (
	"github.com/morlay/goaliyun"
)

type UpgradeInstanceRequest struct {
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	ClientToken      string `position:"Query" name:"ClientToken"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	PackageCode      string `position:"Query" name:"PackageCode"`
	ExtDomainPackage int64  `position:"Query" name:"ExtDomainPackage"`
	ExtBandwidth     int64  `position:"Query" name:"ExtBandwidth"`
	RegionId         string `position:"Query" name:"RegionId"`
}

func (req *UpgradeInstanceRequest) Invoke(client goaliyun.Client) (*UpgradeInstanceResponse, error) {
	resp := &UpgradeInstanceResponse{}
	err := client.Request("waf-openapi", "UpgradeInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpgradeInstanceResponse struct {
	OrderId   goaliyun.String
	RequestId goaliyun.String
}
