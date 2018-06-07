package hsm

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	PageSize        int64  `position:"Query" name:"PageSize"`
	CurrentPage     int64  `position:"Query" name:"CurrentPage"`
	HsmStatus       int64  `position:"Query" name:"HsmStatus"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesRequest) Invoke(client goaliyun.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("hsm", "DescribeInstances", "2018-01-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	RequestId  goaliyun.String
	TotalCount goaliyun.Integer
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId    goaliyun.String
	RegionId      goaliyun.String
	ZoneId        goaliyun.String
	HsmStatus     goaliyun.Integer
	HsmOem        goaliyun.String
	HsmDeviceType goaliyun.String
	VpcId         goaliyun.String
	VswitchId     goaliyun.String
	Ip            goaliyun.String
	Remark        goaliyun.String
	ExpiredTime   goaliyun.Integer
	CreateTime    goaliyun.Integer
	WhiteList     DescribeInstancesWhiteListList
}

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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

type DescribeInstancesWhiteListList []goaliyun.String

func (list *DescribeInstancesWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
