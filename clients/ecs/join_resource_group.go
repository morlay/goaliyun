package ecs

import (
	"github.com/morlay/goaliyun"
)

type JoinResourceGroupRequest struct {
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *JoinResourceGroupRequest) Invoke(client goaliyun.Client) (*JoinResourceGroupResponse, error) {
	resp := &JoinResourceGroupResponse{}
	err := client.Request("ecs", "JoinResourceGroup", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type JoinResourceGroupResponse struct {
	RequestId goaliyun.String
}
