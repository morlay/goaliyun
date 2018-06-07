package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyRouterInterfaceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Spec                 string `position:"Query" name:"Spec"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyRouterInterfaceSpecRequest) Invoke(client goaliyun.Client) (*ModifyRouterInterfaceSpecResponse, error) {
	resp := &ModifyRouterInterfaceSpecResponse{}
	err := client.Request("vpc", "ModifyRouterInterfaceSpec", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyRouterInterfaceSpecResponse struct {
	RequestId goaliyun.String
	Spec      goaliyun.String
}
