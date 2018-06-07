package hsm

import (
	"github.com/morlay/goaliyun"
)

type RenewInstanceRequest struct {
	Period          int64  `position:"Query" name:"Period"`
	PeriodUnit      string `position:"Query" name:"PeriodUnit"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	ClientToken     string `position:"Query" name:"ClientToken"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *RenewInstanceRequest) Invoke(client goaliyun.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("hsm", "RenewInstance", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	RequestId goaliyun.String
}
