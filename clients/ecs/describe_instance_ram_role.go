package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstanceRamRoleRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	RamRoleName          string `position:"Query" name:"RamRoleName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstanceRamRoleRequest) Invoke(client goaliyun.Client) (*DescribeInstanceRamRoleResponse, error) {
	resp := &DescribeInstanceRamRoleResponse{}
	err := client.Request("ecs", "DescribeInstanceRamRole", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstanceRamRoleResponse struct {
	RequestId           goaliyun.String
	RegionId            goaliyun.String
	TotalCount          goaliyun.Integer
	InstanceRamRoleSets DescribeInstanceRamRoleInstanceRamRoleSetList
}

type DescribeInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  goaliyun.String
	RamRoleName goaliyun.String
}

type DescribeInstanceRamRoleInstanceRamRoleSetList []DescribeInstanceRamRoleInstanceRamRoleSet

func (list *DescribeInstanceRamRoleInstanceRamRoleSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceRamRoleInstanceRamRoleSet)
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
