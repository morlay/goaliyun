package vpc

import (
	"github.com/morlay/goaliyun"
)

type AssociateGlobalAccelerationInstanceRequest struct {
	ResourceOwnerId              int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount         string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string `position:"Query" name:"OwnerAccount"`
	BackendServerId              string `position:"Query" name:"BackendServerId"`
	OwnerId                      int64  `position:"Query" name:"OwnerId"`
	GlobalAccelerationInstanceId string `position:"Query" name:"GlobalAccelerationInstanceId"`
	BackendServerRegionId        string `position:"Query" name:"BackendServerRegionId"`
	BackendServerType            string `position:"Query" name:"BackendServerType"`
	RegionId                     string `position:"Query" name:"RegionId"`
}

func (req *AssociateGlobalAccelerationInstanceRequest) Invoke(client goaliyun.Client) (*AssociateGlobalAccelerationInstanceResponse, error) {
	resp := &AssociateGlobalAccelerationInstanceResponse{}
	err := client.Request("vpc", "AssociateGlobalAccelerationInstance", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AssociateGlobalAccelerationInstanceResponse struct {
	RequestId goaliyun.String
}
