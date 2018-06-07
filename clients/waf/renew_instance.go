package waf

import (
	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	Duration     int64  `position:"Query" name:"Duration"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("waf-openapi", "RenewInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	OrderId   goaliyun.String
	RequestId goaliyun.String
}
