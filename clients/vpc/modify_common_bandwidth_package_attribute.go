package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyCommonBandwidthPackageAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCommonBandwidthPackageAttributeRequest) Invoke(client goaliyun.Client) (*ModifyCommonBandwidthPackageAttributeResponse, error) {
	resp := &ModifyCommonBandwidthPackageAttributeResponse{}
	err := client.Request("vpc", "ModifyCommonBandwidthPackageAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCommonBandwidthPackageAttributeResponse struct {
	RequestId goaliyun.String
}
