package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteVServerGroupRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteVServerGroupRequest) Invoke(client goaliyun.Client) (*DeleteVServerGroupResponse, error) {
	resp := &DeleteVServerGroupResponse{}
	err := client.Request("slb", "DeleteVServerGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteVServerGroupResponse struct {
	RequestId goaliyun.String
}
