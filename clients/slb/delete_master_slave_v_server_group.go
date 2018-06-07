package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteMasterSlaveVServerGroupRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveVServerGroupId string `position:"Query" name:"MasterSlaveVServerGroupId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
	RegionId                  string `position:"Query" name:"RegionId"`
}

func (req *DeleteMasterSlaveVServerGroupRequest) Invoke(client goaliyun.Client) (*DeleteMasterSlaveVServerGroupResponse, error) {
	resp := &DeleteMasterSlaveVServerGroupResponse{}
	err := client.Request("slb", "DeleteMasterSlaveVServerGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMasterSlaveVServerGroupResponse struct {
	RequestId goaliyun.String
}
