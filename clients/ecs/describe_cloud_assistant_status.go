package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCloudAssistantStatusRequest struct {
	ResourceOwnerId      int64                                       `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                      `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                       `position:"Query" name:"OwnerId"`
	InstanceIds          *DescribeCloudAssistantStatusInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	RegionId             string                                      `position:"Query" name:"RegionId"`
}

func (req *DescribeCloudAssistantStatusRequest) Invoke(client goaliyun.Client) (*DescribeCloudAssistantStatusResponse, error) {
	resp := &DescribeCloudAssistantStatusResponse{}
	err := client.Request("ecs", "DescribeCloudAssistantStatus", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCloudAssistantStatusResponse struct {
	RequestId                       goaliyun.String
	InstanceCloudAssistantStatusSet DescribeCloudAssistantStatusInstanceCloudAssistantStatusList
}

type DescribeCloudAssistantStatusInstanceCloudAssistantStatus struct {
	InstanceId           goaliyun.String
	CloudAssistantStatus goaliyun.String
}

type DescribeCloudAssistantStatusInstanceIdList []string

func (list *DescribeCloudAssistantStatusInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeCloudAssistantStatusInstanceCloudAssistantStatusList []DescribeCloudAssistantStatusInstanceCloudAssistantStatus

func (list *DescribeCloudAssistantStatusInstanceCloudAssistantStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCloudAssistantStatusInstanceCloudAssistantStatus)
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
