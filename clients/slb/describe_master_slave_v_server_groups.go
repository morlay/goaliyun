package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeMasterSlaveVServerGroupsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeMasterSlaveVServerGroupsRequest) Invoke(client goaliyun.Client) (*DescribeMasterSlaveVServerGroupsResponse, error) {
	resp := &DescribeMasterSlaveVServerGroupsResponse{}
	err := client.Request("slb", "DescribeMasterSlaveVServerGroups", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeMasterSlaveVServerGroupsResponse struct {
	RequestId                goaliyun.String
	MasterSlaveVServerGroups DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList
}

type DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup struct {
	MasterSlaveVServerGroupId   goaliyun.String
	MasterSlaveVServerGroupName goaliyun.String
}

type DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList []DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup

func (list *DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveVServerGroupsMasterSlaveVServerGroup)
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
