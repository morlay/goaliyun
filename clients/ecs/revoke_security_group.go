package ecs

import (
	"github.com/morlay/goaliyun"
)

type RevokeSecurityGroupRequest struct {
	NicType                 string `position:"Query" name:"NicType"`
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange         string `position:"Query" name:"SourcePortRange"`
	ClientToken             string `position:"Query" name:"ClientToken"`
	SecurityGroupId         string `position:"Query" name:"SecurityGroupId"`
	Description             string `position:"Query" name:"Description"`
	SourceGroupOwnerId      int64  `position:"Query" name:"SourceGroupOwnerId"`
	SourceGroupOwnerAccount string `position:"Query" name:"SourceGroupOwnerAccount"`
	Policy                  string `position:"Query" name:"Policy"`
	PortRange               string `position:"Query" name:"PortRange"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol              string `position:"Query" name:"IpProtocol"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	SourceCidrIp            string `position:"Query" name:"SourceCidrIp"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	Priority                string `position:"Query" name:"Priority"`
	DestCidrIp              string `position:"Query" name:"DestCidrIp"`
	SourceGroupId           string `position:"Query" name:"SourceGroupId"`
	RegionId                string `position:"Query" name:"RegionId"`
}

func (req *RevokeSecurityGroupRequest) Invoke(client goaliyun.Client) (*RevokeSecurityGroupResponse, error) {
	resp := &RevokeSecurityGroupResponse{}
	err := client.Request("ecs", "RevokeSecurityGroup", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RevokeSecurityGroupResponse struct {
	RequestId goaliyun.String
}
