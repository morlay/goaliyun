package nas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeAccessGroupsRequest struct {
	PageSize        int64  `position:"Query" name:"PageSize"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	PageNumber      int64  `position:"Query" name:"PageNumber"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeAccessGroupsRequest) Invoke(client goaliyun.Client) (*DescribeAccessGroupsResponse, error) {
	resp := &DescribeAccessGroupsResponse{}
	err := client.Request("nas", "DescribeAccessGroups", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeAccessGroupsResponse struct {
	RequestId    goaliyun.String
	TotalCount   goaliyun.Integer
	PageSize     goaliyun.Integer
	PageNumber   goaliyun.Integer
	AccessGroups DescribeAccessGroupsAccessGroupList
}

type DescribeAccessGroupsAccessGroup struct {
	AccessGroupName  goaliyun.String
	AccessGroupType  goaliyun.String
	RuleCount        goaliyun.Integer
	MountTargetCount goaliyun.Integer
	Description      goaliyun.String
}

type DescribeAccessGroupsAccessGroupList []DescribeAccessGroupsAccessGroup

func (list *DescribeAccessGroupsAccessGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessGroupsAccessGroup)
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
