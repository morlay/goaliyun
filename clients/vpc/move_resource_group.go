package vpc

import (
	"github.com/morlay/goaliyun"
)

type MoveResourceGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	NewResourceGroupId   string `position:"Query" name:"NewResourceGroupId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *MoveResourceGroupRequest) Invoke(client goaliyun.Client) (*MoveResourceGroupResponse, error) {
	resp := &MoveResourceGroupResponse{}
	err := client.Request("vpc", "MoveResourceGroup", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type MoveResourceGroupResponse struct {
	RequestId goaliyun.String
}
