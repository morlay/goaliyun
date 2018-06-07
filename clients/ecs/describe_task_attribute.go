package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeTaskAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TaskId               string `position:"Query" name:"TaskId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeTaskAttributeRequest) Invoke(client goaliyun.Client) (*DescribeTaskAttributeResponse, error) {
	resp := &DescribeTaskAttributeResponse{}
	err := client.Request("ecs", "DescribeTaskAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeTaskAttributeResponse struct {
	RequestId            goaliyun.String
	TaskId               goaliyun.String
	RegionId             goaliyun.String
	TaskAction           goaliyun.String
	TaskStatus           goaliyun.String
	TaskProcess          goaliyun.String
	SupportCancel        goaliyun.String
	TotalCount           goaliyun.Integer
	SuccessCount         goaliyun.Integer
	FailedCount          goaliyun.Integer
	CreationTime         goaliyun.String
	FinishedTime         goaliyun.String
	OperationProgressSet DescribeTaskAttributeOperationProgressList
}

type DescribeTaskAttributeOperationProgress struct {
	OperationStatus goaliyun.String
	ErrorCode       goaliyun.String
	ErrorMsg        goaliyun.String
	RelatedItemSet  DescribeTaskAttributeRelatedItemList
}

type DescribeTaskAttributeRelatedItem struct {
	Name  goaliyun.String
	Value goaliyun.String
}

type DescribeTaskAttributeOperationProgressList []DescribeTaskAttributeOperationProgress

func (list *DescribeTaskAttributeOperationProgressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTaskAttributeOperationProgress)
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

type DescribeTaskAttributeRelatedItemList []DescribeTaskAttributeRelatedItem

func (list *DescribeTaskAttributeRelatedItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTaskAttributeRelatedItem)
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
