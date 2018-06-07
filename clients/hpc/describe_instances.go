package hpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesRequest struct {
	TOKEN        string `position:"Query" name:"TOKEN"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceType string `position:"Query" name:"InstanceType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesRequest) Invoke(client goaliyun.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("hpc", "DescribeInstances", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	RequestId goaliyun.String
	Instances DescribeInstancesInstanceList
}

type DescribeInstancesInstance struct {
	InstanceId                goaliyun.String
	RegionId                  goaliyun.String
	InstanceType              DescribeInstancesInstanceType
	PackageId                 DescribeInstancesPackageId
	Status                    DescribeInstancesStatus
	InnerIpAddress            goaliyun.String
	JumpserverStatus          DescribeInstancesJumpserverStatus
	JumpserverInnerIpAddress  goaliyun.String
	JumpServerPublicIpAddress goaliyun.String
	CreationTime              goaliyun.String
	ExpireTime                goaliyun.String
}

type DescribeInstancesInstanceType struct {
	StringValue goaliyun.String
}

type DescribeInstancesPackageId struct {
	StringValue goaliyun.String
}

type DescribeInstancesStatus struct {
	StringValue goaliyun.String
}

type DescribeInstancesJumpserverStatus struct {
	StringValue goaliyun.String
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
