package hpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeInstancesInSecurityGroupRequest struct {
	TOKEN    string `position:"Query" name:"TOKEN"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeInstancesInSecurityGroupRequest) Invoke(client goaliyun.Client) (*DescribeInstancesInSecurityGroupResponse, error) {
	resp := &DescribeInstancesInSecurityGroupResponse{}
	err := client.Request("hpc", "DescribeInstancesInSecurityGroup", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeInstancesInSecurityGroupResponse struct {
	RequestId goaliyun.String
	Records   DescribeInstancesInSecurityGroupRecordList
}

type DescribeInstancesInSecurityGroupRecord struct {
	RegionId       goaliyun.String
	InstanceId     goaliyun.String
	InnerIpAddress goaliyun.String
}

type DescribeInstancesInSecurityGroupRecordList []DescribeInstancesInSecurityGroupRecord

func (list *DescribeInstancesInSecurityGroupRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInSecurityGroupRecord)
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
