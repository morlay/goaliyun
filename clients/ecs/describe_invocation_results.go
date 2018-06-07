package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInvocationResultsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	CommandId            string `position:"Query" name:"CommandId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	InvokeId             string `position:"Query" name:"InvokeId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	InvokeRecordStatus   string `position:"Query" name:"InvokeRecordStatus"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInvocationResultsRequest) Invoke(client goaliyun.Client) (*DescribeInvocationResultsResponse, error) {
	resp := &DescribeInvocationResultsResponse{}
	err := client.Request("ecs", "DescribeInvocationResults", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInvocationResultsResponse struct {
	RequestId  goaliyun.String
	Invocation DescribeInvocationResultsInvocation
}

type DescribeInvocationResultsInvocation struct {
	PageSize          goaliyun.Integer
	PageNumber        goaliyun.Integer
	TotalCount        goaliyun.Integer
	InvocationResults DescribeInvocationResultsInvocationResultList
}

type DescribeInvocationResultsInvocationResult struct {
	CommandId          goaliyun.String
	InvokeId           goaliyun.String
	InstanceId         goaliyun.String
	FinishedTime       goaliyun.String
	Output             goaliyun.String
	InvokeRecordStatus goaliyun.String
	ExitCode           goaliyun.Integer
}

type DescribeInvocationResultsInvocationResultList []DescribeInvocationResultsInvocationResult

func (list *DescribeInvocationResultsInvocationResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationResultsInvocationResult)
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
