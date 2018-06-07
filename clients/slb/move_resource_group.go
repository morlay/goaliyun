package slb

import (
	"github.com/morlay/goaliyun"
)

type MoveResourceGroupRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	Tags                 string `position:"Query" name:"Tags"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	NewResourceGroupId   string `position:"Query" name:"NewResourceGroupId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *MoveResourceGroupRequest) Invoke(client goaliyun.Client) (*MoveResourceGroupResponse, error) {
	resp := &MoveResourceGroupResponse{}
	err := client.Request("slb", "MoveResourceGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MoveResourceGroupResponse struct {
	RequestId goaliyun.String
}
