package waf

import (
	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	Duration          int64  `position:"Query" name:"Duration"`
	PricingCycle      string `position:"Query" name:"PricingCycle"`
	PackageCode       string `position:"Query" name:"PackageCode"`
	ExtDomainPackage  int64  `position:"Query" name:"ExtDomainPackage"`
	ExtBandwidth      int64  `position:"Query" name:"ExtBandwidth"`
	IsAutoRenew       string `position:"Query" name:"IsAutoRenew"`
	AutoRenewDuration int64  `position:"Query" name:"AutoRenewDuration"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("waf-openapi", "CreateInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	OrderId    goaliyun.String
	InstanceId goaliyun.String
	RequestId  goaliyun.String
}
