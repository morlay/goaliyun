package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSecurityGroupReferencesRequest struct {
	ResourceOwnerId      int64                                               `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                              `position:"Query" name:"OwnerAccount"`
	SecurityGroupIds     *DescribeSecurityGroupReferencesSecurityGroupIdList `position:"Query" type:"Repeated" name:"SecurityGroupId"`
	OwnerId              int64                                               `position:"Query" name:"OwnerId"`
	RegionId             string                                              `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityGroupReferencesRequest) Invoke(client goaliyun.Client) (*DescribeSecurityGroupReferencesResponse, error) {
	resp := &DescribeSecurityGroupReferencesResponse{}
	err := client.Request("ecs", "DescribeSecurityGroupReferences", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupReferencesResponse struct {
	RequestId               goaliyun.String
	SecurityGroupReferences DescribeSecurityGroupReferencesSecurityGroupReferenceList
}

type DescribeSecurityGroupReferencesSecurityGroupReference struct {
	SecurityGroupId           goaliyun.String
	ReferencingSecurityGroups DescribeSecurityGroupReferencesReferencingSecurityGroupList
}

type DescribeSecurityGroupReferencesReferencingSecurityGroup struct {
	AliUid          goaliyun.String
	SecurityGroupId goaliyun.String
}

type DescribeSecurityGroupReferencesSecurityGroupIdList []string

func (list *DescribeSecurityGroupReferencesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeSecurityGroupReferencesSecurityGroupReferenceList []DescribeSecurityGroupReferencesSecurityGroupReference

func (list *DescribeSecurityGroupReferencesSecurityGroupReferenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupReferencesSecurityGroupReference)
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

type DescribeSecurityGroupReferencesReferencingSecurityGroupList []DescribeSecurityGroupReferencesReferencingSecurityGroup

func (list *DescribeSecurityGroupReferencesReferencingSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupReferencesReferencingSecurityGroup)
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
