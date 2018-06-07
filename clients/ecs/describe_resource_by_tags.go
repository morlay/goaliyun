package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeResourceByTagsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
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

func (req *DescribeResourceByTagsRequest) Invoke(client goaliyun.Client) (*DescribeResourceByTagsResponse, error) {
	resp := &DescribeResourceByTagsResponse{}
	err := client.Request("ecs", "DescribeResourceByTags", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeResourceByTagsResponse struct {
	RequestId  goaliyun.String
	PageSize   goaliyun.Integer
	PageNumber goaliyun.Integer
	TotalCount goaliyun.Integer
	Resources  DescribeResourceByTagsResourceList
}

type DescribeResourceByTagsResource struct {
	ResourceId   goaliyun.String
	ResourceType goaliyun.String
	RegionId     goaliyun.String
}

type DescribeResourceByTagsResourceList []DescribeResourceByTagsResource

func (list *DescribeResourceByTagsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourceByTagsResource)
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
