package vpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeServerRelatedGlobalAccelerationInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ServerType           string `position:"Query" name:"ServerType"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ServerId             string `position:"Query" name:"ServerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeServerRelatedGlobalAccelerationInstancesRequest) Invoke(client goaliyun.Client) (*DescribeServerRelatedGlobalAccelerationInstancesResponse, error) {
	resp := &DescribeServerRelatedGlobalAccelerationInstancesResponse{}
	err := client.Request("vpc", "DescribeServerRelatedGlobalAccelerationInstances", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeServerRelatedGlobalAccelerationInstancesResponse struct {
	RequestId                   goaliyun.String
	GlobalAccelerationInstances DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     goaliyun.String
	GlobalAccelerationInstanceId goaliyun.String
	IpAddress                    goaliyun.String
	ServerIpAddress              goaliyun.String
}

type DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList []DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance

func (list *DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance)
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
