package vpc

import (
	"github.com/morlay/goaliyun"
)

type UnassociateGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	InstanceType                 string `position:"Query" name:"InstanceType"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *UnassociateGlobalAccelerationInstanceRequest) Invoke(client goaliyun.Client) (*UnassociateGlobalAccelerationInstanceResponse, error) {
	resp := &UnassociateGlobalAccelerationInstanceResponse{}
	err := client.Request("vpc", "UnassociateGlobalAccelerationInstance", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UnassociateGlobalAccelerationInstanceResponse struct {
	RequestId goaliyun.String
}
