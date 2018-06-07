package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSecurityGroupsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	FuzzyQuery           string `position:"Query" name:"FuzzyQuery"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	IsQueryEcsCount      string `position:"Query" name:"IsQueryEcsCount"`
	NetworkType          string `position:"Query" name:"NetworkType"`
	SecurityGroupName    string `position:"Query" name:"SecurityGroupName"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	ResourceGroupId      string `position:"Query" name:"ResourceGroupId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	DryRun               string `position:"Query" name:"DryRun"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SecurityGroupIds     string `position:"Query" name:"SecurityGroupIds"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	VpcId                string `position:"Query" name:"VpcId"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityGroupsRequest) Invoke(client goaliyun.Client) (*DescribeSecurityGroupsResponse, error) {
	resp := &DescribeSecurityGroupsResponse{}
	err := client.Request("ecs", "DescribeSecurityGroups", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupsResponse struct {
	RequestId      goaliyun.String
	RegionId       goaliyun.String
	TotalCount     goaliyun.Integer
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	SecurityGroups DescribeSecurityGroupsSecurityGroupList
}

type DescribeSecurityGroupsSecurityGroup struct {
	SecurityGroupId         goaliyun.String
	Description             goaliyun.String
	SecurityGroupName       goaliyun.String
	VpcId                   goaliyun.String
	CreationTime            goaliyun.String
	AvailableInstanceAmount goaliyun.Integer
	EcsCount                goaliyun.Integer
	Tags                    DescribeSecurityGroupsTagList
}

type DescribeSecurityGroupsTag struct {
	TagKey   goaliyun.String
	TagValue goaliyun.String
}

type DescribeSecurityGroupsSecurityGroupList []DescribeSecurityGroupsSecurityGroup

func (list *DescribeSecurityGroupsSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupsSecurityGroup)
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

type DescribeSecurityGroupsTagList []DescribeSecurityGroupsTag

func (list *DescribeSecurityGroupsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupsTag)
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
