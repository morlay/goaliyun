package ecs

import (
	"github.com/morlay/goaliyun"
)

type ModifySecurityGroupEgressRuleRequest struct {
	NicType               string `position:"Query" name:"NicType"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange       string `position:"Query" name:"SourcePortRange"`
	ClientToken           string `position:"Query" name:"ClientToken"`
	SecurityGroupId       string `position:"Query" name:"SecurityGroupId"`
	Description           string `position:"Query" name:"Description"`
	Policy                string `position:"Query" name:"Policy"`
	PortRange             string `position:"Query" name:"PortRange"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol            string `position:"Query" name:"IpProtocol"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	SourceCidrIp          string `position:"Query" name:"SourceCidrIp"`
	DestGroupId           string `position:"Query" name:"DestGroupId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	DestGroupOwnerAccount string `position:"Query" name:"DestGroupOwnerAccount"`
	Priority              string `position:"Query" name:"Priority"`
	DestCidrIp            string `position:"Query" name:"DestCidrIp"`
	DestGroupOwnerId      int64  `position:"Query" name:"DestGroupOwnerId"`
	RegionId              string `position:"Query" name:"RegionId"`
}

func (req *ModifySecurityGroupEgressRuleRequest) Invoke(client goaliyun.Client) (*ModifySecurityGroupEgressRuleResponse, error) {
	resp := &ModifySecurityGroupEgressRuleResponse{}
	err := client.Request("ecs", "ModifySecurityGroupEgressRule", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifySecurityGroupEgressRuleResponse struct {
	RequestId goaliyun.String
}
