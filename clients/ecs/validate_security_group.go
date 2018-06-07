package ecs

import (
	"github.com/morlay/goaliyun"
)

type ValidateSecurityGroupRequest struct {
	NicType              string `position:"Query" name:"NicType"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SourcePort           int64  `position:"Query" name:"SourcePort"`
	SourceIp             string `position:"Query" name:"SourceIp"`
	Direction            string `position:"Query" name:"Direction"`
	DestIp               string `position:"Query" name:"DestIp"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string `position:"Query" name:"IpProtocol"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	DestPort             int64  `position:"Query" name:"DestPort"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ValidateSecurityGroupRequest) Invoke(client goaliyun.Client) (*ValidateSecurityGroupResponse, error) {
	resp := &ValidateSecurityGroupResponse{}
	err := client.Request("ecs", "ValidateSecurityGroup", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ValidateSecurityGroupResponse struct {
	RequestId          goaliyun.String
	Policy             goaliyun.String
	TriggeredGroupRule ValidateSecurityGroupTriggeredGroupRule
}

type ValidateSecurityGroupTriggeredGroupRule struct {
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
