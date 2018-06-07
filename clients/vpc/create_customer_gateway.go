package vpc

import (
	"github.com/morlay/goaliyun"
)

type CreateCustomerGatewayRequest struct {
	IpAddress            string `position:"Query" name:"IpAddress"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateCustomerGatewayRequest) Invoke(client goaliyun.Client) (*CreateCustomerGatewayResponse, error) {
	resp := &CreateCustomerGatewayResponse{}
	err := client.Request("vpc", "CreateCustomerGateway", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateCustomerGatewayResponse struct {
	RequestId         goaliyun.String
	CustomerGatewayId goaliyun.String
	IpAddress         goaliyun.String
	Name              goaliyun.String
	Description       goaliyun.String
	CreateTime        goaliyun.Integer
}
