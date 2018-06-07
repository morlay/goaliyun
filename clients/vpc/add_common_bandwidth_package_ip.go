package vpc

import (
	"github.com/morlay/goaliyun"
)

type AddCommonBandwidthPackageIpRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpInstanceId         string `position:"Query" name:"IpInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *AddCommonBandwidthPackageIpRequest) Invoke(client goaliyun.Client) (*AddCommonBandwidthPackageIpResponse, error) {
	resp := &AddCommonBandwidthPackageIpResponse{}
	err := client.Request("vpc", "AddCommonBandwidthPackageIp", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddCommonBandwidthPackageIpResponse struct {
	RequestId goaliyun.String
}
