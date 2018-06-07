package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyVpnGatewayAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyVpnGatewayAttributeRequest) Invoke(client goaliyun.Client) (*ModifyVpnGatewayAttributeResponse, error) {
	resp := &ModifyVpnGatewayAttributeResponse{}
	err := client.Request("vpc", "ModifyVpnGatewayAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyVpnGatewayAttributeResponse struct {
	RequestId      goaliyun.String
	VpnGatewayId   goaliyun.String
	VpcId          goaliyun.String
	VSwitchId      goaliyun.String
	InternetIp     goaliyun.String
	IntranetIp     goaliyun.String
	CreateTime     goaliyun.Integer
	EndTime        goaliyun.Integer
	Spec           goaliyun.String
	Name           goaliyun.String
	Description    goaliyun.String
	Status         goaliyun.String
	BusinessStatus goaliyun.String
}
