package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyCommonBandwidthPackageSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCommonBandwidthPackageSpecRequest) Invoke(client goaliyun.Client) (*ModifyCommonBandwidthPackageSpecResponse, error) {
	resp := &ModifyCommonBandwidthPackageSpecResponse{}
	err := client.Request("vpc", "ModifyCommonBandwidthPackageSpec", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCommonBandwidthPackageSpecResponse struct {
	RequestId goaliyun.String
}
