package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyBandwidthPackageAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyBandwidthPackageAttributeRequest) Invoke(client goaliyun.Client) (*ModifyBandwidthPackageAttributeResponse, error) {
	resp := &ModifyBandwidthPackageAttributeResponse{}
	err := client.Request("vpc", "ModifyBandwidthPackageAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyBandwidthPackageAttributeResponse struct {
	RequestId goaliyun.String
}
