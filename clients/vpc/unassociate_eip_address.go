package vpc

import (
	"github.com/morlay/goaliyun"
)

type UnassociateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	Force                string `position:"Query" name:"Force"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *UnassociateEipAddressRequest) Invoke(client goaliyun.Client) (*UnassociateEipAddressResponse, error) {
	resp := &UnassociateEipAddressResponse{}
	err := client.Request("vpc", "UnassociateEipAddress", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnassociateEipAddressResponse struct {
	RequestId goaliyun.String
}
