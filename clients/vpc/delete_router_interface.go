package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteRouterInterfaceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteRouterInterfaceRequest) Invoke(client goaliyun.Client) (*DeleteRouterInterfaceResponse, error) {
	resp := &DeleteRouterInterfaceResponse{}
	err := client.Request("vpc", "DeleteRouterInterface", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteRouterInterfaceResponse struct {
	RequestId goaliyun.String
}
