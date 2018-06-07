package slb

import (
	"github.com/morlay/goaliyun"
)

type DeleteMasterSlaveServerGroupRequest struct {
	Access_key_id            string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveServerGroupId string `position:"Query" name:"MasterSlaveServerGroupId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Tags                     string `position:"Query" name:"Tags"`
	RegionId                 string `position:"Query" name:"RegionId"`
}

func (req *DeleteMasterSlaveServerGroupRequest) Invoke(client goaliyun.Client) (*DeleteMasterSlaveServerGroupResponse, error) {
	resp := &DeleteMasterSlaveServerGroupResponse{}
	err := client.Request("slb", "DeleteMasterSlaveServerGroup", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteMasterSlaveServerGroupResponse struct {
	RequestId goaliyun.String
}
