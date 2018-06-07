package vpc

import (
	"github.com/morlay/goaliyun"
)

type RemoveCommonBandwidthPackageIpRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpInstanceId         string `position:"Query" name:"IpInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *RemoveCommonBandwidthPackageIpRequest) Invoke(client goaliyun.Client) (*RemoveCommonBandwidthPackageIpResponse, error) {
	resp := &RemoveCommonBandwidthPackageIpResponse{}
	err := client.Request("vpc", "RemoveCommonBandwidthPackageIp", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveCommonBandwidthPackageIpResponse struct {
	RequestId goaliyun.String
}
