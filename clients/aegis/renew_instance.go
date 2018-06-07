package aegis

import (
	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	Duration     int64  `position:"Query" name:"Duration"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	VmNumber     string `position:"Query" name:"VmNumber"`
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("aegis", "RenewInstance", "2016-11-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	OrderId   goaliyun.String
	RequestId goaliyun.String
}
