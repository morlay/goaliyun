package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeClassicLinkInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             string `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           string `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeClassicLinkInstancesRequest) Invoke(client goaliyun.Client) (*DescribeClassicLinkInstancesResponse, error) {
	resp := &DescribeClassicLinkInstancesResponse{}
	err := client.Request("ecs", "DescribeClassicLinkInstances", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeClassicLinkInstancesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	PageNumber goaliyun.Integer
	PageSize   goaliyun.Integer
	Links      DescribeClassicLinkInstancesLinkList
}

type DescribeClassicLinkInstancesLink struct {
	InstanceId goaliyun.String
	VpcId      goaliyun.String
}

type DescribeClassicLinkInstancesLinkList []DescribeClassicLinkInstancesLink

func (list *DescribeClassicLinkInstancesLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClassicLinkInstancesLink)
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
