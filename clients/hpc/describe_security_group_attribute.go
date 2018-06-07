package hpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSecurityGroupAttributeRequest struct {
	TOKEN    string `position:"Query" name:"TOKEN"`
	NicType  string `position:"Query" name:"NicType"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityGroupAttributeRequest) Invoke(client goaliyun.Client) (*DescribeSecurityGroupAttributeResponse, error) {
	resp := &DescribeSecurityGroupAttributeResponse{}
	err := client.Request("hpc", "DescribeSecurityGroupAttribute", "2016-12-13", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupAttributeResponse struct {
	Records DescribeSecurityGroupAttributeRecords
}

type DescribeSecurityGroupAttributeRecords struct {
	RegionId    goaliyun.String
	Permissions DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermission struct {
	SourceIp goaliyun.String
	Policy   DescribeSecurityGroupAttributePolicy
	NicType  DescribeSecurityGroupAttributeNicType
	Priority goaliyun.String
	Time     goaliyun.String
}

type DescribeSecurityGroupAttributePolicy struct {
	StringValue goaliyun.String
}

type DescribeSecurityGroupAttributeNicType struct {
	StringValue goaliyun.String
}

type DescribeSecurityGroupAttributePermissionList []DescribeSecurityGroupAttributePermission

func (list *DescribeSecurityGroupAttributePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupAttributePermission)
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
