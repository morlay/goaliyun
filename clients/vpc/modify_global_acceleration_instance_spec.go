package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyGlobalAccelerationInstanceSpecRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                    string `position:"Query" name:"Bandwidth"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *ModifyGlobalAccelerationInstanceSpecRequest) Invoke(client goaliyun.Client) (*ModifyGlobalAccelerationInstanceSpecResponse, error) {
	resp := &ModifyGlobalAccelerationInstanceSpecResponse{}
	err := client.Request("vpc", "ModifyGlobalAccelerationInstanceSpec", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyGlobalAccelerationInstanceSpecResponse struct {
	RequestId goaliyun.String
}
