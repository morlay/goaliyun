package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMasterSlaveServerGroupsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMasterSlaveServerGroupsRequest) Invoke(client goaliyun.Client) (*DescribeMasterSlaveServerGroupsResponse, error) {
	resp := &DescribeMasterSlaveServerGroupsResponse{}
	err := client.Request("slb", "DescribeMasterSlaveServerGroups", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMasterSlaveServerGroupsResponse struct {
	RequestId               goaliyun.String
	MasterSlaveServerGroups DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroup struct {
	MasterSlaveServerGroupId   goaliyun.String
	MasterSlaveServerGroupName goaliyun.String
}

type DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList []DescribeMasterSlaveServerGroupsMasterSlaveServerGroup

func (list *DescribeMasterSlaveServerGroupsMasterSlaveServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupsMasterSlaveServerGroup)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
