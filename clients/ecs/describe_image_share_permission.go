package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeImageSharePermissionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeImageSharePermissionRequest) Invoke(client goaliyun.Client) (*DescribeImageSharePermissionResponse, error) {
	resp := &DescribeImageSharePermissionResponse{}
	err := client.Request("ecs", "DescribeImageSharePermission", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeImageSharePermissionResponse struct {
	RequestId   goaliyun.String
	RegionId    goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	ImageId     goaliyun.String
	ShareGroups DescribeImageSharePermissionShareGroupList
	Accounts    DescribeImageSharePermissionAccountList
}

type DescribeImageSharePermissionShareGroup struct {
	Group goaliyun.String
}

type DescribeImageSharePermissionAccount struct {
	AliyunId goaliyun.String
}

type DescribeImageSharePermissionShareGroupList []DescribeImageSharePermissionShareGroup

func (list *DescribeImageSharePermissionShareGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSharePermissionShareGroup)
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

type DescribeImageSharePermissionAccountList []DescribeImageSharePermissionAccount

func (list *DescribeImageSharePermissionAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSharePermissionAccount)
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
