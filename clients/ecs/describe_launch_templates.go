package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeLaunchTemplatesRequest struct {
	LaunchTemplateNames  *DescribeLaunchTemplatesLaunchTemplateNameList `position:"Query" type:"Repeated" name:"LaunchTemplateName"`
	ResourceOwnerId      int64                                          `position:"Query" name:"ResourceOwnerId"`
	PageNumber           int64                                          `position:"Query" name:"PageNumber"`
	PageSize             int64                                          `position:"Query" name:"PageSize"`
	LaunchTemplateIds    *DescribeLaunchTemplatesLaunchTemplateIdList   `position:"Query" type:"Repeated" name:"LaunchTemplateId"`
	ResourceOwnerAccount string                                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                         `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                          `position:"Query" name:"OwnerId"`
	RegionId             string                                         `position:"Query" name:"RegionId"`
}

func (req *DescribeLaunchTemplatesRequest) Invoke(client goaliyun.Client) (*DescribeLaunchTemplatesResponse, error) {
	resp := &DescribeLaunchTemplatesResponse{}
	err := client.Request("ecs", "DescribeLaunchTemplates", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeLaunchTemplatesResponse struct {
	RequestId          goaliyun.String
	TotalCount         goaliyun.Integer
	PageNumber         goaliyun.Integer
	PageSize           goaliyun.Integer
	LaunchTemplateSets DescribeLaunchTemplatesLaunchTemplateSetList
}

type DescribeLaunchTemplatesLaunchTemplateSet struct {
	CreateTime           goaliyun.String
	ModifiedTime         goaliyun.String
	LaunchTemplateId     goaliyun.String
	LaunchTemplateName   goaliyun.String
	DefaultVersionNumber goaliyun.Integer
	LatestVersionNumber  goaliyun.Integer
	CreatedBy            goaliyun.String
}

type DescribeLaunchTemplatesLaunchTemplateNameList []string

func (list *DescribeLaunchTemplatesLaunchTemplateNameList) UnmarshalJSON(data []byte) error {
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

type DescribeLaunchTemplatesLaunchTemplateIdList []string

func (list *DescribeLaunchTemplatesLaunchTemplateIdList) UnmarshalJSON(data []byte) error {
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

type DescribeLaunchTemplatesLaunchTemplateSetList []DescribeLaunchTemplatesLaunchTemplateSet

func (list *DescribeLaunchTemplatesLaunchTemplateSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLaunchTemplatesLaunchTemplateSet)
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
