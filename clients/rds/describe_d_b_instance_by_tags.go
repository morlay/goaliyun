package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDBInstanceByTagsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	ProxyId              string `position:"Query" name:"ProxyId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeDBInstanceByTagsRequest) Invoke(client goaliyun.Client) (*DescribeDBInstanceByTagsResponse, error) {
	resp := &DescribeDBInstanceByTagsResponse{}
	err := client.Request("rds", "DescribeDBInstanceByTags", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDBInstanceByTagsResponse struct {
	RequestId        goaliyun.String
	PageNumber       goaliyun.Integer
	PageRecordCount  goaliyun.Integer
	TotalRecordCount goaliyun.Integer
	Items            DescribeDBInstanceByTagsDBInstanceTagList
}

type DescribeDBInstanceByTagsDBInstanceTag struct {
	DBInstanceId goaliyun.String
	Tags         DescribeDBInstanceByTagsTagList
}

type DescribeDBInstanceByTagsTag struct {
	TagKey   goaliyun.String
	TagValue goaliyun.String
}

type DescribeDBInstanceByTagsDBInstanceTagList []DescribeDBInstanceByTagsDBInstanceTag

func (list *DescribeDBInstanceByTagsDBInstanceTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceByTagsDBInstanceTag)
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

type DescribeDBInstanceByTagsTagList []DescribeDBInstanceByTagsTag

func (list *DescribeDBInstanceByTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceByTagsTag)
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
