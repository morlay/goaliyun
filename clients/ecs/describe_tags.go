package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTagsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTagsRequest) Invoke(client goaliyun.Client) (*DescribeTagsResponse, error) {
	resp := &DescribeTagsResponse{}
	err := client.Request("ecs", "DescribeTags", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTagsResponse struct {
	RequestId  goaliyun.String
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	TotalCount goaliyun.Integer
	Tags       DescribeTagsTagList
}

type DescribeTagsTag struct {
	TagKey            goaliyun.String
	TagValue          goaliyun.String
	ResourceTypeCount DescribeTagsResourceTypeCount
}

type DescribeTagsResourceTypeCount struct {
	Instance      goaliyun.Integer
	Disk          goaliyun.Integer
	Volume        goaliyun.Integer
	Image         goaliyun.Integer
	Snapshot      goaliyun.Integer
	Securitygroup goaliyun.Integer
}

type DescribeTagsTagList []DescribeTagsTag

func (list *DescribeTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTag)
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
