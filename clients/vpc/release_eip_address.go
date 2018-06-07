package vpc

import (
	"github.com/morlay/goaliyun"
)

type ReleaseEipAddressRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ReleaseEipAddressRequest) Invoke(client goaliyun.Client) (*ReleaseEipAddressResponse, error) {
	resp := &ReleaseEipAddressResponse{}
	err := client.Request("vpc", "ReleaseEipAddress", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ReleaseEipAddressResponse struct {
	RequestId goaliyun.String
}
