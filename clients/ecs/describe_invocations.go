package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInvocationsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InvokeStatus         string `position:"Query" name:"InvokeStatus"`
	CommandId            string `position:"Query" name:"CommandId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	InvokeId             string `position:"Query" name:"InvokeId"`
	Timed                string `position:"Query" name:"Timed"`
	CommandName          string `position:"Query" name:"CommandName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CommandType          string `position:"Query" name:"CommandType"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeInvocationsRequest) Invoke(client goaliyun.Client) (*DescribeInvocationsResponse, error) {
	resp := &DescribeInvocationsResponse{}
	err := client.Request("ecs", "DescribeInvocations", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInvocationsResponse struct {
	RequestId   goaliyun.String
	TotalCount  goaliyun.Integer
	PageNumber  goaliyun.Integer
	PageSize    goaliyun.Integer
	Invocations DescribeInvocationsInvocationList
}

type DescribeInvocationsInvocation struct {
	InvokeId        goaliyun.String
	CommandId       goaliyun.String
	CommandType     goaliyun.String
	CommandName     goaliyun.String
	Frequency       goaliyun.String
	Timed           bool
	InvokeStatus    goaliyun.String
	InvokeInstances DescribeInvocationsInvokeInstanceList
}

type DescribeInvocationsInvokeInstance struct {
	InstanceId           goaliyun.String
	InstanceInvokeStatus goaliyun.String
}

type DescribeInvocationsInvocationList []DescribeInvocationsInvocation

func (list *DescribeInvocationsInvocationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvocation)
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

type DescribeInvocationsInvokeInstanceList []DescribeInvocationsInvokeInstance

func (list *DescribeInvocationsInvokeInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvokeInstance)
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
