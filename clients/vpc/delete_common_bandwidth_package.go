package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteCommonBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCommonBandwidthPackageRequest) Invoke(client goaliyun.Client) (*DeleteCommonBandwidthPackageResponse, error) {
	resp := &DeleteCommonBandwidthPackageResponse{}
	err := client.Request("vpc", "DeleteCommonBandwidthPackage", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCommonBandwidthPackageResponse struct {
	RequestId goaliyun.String
}
