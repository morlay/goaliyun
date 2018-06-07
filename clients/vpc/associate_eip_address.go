package vpc

import (
	"github.com/morlay/goaliyun"
)

type AssociateEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceRegionId     string `position:"Query" name:"InstanceRegionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AssociateEipAddressRequest) Invoke(client goaliyun.Client) (*AssociateEipAddressResponse, error) {
	resp := &AssociateEipAddressResponse{}
	err := client.Request("vpc", "AssociateEipAddress", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssociateEipAddressResponse struct {
	RequestId goaliyun.String
}
