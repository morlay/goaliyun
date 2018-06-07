package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeSecurityGroupAttributeRequest struct {
	NicType              string `position:"Query" name:"NicType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SecurityGroupId      string `position:"Query" name:"SecurityGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Direction            string `position:"Query" name:"Direction"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeSecurityGroupAttributeRequest) Invoke(client goaliyun.Client) (*DescribeSecurityGroupAttributeResponse, error) {
	resp := &DescribeSecurityGroupAttributeResponse{}
	err := client.Request("ecs", "DescribeSecurityGroupAttribute", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupAttributeResponse struct {
	RequestId         goaliyun.String
	RegionId          goaliyun.String
	SecurityGroupId   goaliyun.String
	Description       goaliyun.String
	SecurityGroupName goaliyun.String
	VpcId             goaliyun.String
	InnerAccessPolicy goaliyun.String
	Permissions       DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermission struct {
	IpProtocol              goaliyun.String
	PortRange               goaliyun.String
	SourceGroupId           goaliyun.String
	SourceGroupName         goaliyun.String
	SourceCidrIp            goaliyun.String
	Policy                  goaliyun.String
	NicType                 goaliyun.String
	SourceGroupOwnerAccount goaliyun.String
	DestGroupId             goaliyun.String
	DestGroupName           goaliyun.String
	DestCidrIp              goaliyun.String
	DestGroupOwnerAccount   goaliyun.String
	Priority                goaliyun.String
	Direction               goaliyun.String
	Description             goaliyun.String
	CreateTime              goaliyun.String
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
