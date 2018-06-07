package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *DeleteGlobalAccelerationInstanceRequest) Invoke(client goaliyun.Client) (*DeleteGlobalAccelerationInstanceResponse, error) {
	resp := &DeleteGlobalAccelerationInstanceResponse{}
	err := client.Request("vpc", "DeleteGlobalAccelerationInstance", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteGlobalAccelerationInstanceResponse struct {
	RequestId goaliyun.String
}
