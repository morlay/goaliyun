package vpc

import (
	"github.com/morlay/goaliyun"
)

type ActivateRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ActivateRouterInterfaceRequest) Invoke(client goaliyun.Client) (*ActivateRouterInterfaceResponse, error) {
	resp := &ActivateRouterInterfaceResponse{}
	err := client.Request("vpc", "ActivateRouterInterface", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ActivateRouterInterfaceResponse struct {
	RequestId goaliyun.String
}
