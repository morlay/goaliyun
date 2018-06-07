package vpc

import (
	"github.com/morlay/goaliyun"
)

type ModifyCustomerGatewayAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *ModifyCustomerGatewayAttributeRequest) Invoke(client goaliyun.Client) (*ModifyCustomerGatewayAttributeResponse, error) {
	resp := &ModifyCustomerGatewayAttributeResponse{}
	err := client.Request("vpc", "ModifyCustomerGatewayAttribute", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyCustomerGatewayAttributeResponse struct {
	RequestId         goaliyun.String
	CustomerGatewayId goaliyun.String
	IpAddress         goaliyun.String
	Name              goaliyun.String
	Description       goaliyun.String
	CreateTime        goaliyun.Integer
}
