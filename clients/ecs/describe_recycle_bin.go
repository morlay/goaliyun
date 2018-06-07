package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeRecycleBinRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeRecycleBinRequest) Invoke(client goaliyun.Client) (*DescribeRecycleBinResponse, error) {
	resp := &DescribeRecycleBinResponse{}
	err := client.Request("ecs", "DescribeRecycleBin", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeRecycleBinResponse struct {
	RequestId        goaliyun.String
	TotalCount       goaliyun.Integer
	RecycleBinModels DescribeRecycleBinRecycleBinModelList
}

type DescribeRecycleBinRecycleBinModel struct {
	ResourceId        goaliyun.String
	RegionId          goaliyun.String
	ResourceType      goaliyun.String
	CreationTime      goaliyun.String
	Status            goaliyun.String
	RelationResources DescribeRecycleBinRelationResourceList
}

type DescribeRecycleBinRelationResource struct {
	RelationResourceId   goaliyun.String
	RelationResourceType goaliyun.String
}

type DescribeRecycleBinRecycleBinModelList []DescribeRecycleBinRecycleBinModel

func (list *DescribeRecycleBinRecycleBinModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecycleBinRecycleBinModel)
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

type DescribeRecycleBinRelationResourceList []DescribeRecycleBinRelationResource

func (list *DescribeRecycleBinRelationResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecycleBinRelationResource)
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
