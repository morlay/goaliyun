package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyGlobalAccelerationInstanceAttributesRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	Name                         string `position:"Query" name:"Name"`
	Description                  string `position:"Query" name:"Description"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *ModifyGlobalAccelerationInstanceAttributesRequest) Invoke(client goaliyun.Client) (*ModifyGlobalAccelerationInstanceAttributesResponse, error) {
	resp := &ModifyGlobalAccelerationInstanceAttributesResponse{}
	err := client.Request("vpc", "ModifyGlobalAccelerationInstanceAttributes", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyGlobalAccelerationInstanceAttributesResponse struct {
	RequestId goaliyun.String
}
