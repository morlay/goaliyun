package vpc

import (
	"github.com/morlay/goaliyun"
)

type RemoveGlobalAccelerationInstanceIpRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	IpInstanceId                 string `position:"Query" name:"IpInstanceId"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *RemoveGlobalAccelerationInstanceIpRequest) Invoke(client goaliyun.Client) (*RemoveGlobalAccelerationInstanceIpResponse, error) {
	resp := &RemoveGlobalAccelerationInstanceIpResponse{}
	err := client.Request("vpc", "RemoveGlobalAccelerationInstanceIp", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveGlobalAccelerationInstanceIpResponse struct {
	RequestId goaliyun.String
}
