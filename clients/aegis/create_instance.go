package aegis

import (
	"github.com/morlay/goaliyun"
)

type CreateInstanceRequest struct {
	Duration          int64  `position:"Query" name:"Duration"`
	IsAutoRenew       string `position:"Query" name:"IsAutoRenew"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	VmNumber          int64  `position:"Query" name:"VmNumber"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	VersionCode       int64  `position:"Query" name:"VersionCode"`
	PricingCycle      string `position:"Query" name:"PricingCycle"`
	AutoRenewDuration int64  `position:"Query" name:"AutoRenewDuration"`
	RegionId          string `position:"Query" name:"RegionId"`
}

func (req *CreateInstanceRequest) Invoke(client goaliyun.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("aegis", "CreateInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	OrderId    goaliyun.String
	InstanceId goaliyun.String
	RequestId  goaliyun.String
}
