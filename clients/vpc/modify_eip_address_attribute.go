package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyEipAddressAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyEipAddressAttributeRequest) Invoke(client goaliyun.Client) (*ModifyEipAddressAttributeResponse, error) {
	resp := &ModifyEipAddressAttributeResponse{}
	err := client.Request("vpc", "ModifyEipAddressAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyEipAddressAttributeResponse struct {
	RequestId goaliyun.String
}
