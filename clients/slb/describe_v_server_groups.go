package slb

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeVServerGroupsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeVServerGroupsRequest) Invoke(client goaliyun.Client) (*DescribeVServerGroupsResponse, error) {
	resp := &DescribeVServerGroupsResponse{}
	err := client.Request("slb", "DescribeVServerGroups", "2014-05-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeVServerGroupsResponse struct {
	RequestId     goaliyun.String
	VServerGroups DescribeVServerGroupsVServerGroupList
}

type DescribeVServerGroupsVServerGroup struct {
	VServerGroupId   goaliyun.String
	VServerGroupName goaliyun.String
}

type DescribeVServerGroupsVServerGroupList []DescribeVServerGroupsVServerGroup

func (list *DescribeVServerGroupsVServerGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVServerGroupsVServerGroup)
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
