package vpc

import (
	"github.com/morlay/goaliyun"
)

type DeleteCustomerGatewayRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteCustomerGatewayRequest) Invoke(client goaliyun.Client) (*DeleteCustomerGatewayResponse, error) {
	resp := &DeleteCustomerGatewayResponse{}
	err := client.Request("vpc", "DeleteCustomerGateway", "2016-04-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteCustomerGatewayResponse struct {
	RequestId goaliyun.String
}
