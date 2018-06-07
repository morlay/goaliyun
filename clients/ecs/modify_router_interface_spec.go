package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifyRouterInterfaceSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	RouterInterfaceId    string `position:"Query" name:"RouterInterfaceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Spec                 string `position:"Query" name:"Spec"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyRouterInterfaceSpecRequest) Invoke(client goaliyun.Client) (*ModifyRouterInterfaceSpecResponse, error) {
	resp := &ModifyRouterInterfaceSpecResponse{}
	err := client.Request("ecs", "ModifyRouterInterfaceSpec", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyRouterInterfaceSpecResponse struct {
	RequestId goaliyun.String
	Spec      goaliyun.String
}
